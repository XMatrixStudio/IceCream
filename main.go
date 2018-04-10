package main

import (
	"net/http"

	"github.com/XMatrixStudio/IceCream/router"
	"github.com/XMatrixStudio/icecream/model"
)

func main() {
	err := model.InitMongo() // 连接Mongo数据库
	if err != nil {
		panic(err)
	}
	http.ListenAndServe(":30020", router.Router())
}
