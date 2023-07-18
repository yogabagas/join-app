package sql

import (
	"database/sql"
	"fmt"
	"github/yogabagas/print-in/config"
	"github/yogabagas/print-in/shared/constant"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	MySQL *sql.DB
}

var DBConn *DB

func NewDBConn(kind string) (*DB, error) {

	if DBConn == nil {
		switch kind {
		case constant.MySQL.String():

			schemaURL := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
				config.GlobalCfg.DB.SQL.User,
				config.GlobalCfg.DB.SQL.Password,
				config.GlobalCfg.DB.SQL.Host,
				config.GlobalCfg.DB.SQL.Schema)

			db, err := sql.Open(constant.MySQL.String(), schemaURL)
			if err != nil {
				log.Panic(err.Error())
				panic(err)
			}

			if err := db.Ping(); err != nil {
				log.Panic(err)
			}
			DBConn = &DB{MySQL: db}
		}
	}

	return DBConn, nil
}
