package cmd

import (
	"github/yogabagas/join-app/pkg/database/sql"
	"github/yogabagas/join-app/shared/constant"
)

var (
	configURL string

	sqlDB = sql.DBConn
)

func InitSQLModule() (*sql.DB, error) {
	return sql.NewDBConn(constant.MySQL.String())
}
