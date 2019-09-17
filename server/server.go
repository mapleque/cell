package server

// Server Http server
type Server struct {
	addr       string
	staticPath string
	db         DB
	log        Logger

	httpServer *http.Server
	handler    *http.ServeMux
}

// NewServer Create a Server entity
func NewServer(addr, staticPath string, db DB, log Logger) *Server {
	handler := http.DefaultServeMux
	s := &Server{
		addr:       addr,
		staticPath: staticPath,
		db:         db,
		log:        log,
		httpServer: &http.Server{
			Addr:           addr,
			Handler:        s.handler,
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
func (s *Server) Run() {
	s.log.Printf("[Info] http server listen on %s\n", s.addr)
	s.log.Fatalln("[Fatal]", s.httpServer.ListenAndServe())
}
