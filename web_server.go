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
	ERROR_INVALID_CLIENT            = "invalid_client"
	ERROR_INVALID_GRANT             = "invalid_grant"
	ERROR_UNSUPPORTED_GRANT_TYPE    = "unsupported_grant_type"
)

type Server struct {
	ds           IDataService
	templatePath string
}

func NewServer(ds IDataService, templatePath string) *Server {
	return &Server{
		ds:           ds,
		templatePath: templatePath,
	}
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
	server.POST("/userinfo", this.HandleUserinfo)
	server.Run()
}

func (this *Server) tpl(c *web.Context, file string, data interface{}) {
	t := template.Must(template.ParseFiles(this.templatePath + file))
	t.Execute(c.ResponseWriter, data)
}

type Tips struct {
	Status Status
	Tips   string
}

type Status string

const (
	SUCCESS Status = "success"
	WARNING        = "warning"
	ERROR          = "error"
)
