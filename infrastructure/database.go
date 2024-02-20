package infrastructure

import (
	"database/sql"

	"github.com/fiaufar/sawit-pro-test/util"
	_ "github.com/lib/pq"
)

type DbConnection struct {
	Db *sql.DB
}

type NewDbConnectionOptions struct {
	Dsn string
}

func NewDbConnection(opts NewDbConnectionOptions) *DbConnection {
	db, err := sql.Open("postgres", opts.Dsn)
	if err != nil {
		util.Log.Error(err)
		panic(err)
	}

	if err := db.Ping(); err != nil {
		util.Log.Error(err)
	}

	return &DbConnection{
		Db: db,
	}
}
