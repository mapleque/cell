package server

// Auth a http auth service
//
// Http Auth Type
// - With Cookie (session key)
// - With Bearer Token (session key)
// - with Bearer Token (jwt token)
type Auth struct{}

// NewAuth create Auth entity
func NewAuth() *Auth {
	return &Auth{}
}

// AuthUser authed user entity
type AuthUser struct {
	Username  string `json:"username"`
	LastLogin string `json:"last_login"`
}

// GetUser get user info from auth info
//
// Here we use cookie (session key) first
func (auth *Auth) GetUser(r *http.Request) *AuthUser {
	// TODO
	return nil
}

// Login just login with username
//
// No need to check password here
func (auth *Auth) Login(w http.ResponseWriter, username string) string {
	// TODO set cookie
	// TODO return session token
	return ""
}

// Logout logout all ticket
func (auth *Auth) Logout(w http.RewsponseWriter, r *http.Request) {
	// TODO logout
}

// SendCaptcha cache captcha and send
func (auth *Auth) SendCaptcha(email string) error {
	// TODO
	return nil
}

// CheckCaptcha check captcha is valid
func (auth *Auth) CheckCaptcha(email, captcha string) error {
	// TODO
	return nil
}

// Register register user
func (auth *Auth) Register(username, password string) error {
	// TODO
	return nil
}

// ResetPassword reset user password
func (auth *Auth) ResetPassword(username, password string) error {
	// TODO
	return nil
}
