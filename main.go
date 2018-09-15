package main

import (
	"database/sql"
	"makebex-backend/server"
	"makebex-backend/server/bot"
	"makebex-backend/server/config"
	"makebex-backend/server/views/db"
)

type Env struct {
	DB *sql.DB
}

var env = Env {}

func main() {
	env.DB = db.DB()

	// For dev.
	go bot.StartBot(env.DB)

	server.On(config.Addr, env.DB)
}
