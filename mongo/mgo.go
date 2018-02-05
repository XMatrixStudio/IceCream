package mongo

import (
	"fmt"
	"log"

	"github.com/XMatrixStudio/icecream/config"
	"gopkg.in/mgo.v2"
)

// Init 初始化连接
func Init() {
	session, err := mgo.Dial(
		"mongodb://" +
			config.Mongo.User +
			":" + config.Mongo.Password +
			"@" + config.Mongo.Host +
			":" + config.Mongo.Port +
			"/" + config.Mongo.Name)
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("Mongodb Connent Success")
	}
	db := session.DB(config.Mongo.Name)
	db.C("users")
}
