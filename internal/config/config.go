package config

import "flag"

var (
	clickhouseUrlFlag = flag.String("url", "", "setting up a clickhouse url")
	databaseFlag      = flag.String("database", "", "setting up a database name")
	tableFlag         = flag.String("table", "", "setting up a table name")
	usernameFlag      = flag.String("username", "", "setting up a clickhouse username")
	passwordFlag      = flag.String("password", "", "setting up a clichouse password")
	// clickhouseWriteContextTimeoutFlag = flag.String("timeout", "", "clickhouse write context timeout")

	clickMailUrlFlag = flag.String("clickmail", "", "clickhouse mail ru lib url")
)

type Config struct {
	ClickhouseURL string
	Database      string
	Table         string
	Username      string
	Password      string

	ClickmailURL string
}

func GetConfig() Config {
	flag.Parse()

	c := Config{
		ClickhouseURL: *clickhouseUrlFlag,
		Database:      *databaseFlag,
		Table:         *tableFlag,
		Username:      *usernameFlag,
		Password:      *passwordFlag,

		ClickmailURL: *clickMailUrlFlag,
	}

	return c
}
