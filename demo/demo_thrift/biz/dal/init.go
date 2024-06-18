package dal

import "github.com/Ccmuyu/biz-demo/gomall/demo_thrift/biz/dal/mysql"

func Init() {
	// redis.Init()

	mysql.Init()
}
