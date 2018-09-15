package config

import "fmt"

//-* http Methods
const (
	Get = "GET"
	Post = "POST"
	Put = "PUT"
	Patch = "PATCH"
	Delete = "DELETE"
	Options = "OPTIONS"
	Head = "HEAD"
)

const (
	DBHost = "127.0.0.1"
	DBPort = "5432"
	DBUser = "server"
	DBPass = ""
	DBName = "rest_server"
	ConnectionFormat = "host=%s port=%s sslmode=disable dbname=%s user=%s password=%s"
	DriverName = "postgres"
)

// Second parameter for the sql.Open() function. /-* db/db.go *-/
var DataSourceName = fmt.Sprintf(
	ConnectionFormat,
	DBHost, DBPort, DBName, DBUser, DBPass,
)

var API = "api/"
var VersionOne  = API + "v1/"

var Addr = "0.0.0.0:10000"

