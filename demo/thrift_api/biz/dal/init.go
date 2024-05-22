package dal

import (
	"github.com/Ccmuyu/biz-demo/gomall/thrift_api/biz/dal/mysql"
	"github.com/Ccmuyu/biz-demo/gomall/thrift_api/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
