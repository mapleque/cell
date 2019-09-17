package server

// Oidc oidc service
type Oidc struct {
	Keys []*jwt.RsaKeySet `json:"keys"`

	db  DB
	log Logger
}

// NewOidc create oidc entity
func NewOidc(db DB, log Logger) *Oidc {
	return &Oidc{
		Keys: []*jwt.RsaKeySet{},
		db:   db,
		log:  log,
	}
}

// AddKeyPair add a key pair
func (oidc *Oidc) AddKeyPair(keyID, publicKey, privateKey string) {
	ks, err := jwt.NewRsaKeySet(
		keyID,
		publicKey,
		privateKey,
		"RS256",
	)
	if err != nil {
		oidc.log.Error("add key pair error", err)
		return
	}
	this.Keys = append(this.Keys, ks)
}

// GetKeys get the jwks
//
// Here we use Oidc entity for jwks, cause of
// the property Keys is the only exported one.
func (oidc *Oidc) GetKeys() *Oidc {
	return oidc
}

// OidcClient the oidc client entity
//
// User should register the oidc client first.
//
// The oidc client data saved in db.
type OidcClient struct {
	ClientID    string `json:"client_id"`
	ClientDesc  string `json:"client_desc"`
	RedirectURI string `json:"redirect_uri"`
}

// FindClient find oidc client by client id
func (oidc *Oidc) FindClient(clientID string) (*OidcClient, bool) {
	// TODO
}

// OidcCode oidc code entity
type OidcCode struct {
	Code         string `json:"code"`
	State        string `json:"state"`
	IDToken      string `json:"id_token,omitempty"`
	AccessToken  string `json:"access_token,omitempty"`
	TokenType    string `json:"token_type,omitempty"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Openid       string `json:"openid,omitempty"`
}

// Auth build a auth code for user on the client
func (oidc *Oidc) Auth(
	user *AuthUser,
	responseType,
	clientID,
	redirectURI,
	scope,
	state string,
) (*OidcCode, error) {

	code := &OidcCode{}

	// TODO general code and save

	return code, nil
}

// CheckClient check the client id and secret match or not
func (oidc *Oidc) CheckClient(id, secret string) bool {
	// TODO
}

// OidcToken oidc token entity
type OidcToken struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Openid       string `json:"openid"`
}

// GrantToken grant the auth code a token
func (oidc *Oidc) GrantToken(code string) (*OidcToken, error) {
	// TODO find the code and user, build token
	return nil, nil
}

// RefreshToken refresh the token with refresh token
func (oidc *Oidc) RefreshToken(token string) (*OidcToken, error) {
	// TODO find the token and refresh
	return nil, nil
}

// OidcUser oidc userinfo entity
type OidcUser struct {
	Username string `json:"username"`
	Openid   string `json:"openid"`
}

// FindUser find user by access token
func (oidc *Oidc) FindUser(token string) (*OidcUser, error) {
	// TODO find userinfo by token
	return nil, nil
}
