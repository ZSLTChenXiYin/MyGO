package redis

import (
	"github.com/ZSLTChenXiYin/MyGO/configure"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	cli *redis.Client
}

func NewRedis(conf configure.Configuration) *Redis {
	return &Redis{
		cli: redis.NewClient(&redis.Options{
			Addr:     conf.Redis().Address(),
			Password: conf.Redis().Password(),
			DB:       conf.Redis().DB(),
		}),
	}
}

func (r *Redis) Redis() *redis.Client { return r.cli }
