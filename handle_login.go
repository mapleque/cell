package cell

import (
	"github.com/mapleque/kelp/web"
	"html/template"
)

func (this *Server) HandleLoginPage(c *web.Context) {
	this.tpl(c, "login.html", nil)
}

func (this *Server) HandleLogin(c *web.Context) {
}

func (this *Server) HandleLogout(c *web.Context) {}
