package cell

import (
	"github.com/mapleque/kelp/web"
	"html/template"
)

func (this *Server) HandleRegisterPage(c *web.Context) {
	this.tpl(c, "register.html", nil)
}

func (this *Server) HandleRegister(c *web.Context) {}
