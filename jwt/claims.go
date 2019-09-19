package jwt

// Claims is an interface of claims
type Claims interface {
	Valid() error
	SetIssuer(iss string)
	SetSubject(sub string)
	SetAudience(aud string)
	SetExpirationTime(exp int64)
	SetNotBefore(nbf int64)
	SetIssuedAt(iat int64)
	SetJWTID(jti string)
	GetIssuer() string
	GetSubject() string
	GetAudience() string
	GetExpirationTime() int64
	GetNotBefore() int64
	GetIssuedAt() int64
	GetJWTID() string
}

// StdClaims implement Claims interface
// which include all standard properties.
//
// Extend to add more public value if you need.
type StdClaims struct {
	Issuer         string `json:"iss,omitempty"`
	Subject        string `json:"sub,omitempty"`
	Audience       string `json:"aud,omitempty"`
	ExpirationTime int64  `json:"exp,omitempty"`
	NotBefore      int64  `json:"nbf,omitempty"`
	IssuedAt       int64  `json:"iat,omitempty"`
	JWTID          string `json:"jti,omitempty"`
}

// NewStdClaims returns a StdClaims entity with default value
func NewStdClaims() Claims {
	return &StdClaims{}
}

// Valid Valid check the claims is valid.
//
// TODO implement valid method
func (c *StdClaims) Valid() error {
	return nil
}

// SetIssuer SetIssuer set issuer = iss.
func (c *StdClaims) SetIssuer(iss string) {
	c.Issuer = iss
}

// SetSubject SetSubject set subject = sub.
func (c *StdClaims) SetSubject(sub string) {
	c.Subject = sub
}

// SetAudience SetAudience set audience = aud.
func (c *StdClaims) SetAudience(aud string) {
	c.Audience = aud
}

// SetExpirationTime SetExpirationTime set expireation time = exp.
func (c *StdClaims) SetExpirationTime(exp int64) {
	c.ExpirationTime = exp
}

// SetNotBefore SetNotBefore set not before = nbf.
func (c *StdClaims) SetNotBefore(nbf int64) {
	c.NotBefore = nbf
}

// SetIssuedAt SetIssuedAt set issued at = iat.
func (c *StdClaims) SetIssuedAt(iat int64) {
	c.IssuedAt = iat
}

// SetJWTID SetJWTID set jwtid = jti.
func (c *StdClaims) SetJWTID(jti string) {
	c.JWTID = jti
}

// GetIssuer GetIssuer get issuer.
func (c *StdClaims) GetIssuer() string {
	return c.Issuer
}

// GetSubject GetSubject get subject.
func (c *StdClaims) GetSubject() string {
	return c.Subject
}

// GetAudience GetAudience get audience.
func (c *StdClaims) GetAudience() string {
	return c.Audience
}

// GetExpirationTime GetExpirationTime get expiration time
func (c *StdClaims) GetExpirationTime() int64 {
	return c.ExpirationTime
}

// GetNotBefore GetNotBefore get not before
func (c *StdClaims) GetNotBefore() int64 {
	return c.NotBefore
}

// GetIssuedAt GetIssuedAt get issued at
func (c *StdClaims) GetIssuedAt() int64 {
	return c.IssuedAt
}

// GetJWTID GetJWTID get jwtid
func (c *StdClaims) GetJWTID() string {
	return c.JWTID
}
