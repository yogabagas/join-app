package cmd

import (
	"fmt"
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
		Host: "localhost:6379",
		User: url.UserPassword("", "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81"),
	}
	client, err := redis.NewCache(&redisCreds)
	if err != nil {
		fmt.Println(err)
	}

	return client
}
