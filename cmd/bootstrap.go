package cmd

import (
	"fmt"
	"github/yogabagas/join-app/config"
	"github/yogabagas/join-app/pkg/cache"
	"github/yogabagas/join-app/pkg/cache/redis"
	"github/yogabagas/join-app/pkg/database/sql"
	"github/yogabagas/join-app/shared/constant"
	"net/url"
)

var (
	configURL string

	sqlDB       = sql.DBConn
	redisClient cache.Cache
)

func InitSQLModule() (*sql.DB, error) {
	return sql.NewDBConn(constant.PostgreSQL.String())
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
