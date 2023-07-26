package cmd

import (
	"fmt"
	"github/yogabagas/print-in/config"
	"github/yogabagas/print-in/pkg/cache"
	"github/yogabagas/print-in/pkg/cache/redis"
	"github/yogabagas/print-in/pkg/database/sql"
	"github/yogabagas/print-in/shared/constant"
	"net/url"
)

var (
	configURL string

	sqlDB       = sql.DBConn
	redisClient cache.Cache
)

func InitSQLModule() (*sql.DB, error) {
	return sql.NewDBConn(constant.MySQL.String())
}

func InitCache() cache.Cache {
	redisCreds := url.URL{
		Host: config.GlobalCfg.Cache.Redis.Host,
		User: url.UserPassword(config.GlobalCfg.Cache.Redis.User, config.GlobalCfg.Cache.Redis.Password),
	}
	client, err := redis.NewCache(&redisCreds)
	if err != nil {
		fmt.Println(err)
	}

	return client
}
