package db_test

import (
	"github.com/wei840222/blog/db"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	db.TruncateAllTable()
	user, err := db.CreateUser(&db.User{
		Name: "TestCreateUser",
	})

	assert.Nil(t, err)
	assert.Equal(t, "TestCreateUser", user.Name)
}

func TestGetUser(t *testing.T) {
	db.TruncateAllTable()
	userToGet, err := db.CreateUser(&db.User{
		Name: "TestGetUser",
	})

	assert.Nil(t, err)

	user, err := db.GetUser(userToGet.ID)
	assert.Nil(t, err)
	assert.Equal(t, "TestGetUser", user.Name)
}
