package cell

import (
	"github.com/mapleque/kelp/web"
	"html/template"
)

func (this *Server) HandleProfilePage(c *web.Context) {
	this.tpl(c, "profile.html", nil)
}

func (this *Server) HandleProfile(c *web.Context) {}
