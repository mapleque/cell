package server

import "net/http"

type userLoginParam struct {
	EncST            string `json:"st"`
	EncAuthenticator string `json:"authenticator"`
	Remember         bool   `json:"remember"`
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

	token := s.auth.Login(w, r, username, in.Remember)
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
	if err := bind(r, in); err != nil {
		resp(w, 10001, err)
		return
	}

	if err := checkEmailAddress(in.Username); err != nil {
		resp(w, 10001, err)
		return
	}

	if s.auth.CheckUserExist(in.Username) {
		resp(w, 10001, "该账号已注册")
		return
	}

	if err := s.auth.SendCaptcha(in.Username, "Cell: 您注册的验证码", "register.html"); err != nil {
		resp(w, 11001, err)
		return
	}
	resp(w, 0, nil)
}

func (s *Server) handleUserRegister(w http.ResponseWriter, r *http.Request) {
	in := &userRegisterParam{}
	if err := bind(r, in); err != nil {
		resp(w, 10001, err)
		return
	}

	if err := s.auth.CheckCaptcha(in.Username, in.Captcha); err != nil {
		resp(w, 10001, err)
		return
	}

	if err := s.auth.Register(in.Username, in.Password); err != nil {
		resp(w, 11002, err)
		return
	}
	resp(w, 0, nil)
}

func (s *Server) handleUserForgot(w http.ResponseWriter, r *http.Request) {
	in := &userRegisterParam{}
	if err := bind(r, in); err != nil {
		resp(w, 10001, err)
		return
	}

	if err := checkEmailAddress(in.Username); err != nil {
		resp(w, 10001, err)
		return
	}

	if !s.auth.CheckUserExist(in.Username) {
		resp(w, 10001, "该账号尚未注册")
		return
	}

	if err := s.auth.SendCaptcha(in.Username, "Cell: 您正在申请重置密码", "forgot.html"); err != nil {
		resp(w, 11001, err)
		return
	}
	resp(w, 0, nil)
}

func (s *Server) handleUserReset(w http.ResponseWriter, r *http.Request) {
	in := &userRegisterParam{}
	if err := bind(r, in); err != nil {
		resp(w, 10001, err)
		return
	}

	if err := s.auth.CheckCaptcha(in.Username, in.Captcha); err != nil {
		resp(w, 10001, err)
		return
	}

	if err := s.auth.ResetPassword(in.Username, in.Password); err != nil {
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

func checkEmailAddress(email string) error {
	// TODO
	return nil
}
