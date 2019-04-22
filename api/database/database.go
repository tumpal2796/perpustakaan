package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var Conn map[string]*sqlx.DB

type Options map[string]*struct {
	Master string
}

func InitConnection(config interface{}) {
	var err error
	Conn = make(map[string]*sqlx.DB)
	cfg := config.(Options)

	for key, val := range cfg {
		Conn[key], err = sqlx.Connect("postgres", val.Master)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}
