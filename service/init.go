package service

import (
	"fmt"
	"log"

	"github.com/go-xorm/xorm"
	"github.com/r-hata/sample-api/model"
)

var DbEngine *xorm.Engine

func init() {
	driverName := "mysql"
	DsName := "root:root@(10.0.2.15:3306)/gin?charset=utf8"

	var err error
	DbEngine, err = xorm.NewEngine(driverName, DsName)
	if err != nil {
		log.Fatal(err.Error())
	}
	DbEngine.ShowSQL(true)
	DbEngine.SetMaxOpenConns(2)
	DbEngine.Sync2(new(model.Book))
	fmt.Println("init data base ok")
}
