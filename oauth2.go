package cell

import (
	"github.com/mapleque/kelp/web"
)

type Oauth2 struct {
}

func NewOauth2() *Oauth2 {}

func (this *Oauth2) Serve(host) {
	server := web.New(host)
	// basic
	server.GET("/", HandleHomePage)
	server.GET("/login", HandleLoginPage)
	server.POST("/login", HandleLogin)
	server.GET("/register", HandleRegisterPage)
	server.POST("/register", HandleRegister)
	server.POST("/logout", HandleLogout)
	server.GET("/forgot", HandleForgotPage)
	server.POST("/forgot", HandleForgot)
	server.GET("/reset", HandleResetPage)
	server.POST("/reset", HandleReset)
	server.GET("/profile", HandleProfilePage)
	server.POST("/profile", HandleProfile)
	// for openid
	http.HandleFunc("/.well-known/openid-configuration", HandleOpenidConfiguration)
	// for oauth2
	server.GET("/auth", HandleAuthPage)
	server.POST("/auth", HandleAuth)
	server.POST("/token", HandleToken)
}
