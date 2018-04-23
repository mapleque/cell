package cell

import (
	"github.com/mapleque/kelp/web"
)

func (this *Server) HandleHomePage(c *web.Context)            {}
func (this *Server) HandleLoginPage(c *web.Context)           {}
func (this *Server) HandleLogin(c *web.Context)               {}
func (this *Server) HandleRegisterPage(c *web.Context)        {}
func (this *Server) HandleRegister(c *web.Context)            {}
func (this *Server) HandleLogout(c *web.Context)              {}
func (this *Server) HandleForgotPage(c *web.Context)          {}
func (this *Server) HandleForgot(c *web.Context)              {}
func (this *Server) HandleResetPage(c *web.Context)           {}
func (this *Server) HandleReset(c *web.Context)               {}
func (this *Server) HandleProfilePage(c *web.Context)         {}
func (this *Server) HandleProfile(c *web.Context)             {}
func (this *Server) HandleOpenidConfiguration(c *web.Context) {}
