package server

import (
	"net/http"
	"time"
)

// Server Http server
type Server struct {
	addr       string
	staticPath string
	db         DB
	log        Logger
	oidc       *Oidc
	auth       *Auth
	kerberos   *Kerberos

	httpServer *http.Server
	handler    *http.ServeMux
}

// NewServer Create a Server entity
func NewServer(
	addr, staticPath string,
	db DB,
	log Logger,
	oidc *Oidc,
	auth *Auth,
	kerberos *Kerberos,
) *Server {
	handler := http.DefaultServeMux
	s := &Server{
		addr:       addr,
		staticPath: staticPath,
		db:         db,
		log:        log,
		oidc:       oidc,
		auth:       auth,
		kerberos:   kerberos,

		httpServer: &http.Server{
			Addr:           addr,
			Handler:        handler,
			ReadTimeout:    5 * time.Second,
			WriteTimeout:   5 * time.Second,
			MaxHeaderBytes: http.DefaultMaxHeaderBytes,
		},
		handler: handler,
	}

	s.initRouter()
	return s
}

// Run Run http server
func (s *Server) Run(certFile, keyFile string) {
	s.log.Printf("[Info] http server listen on %s\n", s.addr)
	if certFile != "" && keyFile != "" {
		s.log.Fatalln("[Fatal]", s.httpServer.ListenAndServeTLS(certFile, keyFile))
	} else {
		s.log.Fatalln("[Fatal]", s.httpServer.ListenAndServe())
	}
}
