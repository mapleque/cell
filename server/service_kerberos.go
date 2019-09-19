package server

import (
	"database/sql"
	"encoding/json"
	"time"
)

// Kerberos Kerberos is a service implement kerberos protocal.
//
//
// As an example, this project has using kerberos for user login.
type Kerberos struct {
	db  DB
	log Logger

	tgsSecretKey string
	appSecretKey string
}

// KerberosAuthResult KerberosAuthResult is the result of kerberos auth.
type KerberosAuthResult struct {
	EncCTSK string `json:"ctsk"`
	EncTGT  string `json:"tgt"`
}

// KerberosGrantResult KerberosGrantResult is the result of kerberos grant.
type KerberosGrantResult struct {
	EncCSSK string `json:"cssk"`
	EncST   string `json:"st"`
}

// NewKerberos NewKerberos create a Kerberos service with settings.
func NewKerberos(db DB, log Logger, tgsSecretKey, appSecretKey string) *Kerberos {
	return &Kerberos{
		db,
		log,
		tgsSecretKey,
		appSecretKey,
	}
}

type kerberosTGT struct {
	CTSK     string `json:"ctsk"`
	Username string `json:"username"`
	Expired  int64  `json:"expired"`
}

// Auth Auth implement the process of kerberos auth.
func (k *Kerberos) Auth(username string) (*KerberosAuthResult, error) {
	// find password from db
	var password string
	err := k.db.QueryRowContext(
		dbCtx,
		"SELECT `password` FROM `user` WHERE `username`=? LIMIT 1",
		username,
	).Scan(&password)
	switch {
	case err == sql.ErrNoRows:
		return nil, errUserNotExist
	case err != nil:
		panic(err)
	}

	tgt := &kerberosTGT{
		CTSK:     RandToken(),
		Username: username,
		Expired:  time.Now().Add(5 * time.Minute).Unix(),
	}

	encCTSK := k.encrypt(tgt.CTSK, password)
	encTGT := k.encrypt(tgt, k.tgsSecretKey)
	res := &KerberosAuthResult{
		encCTSK,
		encTGT,
	}
	return res, nil
}

type kerberosAuthenticator struct {
	Username  string `json:"username"`
	Timestamp int64  `json:"timestamp"`
}

type kerberosServiceTicket struct {
	CSSK     string `json:"cssk"`
	Username string `json:"username"`
	Expired  int64  `json:"expired"`
}

// Grant Grant implement the process of kerberos grant.
func (k *Kerberos) Grant(encTGT, appID, encAuthenticator string) (*KerberosGrantResult, error) {
	tgt := &kerberosTGT{}
	if err := k.decrypt(encTGT, k.tgsSecretKey, tgt); err != nil {
		return nil, errTGTInvalid
	}
	if tgt.Expired < time.Now().Unix() {
		return nil, errTGTInvalid
	}
	authenticator := &kerberosAuthenticator{}
	if err := k.decrypt(encAuthenticator, tgt.CTSK, authenticator); err != nil {
		return nil, errAuthenticatorInvalid
	}

	var appSecret string
	err := k.db.QueryRowContext(
		dbCtx,
		"SELECT `secret` FROM `app` WHERE `app_id`=? LIMIT 1",
		appID,
	).Scan(&appSecret)
	switch {
	case err == sql.ErrNoRows:
		return nil, errAppNotExist
	case err != nil:
		panic(err)
	}

	st := &kerberosServiceTicket{
		Username: authenticator.Username,
		Expired:  time.Now().Add(2 * time.Hour).Unix(),
	}

	encST := k.encrypt(st, appSecret)
	encCSSK := k.encrypt(RandToken(), tgt.CTSK)

	res := &KerberosGrantResult{
		encST,
		encCSSK,
	}
	return res, nil
}

// Login Login usually using in application service,
// for validating the kerberos service ticket.
func (k *Kerberos) Login(encST, encAuthenticator string) (string, error) {
	st := &kerberosServiceTicket{}
	if err := k.decrypt(encST, k.appSecretKey, st); err != nil {
		return "", errSTInvalid
	}
	if st.Expired < time.Now().Unix() {
		return "", errSTInvalid
	}

	authenticator := &kerberosAuthenticator{}
	if err := k.decrypt(encAuthenticator, st.CSSK, authenticator); err != nil {
		return "", errAuthenticatorInvalid
	}

	if authenticator.Username != st.Username ||
		time.Now().Unix()-authenticator.Timestamp > int64(300) {
		return "", errAuthenticatorInvalid
	}

	return authenticator.Username, nil
}

func (k *Kerberos) encrypt(tar interface{}, key string) string {
	data, _ := json.Marshal(tar)
	cipher, _ := EncryptAesEcbBase64(data, key)
	return cipher
}

func (k *Kerberos) decrypt(cipher, key string, tar interface{}) error {
	data, err := DecryptAesEcbBase64(cipher, key)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, tar); err != nil {
		return err
	}
	return nil
}
