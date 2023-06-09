package access_token

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetNewAccessTokenConstants(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "expiration time should be 24h")
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()

	assert.False(t, at.IsExpired(), "brand new access token should not be expired")
	assert.EqualValues(t, at.AccessToken, "", "brand new access token should not have defined access token id")
	assert.EqualValues(t, at.UserId, 0, "brand new access token should not have defined user id")
}

func TestGetNewAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.IsExpired(), "empty access token should be expired by default")

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "access token expiring 3 hours from now should NOT be expired")
}
