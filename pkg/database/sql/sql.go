package sql

import (
	"database/sql"
	"fmt"
	"github/yogabagas/join-app/config"
	"github/yogabagas/join-app/shared/constant"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type DB struct {
	MySQL      *sql.DB
	PostgreSQL *sql.DB
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

			log.Printf("database connection : %s", schemaURL)

			db, err := sql.Open(constant.MySQL.String(), schemaURL)
			if err != nil {
				log.Panic(err.Error())
				panic(err)
			}

			if err := db.Ping(); err != nil {
				log.Panic(err)
			}
			DBConn = &DB{MySQL: db}
		case constant.PostgreSQL.String():

			schemaURL := fmt.Sprintf("%s://%s:%s@%s/%s?sslmode=disable",
				constant.PostgreSQL.String(),
				config.GlobalCfg.DB.SQL.User,
				config.GlobalCfg.DB.SQL.Password,
				config.GlobalCfg.DB.SQL.Host,
				config.GlobalCfg.DB.SQL.Schema)

			log.Printf("database connection : %s", schemaURL)

			db, err := sql.Open(constant.PostgreSQL.String(), schemaURL)
			if err != nil {
				log.Panic(err.Error())
				panic(err)
			}

			if err := db.Ping(); err != nil {
				log.Panic(err)
			}
			DBConn = &DB{PostgreSQL: db}
		}
	}

	return DBConn, nil
}
