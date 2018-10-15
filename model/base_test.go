package model

import (
	"fmt"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestGetInfo(t *testing.T) {
	var order COrder
	err := GetInfo(order, nil)
	fmt.Println(err)
}

func TestSaveOrUpdate(t *testing.T) {
	u := &User{
		Name:  "姜建康",
		Phone: "18610773065",
		Pwd:   "11111111",
		Role:  0,
	}
	v := map[string]interface{}{"Phone": "18610773066"}
	SaveOrUpdate(u, v, nil)
	target := &User{}
	err := GetInfo(target, v)
	assert.Equal(t, nil, err)
	assert.Equal(t, "18610773066", target.Phone)
}
