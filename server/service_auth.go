package server

import (
	"database/sql"
	"net/http"
	"time"
)

// Auth a http auth service
//
// Http Auth Type
// - With Cookie (session key)
// - With Bearer Token (session key)
// - with Bearer Token (jwt token)
type Auth struct {
	db   DB
	log  Logger
	mail *Mail
}

// NewAuth create Auth entity
func NewAuth(db DB, log Logger, mail *Mail) *Auth {
	return &Auth{db, log, mail}
}

// AuthUser authed user entity
type AuthUser struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	LastLogin string `json:"last_login"`
}

const (
	authCookieName = "cell_session"
	appID          = "cell"
)

// GetUser get user info from auth info
//
// Here we use cookie (session key) first
func (auth *Auth) GetUser(r *http.Request) (*AuthUser, error) {
	cookie, err := r.Cookie(authCookieName)
	if err != nil {
		return nil, errNoAuthToken
	}

	token := cookie.Value
	user := &AuthUser{}
	var (
		expired time.Time
	)
	err = auth.db.QueryRowContext(
		dbCtx(),
		"SELECT `id`, `username`, `last_login`, `expired_at` "+
			"FROM `user_auth` "+
			"WHERE `app_id` = ? AND `token` = ? AND expired_at > NOW() "+
			"LIMIT 1",
		appID,
		token,
	).Scan(&user.ID, &user.Username, &user.LastLogin, &expired)
	switch {
	case err == sql.ErrNoRows:
		return nil, errAuthTokenInvalid
	case err != nil:
		panic(err)
	}

	if expired.Before(time.Now().Add(30 * time.Minute)) {
		_, _ = auth.db.ExecContext(
			dbCtx(),
			"UPDATE `user_auth` "+
				"SET `expired_at`=DATE_ADD(NOW(),INTERVAL 1 HOUR) "+
				"WHERE `id`=? LIMIT 1",
			user.ID,
		)
	}
	return user, nil
}

// Login just login with username
//
// No need to check password here
func (auth *Auth) Login(
	w http.ResponseWriter,
	r *http.Request,
	username string,
	remember bool,
) string {

	if _, err := auth.db.ExecContext(
		dbCtx(),
		"INSERT IGNORE INTO `user_auth` "+
			"(`app_id`, `username`) "+
			"VALUES(?,?) ",
		appID,
		username,
	); err != nil {
		panic(err)
	}

	// The token is 64 bytes random string,
	// which an optimistic lock unique was using in mysql.
	// So When conflict, it will panic and this request will be failed.
	token := RandToken()
	if _, err := auth.db.ExecContext(
		dbCtx(),
		"UPDATE `user_auth` "+
			"SET `token`=?,`ip`=?,"+
			"`expired_at`=DATE_ADD(NOW(), INTERVAL 1 HOUR),"+
			"`last_login`=NOW() "+
			"WHERE `app_id`=? AND `username`=? LIMIT 1",
		token,
		r.RemoteAddr,
		appID,
		username,
	); err != nil {
		panic(err)
	}

	maxAge := 0
	if remember {
		maxAge = 86400
	}

	// set cookie
	cookie := &http.Cookie{
		Name:   authCookieName,
		Value:  token,
		Path:   "/",
		MaxAge: maxAge,
	}
	http.SetCookie(w, cookie)
	// return token
	return token
}

// Logout logout all ticket
func (auth *Auth) Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(authCookieName)
	if err != nil {
		// no token, do nothing
		return
	}

	token := cookie.Value
	if _, err := auth.db.ExecContext(
		dbCtx(),
		"UPDATE `user_auth` SET `token`=NULL,`expired_at`=NOW() "+
			"WHERE `app_id`=? AND `token`=? LIMIT 1",
		appID,
		token,
	); err != nil {
		panic(err)
	}
	cookie.MaxAge = 0
	http.SetCookie(w, cookie)
}

// CheckUserExist CheckUserExist check the username has been register.
func (auth *Auth) CheckUserExist(username string) bool {
	var id int64
	err := auth.db.QueryRowContext(
		dbCtx(),
		"SELECT `id` FROM `user` WHERE `username` = ? LIMIT 1",
		username,
	).Scan(&id)
	switch {
	case err == sql.ErrNoRows:
		return false
	case err != nil:
		panic(err)
	}
	return true
}

// SendCaptcha cache captcha and send
func (auth *Auth) SendCaptcha(username, subject, template string) error {
	captcha := RandMd5()[0:4]
	if _, err := auth.db.ExecContext(
		dbCtx(),
		"INSERT INTO `user_captcha` "+
			"(`username`, `captcha`, `expired_at`) "+
			"VALUES (?,?,DATE_ADD(NOW(), INTERVAL 5 MINUTE)) "+
			"ON DUPLICATE KEY UPDATE "+
			"`captcha`=?,`expired_at`=DATE_ADD(NOW(), INTERVAL 5 MINUTE)",
		username,
		captcha,
		captcha,
	); err != nil {
		panic(err)
	}
	auth.log.Printf("[Info] user %s captcha %s\n", username, captcha)
	auth.mail.Send(
		[]string{username},
		subject,
		template,
		struct {
			Email   string
			Captcha string
		}{
			username,
			captcha,
		},
	)
	return nil
}

// CheckCaptcha check captcha is valid
func (auth *Auth) CheckCaptcha(username, captcha string) error {
	var id int64
	err := auth.db.QueryRowContext(
		dbCtx(),
		"SELECT `id` FROM `user_captcha` "+
			"WHERE `username`=? AND `captcha`=? AND `expired_at`>NOW() "+
			"LIMIT 1",
		username,
		captcha,
	).Scan(&id)
	switch {
	case err == sql.ErrNoRows:
		return errCaptchaInvalid
	case err != nil:
		panic(err)
	}

	if _, err := auth.db.ExecContext(
		dbCtx(),
		"DELETE FROM `user_captcha` WHERE `id` = ? LIMIT 1",
		id,
	); err != nil {
		auth.log.Printf("[Error] delete captcha failed, cause of: %v\n", err)
	}

	return nil
}

// Register register user
func (auth *Auth) Register(username, password string) error {
	var id int64
	err := auth.db.QueryRowContext(
		dbCtx(),
		"SELECT `id` FROM `user` WHERE `username`=? LIMIT 1",
		username,
	).Scan(&id)
	switch {
	case err == nil:
		return errUserExist
	case err != sql.ErrNoRows:
		panic(err)
	}
	if _, err := auth.db.ExecContext(
		dbCtx(),
		"INSERT INTO `user` (`username`,`password`) VALUES (?,?)",
		username,
		password,
	); err != nil {
		panic(err)
	}
	return nil
}

// ResetPassword reset user password
func (auth *Auth) ResetPassword(username, password string) error {
	if _, err := auth.db.ExecContext(
		dbCtx(),
		"UPDATE `user` SET `password`=?, `update_at`=NOW() "+
			"WHERE `username`=? LIMIT 1",
		password,
		username,
	); err != nil {
		panic(err)
	}

	return nil
}
