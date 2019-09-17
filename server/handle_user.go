package server

type userLoginParam struct {
	EncST            string `json:"enc_st"`
	EncAuthenticator string `json:"enc_authenticator"`
}

func (s *Server) handleUserLogin(w http.ResponseWriter, r *http.Request) {
	// Here we use kerberos token in login
	in := &userLoginParam{}
	err := bind(r, in)
	if err != nil {
		resp(w, 10001, err)
		return
	}

	username, err := s.kerberos.Login(in.EncST, in.EncAuthenticator)
	if err != nil {
		resp(w, 10001, err)
		return
	}

	token := s.auth.Login(w, username)
	resp(w, 0, token)
}

func (s *Server) handleUserLogout(w http.ResponseWriter, r *http.Request) {
	s.auth.Logout(w, r)
	resp(w, 0, nil)
}

type userRegisterParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Captcha  string `json:"captcha"`
}

func (s *Server) handleUserCaptcha(w http.ResponseWriter, r *http.Request) {
	in := &userRegisterParam{}
	err := bind(r, in)
	if err != nil {
		resp(w, 10001, err)
		return
	}

	err := checkEmailAddress(in.Username)
	if err != nil {
		resp(w, 10001, err)
		return
	}

	err := s.auth.SendCaptcha(in.Username)
	resp(w, 0, nil)
}

func (s *Server) handleUserRegister(w http.ResponseWriter, r *http.Request) {
	in := &userRegisterParam{}
	err := bind(r, in)
	if err != nil {
		resp(w, 10001, err)
		return
	}

	err := s.auth.CheckCaptcha(in.Username, in.Captcha)
	if err != nil {
		resp(w, 10001, err)
		return
	}

	err := s.auth.Register(in.Username, in.Password)
	if err != nil {
		resp(w, 11002, err)
		return
	}
	resp(w, 0, nil)
}

func (s *Server) handleUserForgot(w http.ResponseWriter, r *http.Request) {
	in := &userRegisterParam{}
	err := bind(r, in)
	if err != nil {
		resp(w, 10001, err)
		return
	}

	err := checkEmailAddress(in.Username)
	if err != nil {
		resp(w, 10001, err)
		return
	}

	err := s.auth.SendCaptcha(in.Username)
	resp(w, 0, nil)
}

func (s *Server) handleUserReset(w http.ResponseWriter, r *http.Request) {
	in := &userRegisterParam{}
	err := bind(r, in)
	if err != nil {
		resp(w, 10001, err)
		return
	}

	err := s.auth.CheckCaptcha(in.Username, in.Captcha)
	if err != nil {
		resp(w, 10001, err)
		return
	}

	err := s.auth.ResetPassword(in.Username, in.Password)
	if err != nil {
		resp(w, 11002, err)
		return
	}
	resp(w, 0, nil)
}

func (s *Server) handleUserProfile(w http.ResponseWriter, r *http.Request) {
	// TODO
	// user info
	// grant info
	// app info
}
