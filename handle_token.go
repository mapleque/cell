package cell

import (
	"github.com/mapleque/kelp/web"
)

func (this *Server) HandleToken(c *web.Context) {
	grantType := c.Request.FormValue("grant_type")
	if grantType == "" {
		c.Json(map[string]interface{}{
			"error": ERROR_INVALID_REQUEST,
		})
	}
	token := getBasicToken(c)
	if token == "" {
		c.Json(map[string]interface{}{
			"error": ERROR_INVALID_CLIENT,
		})
	}
	if !this.ds.CheckClientToken(token) {
		c.Json(map[string]interface{}{
			"error": ERROR_UNAUTHORIZED_CLIENT,
		})
	}
	switch grantType {
	case "authorization_code":
		code := c.Request.FormValue("code")
		// no need to deal this param
		// redirectUri := c.Request.FormValue("redirect_uri")
		resp, err := this.ds.GeneralToken(code)
		if err != nil {
			c.Json(map[string]interface{}{
				"error": ERROR_INVALID_GRANT,
			})
			return
		}
		c.Json(resp)
	case "refresh_token":
		scope := c.Request.FormValue("scope")
		if !checkSupportScope(scope) {
			c.Json(map[string]interface{}{
				"error": ERROR_INVALID_SCOPE,
			})
			return
		}
		refreshToken := c.Request.FormValue("refresh_token")
		resp, err := this.ds.RefreshToken(refreshToken)
		if err != nil {
			c.Json(map[string]interface{}{
				"error": ERROR_INVALID_GRANT,
			})
			return
		}
		c.Json(resp)
	default:
		c.Json(map[string]interface{}{
			"error": ERROR_UNSUPPORTED_GRANT_TYPE,
		})
	}
}

func getBasicToken(c *web.Context) string {
	token := c.Request.Header.Get("Authorization")
	if len(token) > 6 && token[:6] == "Basic " {
		return token[6:]
	}
	return ""
}
