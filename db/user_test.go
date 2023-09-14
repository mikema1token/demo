package db

import (
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDBUserCRUD(t *testing.T) {
	user := UserModel{
		Model: gorm.Model{
			ID: 1,
		},
		Name: "name",
	}
	err := CreateUser(user)
	assert.Nil(t, err)
	userById, err := GetUserById(int(user.ID))
	assert.Nil(t, err)
	assert.Equal(t, userById, user)

	userList, err := GetUserList()
	assert.Nil(t, err)
	assert.True(t, len(userList) > 0)

	err = DeleteUserById(int(user.ID))
	assert.Nil(t, err)

	_, err = GetUserById(int(user.ID))
	assert.True(t, gorm.IsRecordNotFoundError(err))
}
