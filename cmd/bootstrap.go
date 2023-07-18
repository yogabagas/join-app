package cmd

import (
	"github/yogabagas/print-in/pkg/database/sql"
	"github/yogabagas/print-in/shared/constant"
)

var (
	configURL string

	sqlDB = sql.DBConn
)

func InitSQLModule() (*sql.DB, error) {
	return sql.NewDBConn(constant.MySQL.String())
}
