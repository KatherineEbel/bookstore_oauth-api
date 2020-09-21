package accessToken

import (
	"time"
)

const (
	expTime = 24
)

type AccessToken struct {
	Token    string `json:"token"`
	UserId   int64  `json:"userId"`
	ClientId int64  `json:"clientId"`
	Expires  int64  `json:"expires"`
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expTime * time.Hour).Unix(),
	}
}

func (t AccessToken) Expired() bool {
	return time.Unix(t.Expires, 0).Before(time.Now().UTC())
}
