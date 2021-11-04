package databasesetup

import (
	"github.com/go-pg/pg"
)

func Setup() *pg.DB {
	connection := &pg.Options{
		User:     "postgres",
		Password: "codecraft",
		Addr:     ":8080",
		Database: "practice",
	}
	con := pg.Connect(connection)
	return con
}
