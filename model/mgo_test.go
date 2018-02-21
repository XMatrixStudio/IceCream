package model

import (
	"fmt"
	"testing"
)

func Test_Division_Init(t *testing.T) {
	err := InitMongo()
	if err != nil {
		t.Error(err)
	}
}

func Test_Division_User(t *testing.T) {
	id, err := AddUser()
	fmt.Println(id, err)
}
