package service

import (
	"fmt"
	"log"

	"github.com/go-xorm/xorm"
	"github.com/r-hata/sample-api/model"
	"xorm.io/core"
)

var dbEngine *xorm.Engine

func init() {
	driverName := "mysql"
	dsName := "gin:gin123@tcp([127.0.0.1]:3306)/gin?charset=utf8"

	var err error
	dbEngine, err = xorm.NewEngine(driverName, dsName)
	if err != nil {
		log.Fatal(err.Error())
	}
	dbEngine.ShowSQL(true)
	dbEngine.SetMapper(core.GonicMapper{})
	dbEngine.SetMaxOpenConns(2)
	err = dbEngine.Sync2(new(model.Book))
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("init data base ok")
}
