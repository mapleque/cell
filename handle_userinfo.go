package cell

import (
	"github.com/mapleque/kelp/web"
)

func (this *Server) HandleUserinfo(c *web.Context) {
	accessToken := c.Request.FormValue("access_token")
	if accessToken == "" {
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
	resp, err := this.ds.GetUserinfoByToken(accessToken)
	if err != nil {
		c.Json(map[string]interface{}{
			"error": ERROR_INVALID_GRANT,
		})
		return
	}
	c.Json(resp)
}
