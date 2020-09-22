package accessToken

import (
	"strings"
	"time"

	"github.com/KatherineEbel/bookstore_oauth-api/src/utils/errors"
)

const (
	expTime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func (t *AccessToken) Validate() *errors.RestError {
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

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expTime * time.Hour).Unix(),
	}
}

func (t *AccessToken) Expired() bool {
	return time.Unix(t.Expires, 0).Before(time.Now().UTC())
}
