package cache

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"

	"github.com/phaesoo/keybox/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestAuthKey(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip integration tests")
	}
	suite.Run(t, new(AuthKeyTestSuite))
}

func Test_authKeyHash(t *testing.T) {
	assert := assert.New(t)

	var accessKey string
	assert.NoError(faker.FakeData(&accessKey))

	res := generateKey(accessKey)
	assert.Equal(res, fmt.Sprintf(authKeyHashPrefix, accessKey))
}

type AuthKeyTestSuite struct {
	TestSuite
	cache *Cache
}

func (ts *AuthKeyTestSuite) SetupSuite() {
	ts.TestSuite.SetupSuite()
	ts.cache = NewCache(ts.Pool, strings.Replace(uuid.NewString(), "-", "", -1))
}

func (ts *AuthKeyTestSuite) Test_SetAuthKey() {
	var authKey models.AuthKey
	ts.NoError(faker.FakeData(&authKey))

	err := ts.cache.SetAuthKey(authKey, 1)
	ts.NoError(err)

	// Should be expired after ttl
	time.Sleep(2 * time.Second)
	_, err = ts.cache.AuthKey(authKey.KeyID)
	ts.Error(err, ErrNotFound)
}

func (ts *AuthKeyTestSuite) Test_GetAuthKey() {
	var authKey models.AuthKey
	ts.NoError(faker.FakeData(&authKey))

	err := ts.cache.SetAuthKey(authKey, 1)
	ts.NoError(err)

	res, err := ts.cache.AuthKey(authKey.KeyID)
	ts.NoError(err)
	ts.EqualValues(authKey, res)
}
