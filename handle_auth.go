package cell

import (
	"github.com/mapleque/kelp/web"
	"html/template"
	"net/url"
)

// GET /auth
// Content-Type: application/x-www-form-urlencoded
// params:
// - response_type
//     - code
//     - token
//     - id_token
//     - code token
//     - code id_token
//     - token id_token
//     - code token id_token
//     - none
// - client_id
// - redirect_uri
// - scope
//     - openid
//     - email
//     - profile
// - state
//
// code response:
// 302
// Location redirect_uri with params:
// - code
// - state
//
// token response:
// 302
// Location redirect_uri with params:
// - access_token
// - token_type
//     - bearer
// - expires_in
// - scope
//     - openid
//     - email
//     - profile
// - state
//
// id_token response:
// - id_token
// - state
//
// error response:
// 302
// Location redirect_uri with params:
// - error
//     - invalid_request
//     - unauthorized_client
//     - access_denied
//     - unsupported_response_type
//     - invalid_scope
//     - server_error
//     - temporarily_unavailable
// - error_description
// - error_uri
// - state
func (this *Server) HandleAuthPage(c *web.Context) {
	responseType := c.QueryDefault("response_type", "")
	clientId := c.QueryDefault("client_id", "")
	redirectUri := c.QueryDefault("redirect_uri", "")
	scope := c.QueryDefault("scope", "")
	state := c.QueryDefault("state", "")
	if responseType == "" ||
		clientId == "" ||
		redirectUri == "" ||
		scope == "" {
		c.Redirect(302, errorLocation(redirectUri, ERROR_INVALID_REQUEST))
		return
	}
	if !checkSupportResponseType(responseType) {
		c.Redirect(302, errorLocation(redirectUri, ERROR_UNSUPPORTED_RESPONSE_TYPE))
	}
	if !checkSupportScope(scope) {
		c.Redirect(302, errorLocation(redirectUri, ERROR_INVALID_SCOPE))
	}
	clientName, clientDesc, ok, err := this.ds.CheckClientAuthAndGet(clientId, redirectUri)
	if err != nil {
		c.Redirect(302, errorLocation(redirectUri, ERROR_SERVER_ERROR))
	}
	if !ok {
		c.Redirect(302, errorLocation(redirectUri, ERROR_UNAUTHORIZED_CLIENT))
	}
	t := template.Must(template.ParseFiles("../tpl/auth.html"))
	t.Execute(c.ResponseWriter, struct {
		ResponseType string
		Scope        string
		ClientId     string
		RedirectUri  string
		ClientName   string
		ClientDesc   string
		State        string
	}{
		responseType,
		scope,
		clientId,
		redirectUri,
		clientName,
		clientDesc,
		state,
	})
}

func (this *Server) HandleAuth(c *web.Context) {
	redirectUri := c.Request.FormValue("redirect_uri")
	resp, err := this.ds.Auth(
		c.Request.FormValue("response_type"),
		c.Request.FormValue("client_id"),
		redirectUri,
		c.Request.FormValue("scope"),
		c.Request.FormValue("state"),
	)
	if err != nil {
		c.Redirect(302, errorLocation(redirectUri, ERROR_SERVER_ERROR))
	}
	c.Redirect(302, successLocation(redirectUri, resp))
}

func errorLocation(uri, status string) string {
	return uri + "?" + url.Values{"error": {status}}.Encode()
}

func successLocation(uri string, resp map[string]string) string {
	values := url.Values{}
	for k, v := range resp {
		values.Add(k, v)
	}
	return uri + "?" + values.Encode()
}

func checkSupportResponseType(responseType string) bool {
	return true
}

func checkSupportScope(scope string) bool {
	return true
}
