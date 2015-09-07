package storage

import (
	"log"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jf-marino/simple.go.server/storage/models"
)

var s orm.Ormer

func Orm() orm.Ormer {
	if s == nil {
		s = orm.NewOrm()
	}
	return s
}

func registerModels() {
	orm.RegisterModel(new(models.Event))
}

func Initialize() {
	orm.Debug = true
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	err := orm.RegisterDataBase("default", "mysql", "root:password@/SimpleDb?charset=utf8")
	if err == nil {
		registerModels()
	} else {
		log.Println(err)
	}
}
