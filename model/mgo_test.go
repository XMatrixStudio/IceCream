package model

import (
	"fmt"
	"testing"
)

func Test_Division_1(t *testing.T) {
	err := InitMongo()
	if err != nil {
		t.Error(err)
	}
}

func Test_Division_2(t *testing.T) {
	id, err := AddUser()
	fmt.Println(id, err)
}
