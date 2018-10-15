package service

import (
	"testing"

	"house/controller/form"

	"github.com/magiconair/properties/assert"
)

func TestLoginOrCreateService(t *testing.T) {
	form := form.Login{
		User:     "18610773065",
		Password: "123456",
	}
	u, e := LoginOrCreateService(&form)
	assert.Equal(t, form.User, u.Phone)
	assert.Equal(t, e, nil)
}

func TestUpPwdService(t *testing.T) {
	form := form.Login{
		User:     "18610773065",
		Password: "123458",
	}

	u, e := UpPwdService(&form)
	assert.Equal(t, e, nil)
	assert.Equal(t, u.Pwd, form.Password)
}
