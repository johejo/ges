package mysql

import (
	"database/sql"
	"fmt"
	"net/url"

	"github.com/johejo/gohejo/envutils"
	"github.com/johejo/gohejo/logutils"

	_ "github.com/go-sql-driver/mysql"
)

var logger = logutils.New()

func OpenConnection() *sql.DB {
	username := envutils.GetEnv("MYSQL_USER", "message")
	password := envutils.GetEnv("MYSQL_PASSWORD", "message")
	protocol := envutils.GetEnv("MYSQL_PROTOCOL", "tcp")
	hostname := envutils.GetEnv("MYSQL_HOSTNAME", "localhost")
	port := envutils.GetEnv("MYSQL_PORT", "3306")
	database := envutils.GetEnv("MYSQL_DATABASE", "messagedb")
	charset := envutils.GetEnv("MYSQL_CHARSET", "utf8mb4")
	locale := envutils.GetEnv("MYSQL_LOCALE", "Asia/Tokyo")
	tls := envutils.GetEnv("MYSQL_TLS", "false")

	connStr := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=%s&parseTime=True&loc=%s&tls=%s",
		username, password, protocol, hostname, port, database, charset, url.QueryEscape(locale), tls,
	)

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		panic(err)
	}

	return db
}
