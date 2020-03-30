package service

import (
	"fmt"

	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
	"github.com/go-redis/redis"
)

type Redis struct {
	Client *redis.Client
}

func (i *Redis) Init() *Redis {
	conf, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		fmt.Println("new config failed, err:", err)
		logs.Error("new config failed, err:", err)

	} else {
		host := conf.String("redis::host")
		pwd := conf.String("redis::pwd")
		i.Client = redis.NewClient(&redis.Options{
			Addr:     host,
			Password: pwd,
			DB:       0,
		})
		_, err := i.Client.Ping().Result()
		if err != nil {
			fmt.Printf("ping error[%s]\n", err.Error())
			panic(err.Error())
		}
	}
	return i
}

func (i *Redis) Close() {
	i.Client.Close()
}
