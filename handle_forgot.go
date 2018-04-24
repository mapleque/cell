package cell

import (
	"github.com/mapleque/kelp/web"
	"html/template"
)

func (this *Server) HandleForgotPage(c *web.Context) {
	this.tpl(c, "forgot.html", nil)
}

func (this *Server) HandleForgot(c *web.Context) {
	this.tpl(c, "tips.html", Tips{SUCCESS, "Please check your email to continue!"})
}

func (this *Server) HandleResetPage(c *web.Context) {
	this.tpl(c, "reset.html", nil)
}

func (this *Server) HandleReset(c *web.Context) {
	this.tpl(c, "tips.html", Tips{SUCCESS, "Password has been reset!"})
}
