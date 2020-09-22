package accessToken

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstants(t *testing.T) {
	assert.EqualValues(t, 24, expTime, "expTime should be 24")
}
func TestGetNewAccessToken_NotExpired(t *testing.T) {
	tok := GetNewAccessToken()
	assert.False(t, tok.Expired(), "expected new token not to be expired")
}

func TestGetNewAccessToken_TokenStringIsEmpty(t *testing.T) {
	tok := GetNewAccessToken()
	assert.Empty(t, tok.AccessToken, "new access token should not have token id")
}

func TestGetNewAccessToken_UserIdIsNil(t *testing.T) {
	tok := GetNewAccessToken()
	assert.Empty(t, tok.UserId, "new access token should not have user id")
}

func TestGetNewAccessToken_EmptyTokenExpired(t *testing.T) {
	tok := AccessToken{}
	assert.True(t, tok.Expired(), "expected empty token to be expired by default")
}

func TestGetNewAccessToken_FutureNotExpired(t *testing.T) {
	tok := AccessToken{Expires: time.Now().UTC().Add(time.Hour * 3).Unix()}
	assert.False(t, tok.Expired(), "expected token to NOT be expired")
}
