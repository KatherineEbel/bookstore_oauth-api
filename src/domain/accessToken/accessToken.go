package accessToken

import (
	"fmt"
	"strings"
	"time"

	"github.com/KatherineEbel/bookstore_oauth-api/src/utils/crypto"
	"github.com/KatherineEbel/bookstore_oauth-api/src/utils/errors"
)

const (
	expTime                    = 24
	grantTypePassword          = "password"
	grantTypeClientCredentials = "client_credentials"
)

type Token struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id,omitempty"`
	Expires     int64  `json:"expires"`
}

type Request struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`

	// grant_type password
	UserName string `json:"user_name"`
	Password string `json:"password"`

	// grant_type client_credentials
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (r Request) Validate() *errors.RestError {
	switch r.GrantType {
	case grantTypePassword:
	case grantTypeClientCredentials:
	default:
		return errors.NewBadRequestError("invalid grant type")
	}
	return nil
}

func (t *Token) Validate() *errors.RestError {
	var msgs []string
	t.AccessToken = strings.TrimSpace(t.AccessToken)
	if t.AccessToken == "" {
		msgs = append(msgs, "invalid token id")
	}
	if t.UserId <= 0 {
		msgs = append(msgs, "invalid user id")
	}
	if t.ClientId <= 0 {
		msgs = append(msgs, "invalid client id")
	}
	if t.Expired() {
		msgs = append(msgs, "token expired")
	}
	if len(msgs) > 0 {
		str := strings.Join(msgs, ", ")
		return errors.NewBadRequestError(str)
	}
	return nil
}

func GetNewAccessToken(id int64) Token {
	return Token{
		UserId:  id,
		Expires: time.Now().UTC().Add(expTime * time.Hour).Unix(),
	}
}

func (t *Token) Expired() bool {
	return time.Unix(t.Expires, 0).Before(time.Now().UTC())
}

func (t *Token) Generate() {
	t.AccessToken = crypto.GetMd5(fmt.Sprintf("ke-%d-%d-ran", t.UserId, t.Expires))
}
