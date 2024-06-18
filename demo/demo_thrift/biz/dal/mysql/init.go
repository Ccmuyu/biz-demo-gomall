package mysql

import (
	"fmt"

	"github.com/Ccmuyu/biz-demo/gomall/demo_thrift/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	fmt.Println("dsn 1:\n", conf.GetConf().MySQL.DSN)
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, "root", "111111")
	fmt.Println("dsn:\n", dsn)
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v", DB.Debug().Exec("select version()"))
}
