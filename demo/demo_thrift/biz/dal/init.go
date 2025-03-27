package dal

import (
	"github.com/zxjia2002/gomall/demo/demo_thrift/biz/dal/mysql"
	"github.com/zxjia2002/gomall/demo/demo_thrift/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
