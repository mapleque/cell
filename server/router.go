package server

import "net/http"

func (s *Server) initRouter() {
	s.handler.Handle("/", http.FileServer(http.Dir(s.staticPath)))
	s.handler.Handle(
		"/.well-known/openid-configuration",
		http.FileServer(http.Dir(s.staticPath+"openid-configuration.json")),
	)

	// /api/kerberos
	{
		s.handleAPI("/kerberos/auth", s.handleKerberosAuth)
		s.handleAPI("/kerberos/grant", s.handleKerberosGrant)
	}
	// /api/oidc
	{
		s.handleAPI("/oidc/certs", s.handleOidcCerts)
		s.handleAPI("/oidc/auth", func(w http.ResponseWriter, r *http.Request) {
			u := r.URL
			u.Path = "/oidc/auth.html"
			http.Redirect(w, r, u.String(), 302)
		})
		s.handleAPI("/oidc/authinfo", s.handleOidcAuthinfo)
		s.handleAPI("/oidc/authed", s.handleOidcAuthed)
		s.handleAPI("/oidc/token", s.handleOidcToken)
		s.handleAPI("/oidc/userinfo", s.handleOidcUserinfo)
	}
	// /api/user
	{
		s.handleAPI("/user/login", s.handleUserLogin)
		s.handleAPI("/user/logout", s.handleUserLogout)
		s.handleAPI("/user/captcha", s.handleUserCaptcha)
		s.handleAPI("/user/register", s.handleUserRegister)
		s.handleAPI("/user/forgot", s.handleUserForgot)
		s.handleAPI("/user/reset", s.handleUserReset)
		s.handleAPI("/user/profile", s.handleUserProfile)
	}
	// /api/authorization
	{
		s.handleAPI("/authorization/list", s.handleAuthorizationList)
		s.handleAPI("/authorization/delete", s.handleAuthorizationDelete)
	}
	// /api/app
	{
		s.handleAPI("/app/create", s.handleAppCreate)
		s.handleAPI("/app/update", s.handleAppUpdate)
		s.handleAPI("/app/delete", s.handleAppDelete)
		s.handleAPI("/app/list", s.handleAppList)
	}
}
