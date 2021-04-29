package db

import (
	"database/sql"
	"strings"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"github.com/phaesoo/keybox/internal/models"
	"github.com/stretchr/testify/suite"
)

func TestKeyAuth(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip integration tests")
	}
	suite.Run(t, new(AuthKeyTestSuite))
}

type AuthKeyTestSuite struct {
	TestSuite
	db *DB
}

func (ts *AuthKeyTestSuite) SetupSuite() {
	ts.TestSuite.SetupSuite()
	ts.db = NewDB(ts.Conn, strings.Replace(uuid.NewString(), "-", "", -1))
}

func (ts *AuthKeyTestSuite) Test_SetAuthKey() {
	var authKey models.AuthKey
	ts.NoError(faker.FakeData(&authKey))

	err := ts.db.SetAuthKey(authKey)
	ts.NoError(err)
}

func (ts *AuthKeyTestSuite) Test_AuthKey() {
	var authKey models.AuthKey
	ts.NoError(faker.FakeData(&authKey))

	err := ts.db.SetAuthKey(authKey)
	ts.NoError(err)

	ts.Run("It returns expected object", func() {
		res, err := ts.db.AuthKey(authKey.UserID, authKey.KeyID)
		ts.NoError(err)
		ts.True(authKey.Equal(res))
	})
	ts.Run("It returns error with unknown access key", func() {
		res, err := ts.db.AuthKey("user-1", "key-1")
		ts.Error(err)
		ts.EqualValues(models.AuthKey{}, res)
	})
}

func (ts *AuthKeyTestSuite) Test_DeleteAuthKey() {
	var authKey models.AuthKey
	ts.NoError(faker.FakeData(&authKey))

	err := ts.db.SetAuthKey(authKey)
	ts.NoError(err)

	_, err = ts.db.AuthKey(authKey.UserID, authKey.KeyID)
	ts.NoError(err)

	err = ts.db.DeleteAuthKey(authKey.UserID, authKey.KeyID)
	ts.NoError(err)

	_, err = ts.db.AuthKey(authKey.UserID, authKey.KeyID)
	ts.Error(sql.ErrNoRows)
}
