package cell

import (
	"github.com/mapleque/kelp/web"
	"html/template"
)

func (this *Server) HandleHomePage(c *web.Context) {
	this.tpl(c, "home.html", nil)
}

func (this *Server) HandleOpenidConfiguration(c *web.Context) {
	this.tpl(c, "openid-configuration.json", nil)
}
