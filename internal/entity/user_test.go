package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Paulin", "pr@m.com", "senha!123")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.Equal(t, "Paulin", user.Name)
	assert.Equal(t, "pr@m.com", user.Email)
}

func TestCheckPassword(t *testing.T) {
	user, err := NewUser("Paulin", "pr@m.com", "senha!123")
	assert.Nil(t, err)
	assert.True(t, user.CheckPassword("senha!123"))
	assert.False(t, user.CheckPassword("wrongpassword"))
	assert.NotEqual(t, "senha!123", user.Password)

}
