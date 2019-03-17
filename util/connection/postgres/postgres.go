package postgres

import (
	"fmt"

	config "github.com/fzlrhmn/simple-jwt/util/config"
	"github.com/go-pg/pg"
)

var pgConns *pg.DB

// Initialize is init connection to postgres and set it as single skeleton
func Initialize() bool {
	host := config.Instance.GetString("db.postgres.host")
	port := config.Instance.GetInt("db.postgres.port")
	user := config.Instance.GetString("db.postgres.user")
	password := config.Instance.GetString("db.postgres.password")
	database := config.Instance.GetString("db.postgres.database")

	pgConns = pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%v:%v", host, port),
		User:     user,
		Password: password,
		Database: database,
	})

	return true
}

func GetInstance() *pg.DB {
	return pgConns
}
