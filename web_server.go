package cell

import (
	"github.com/mapleque/kelp/web"
)

const (
	ERROR_INVALID_REQUEST           = "invalid_request"
	ERROR_UNAUTHORIZED_CLIENT       = "unauthorized_client"
	ERROR_ACCESS_DENIED             = "access_denied"
	ERROR_UNSUPPORTED_RESPONSE_TYPE = "unsupported_response_type"
	ERROR_INVALID_SCOPE             = "invalid_scope"
	ERROR_SERVER_ERROR              = "server_error"
	ERROR_TEMPORARILY_UNAVAILABLE   = "temporarily_unavailable"
)

type Server struct {
	ds IDataService
}

func NewServer(ds IDataService) *Server {
	return &Server{ds}
}

func (this *Server) Serve(host string) {
	server := web.New(host)
	// basic
	server.GET("/", this.HandleHomePage)
	server.GET("/login", this.HandleLoginPage)
	server.POST("/login", this.HandleLogin)
	server.GET("/register", this.HandleRegisterPage)
	server.POST("/register", this.HandleRegister)
	server.POST("/logout", this.HandleLogout)
	server.GET("/forgot", this.HandleForgotPage)
	server.POST("/forgot", this.HandleForgot)
	server.GET("/reset", this.HandleResetPage)
	server.POST("/reset", this.HandleReset)
	server.GET("/profile", this.HandleProfilePage)
	server.POST("/profile", this.HandleProfile)
	// for openid
	server.GET("/.well-known/openid-configuration", this.HandleOpenidConfiguration)
	// for oauth2
	server.GET("/auth", this.HandleAuthPage)
	server.POST("/auth", this.HandleAuth)
	server.POST("/token", this.HandleToken)
	server.Run()
}
