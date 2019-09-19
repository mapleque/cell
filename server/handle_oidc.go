package server

import "net/http"

func (s *Server) handleOidcCerts(w http.ResponseWriter, r *http.Request) {
	keys := s.oidc.GetKeys()
	respRaw(w, 0, keys)
}

func (s *Server) handleOidcAuthinfo(w http.ResponseWriter, r *http.Request) {
	responseType := r.FormValue("response_type")
	if !checkSupportResponseType(responseType) {
		resp(w, 10001, "不支持的response_type")
	}

	scope := r.FormValue("scope")
	if !checkSupportScope(scope) {
		resp(w, 10001, "不支持的scope")
		return
	}

	clientID := r.FormValue("client_id")
	if clientID == "" {
		resp(w, 10001, "client_id不能为空")
		return
	}

	client, exist := s.oidc.FindClient(clientID)
	if !exist {
		resp(w, 11001, "client_id未注册")
		return
	}

	resp(w, 0, client)
}

func (s *Server) handleOidcAuthed(w http.ResponseWriter, r *http.Request) {
	responseType := r.FormValue("response_type")
	if !checkSupportResponseType(responseType) {
		resp(w, 10001, "不支持的response_type")
	}
	scope := r.FormValue("scope")
	if !checkSupportScope(scope) {
		resp(w, 10001, "不支持的scope")
		return
	}
	clientID := r.FormValue("client_id")
	redirectURI := r.FormValue("redirect_uri")
	state := r.FormValue("state")

	user, err := s.auth.GetUser(r)
	if err != nil {
		resp(w, 10002, err)
		return
	}

	code, err := s.oidc.Auth(
		user,
		responseType,
		clientID,
		redirectURI,
		scope,
		state,
	)
	if err != nil {
		resp(w, 10001, err)
		return
	}

	resp(w, 0, code)
}

func (s *Server) handleOidcToken(w http.ResponseWriter, r *http.Request) {
	grantType := r.FormValue("grant_type")
	if grantType == "" {
		respRaw(w, 10001, "grant_type不能为空")
		return
	}
	username, password, hasToken := r.BasicAuth()
	if !hasToken {
		respRaw(w, 10002, "需要Basic Auth Token")
		return
	}

	if !s.oidc.CheckClient(username, password) {
		respRaw(w, 10002, "错误的Basic Auth Token")
		return
	}

	switch grantType {
	case "authorization_code":
		code := r.FormValue("code")
		token, err := s.oidc.GrantToken(code)
		if err != nil {
			respRaw(w, 10001, err)
			return
		}
		respRaw(w, 0, token)
		return
	case "refresh_token":
		scope := r.FormValue("scope")
		if !checkSupportScope(scope) {
			respRaw(w, 10001, "不支持的scope")
			return
		}

		refreshToken := r.FormValue("refresh_token")
		if refreshToken == "" {
			respRaw(w, 10001, "refresh_token不能为空")
			return
		}
		token, err := s.oidc.RefreshToken(refreshToken)
		if err != nil {
			respRaw(w, 10001, err)
			return
		}
		respRaw(w, 0, token)
		return
	default:
		respRaw(w, 10001, "不支持的grant_type")
		return
	}
}

func (s *Server) handleOidcUserinfo(w http.ResponseWriter, r *http.Request) {
	accessToken := r.FormValue("access_token")
	if accessToken == "" {
		respRaw(w, 10001, "access_token不能为空")
		return
	}
	user, err := s.oidc.FindUser(accessToken)
	if err != nil {
		respRaw(w, 10001, err)
		return
	}
	respRaw(w, 0, user)
}

func checkSupportScope(scope string) bool {
	switch scope {
	case "openid", "email", "profile":
		return true
	}
	return false
}

func checkSupportResponseType(responseType string) bool {
	switch responseType {
	case "code",
		"token",
		"id_token",
		"code token",
		"code id_token",
		"token id_token",
		"code token id_token",
		"none":
		return true
	}
	return false
}
