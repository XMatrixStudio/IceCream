package main

import (
	"net/http"
	"github.com/XMatrixStudio/IceCream/router"
	"github.com/XMatrixStudio/icecream/lib"
)

func main() {
	err := lib.InitMongo() // 连接Mongo数据库
	if err != nil {
		panic(err)
	}
	http.ListenAndServe(":30020", router.Router())
}
