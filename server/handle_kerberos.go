package server

import "net/http"

type kerberosAuthParam struct {
	Username string `json:"username"`
}

func (s *Server) handleKerberosAuth(w http.ResponseWriter, r *http.Request) {
	in := &kerberosAuthParam{}
	err := bind(r, in)
	if err != nil {
		resp(w, 10001, err)
		return
	}

	ret, err := s.kerberos.Auth(in.Username)
	if err != nil {
		resp(w, 11001, err)
		return
	}

	resp(w, 0, ret)
}

type kerberosGrantParam struct {
	EncTGT               string `json:"tgt"`
	AppID                string `json:"app_id"`
	CTSKEncAuthenticator string `json:"authenticator"`
}

func (s *Server) handleKerberosGrant(w http.ResponseWriter, r *http.Request) {
	in := &kerberosGrantParam{}
	err := bind(r, in)
	if err != nil {
		resp(w, 10001, err)
		return
	}

	ret, err := s.kerberos.Grant(in.EncTGT, in.AppID, in.CTSKEncAuthenticator)
	if err != nil {
		resp(w, 11001, err)
		return
	}
	resp(w, 0, ret)
}
