package dal

import (
	"github.com/Ccmuyu/biz-demo/gomall/demo_thrift/biz/dal/mysql"
	"github.com/Ccmuyu/biz-demo/gomall/demo_thrift/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
