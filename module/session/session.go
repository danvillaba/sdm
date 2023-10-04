package session

import "github.com/jmoiron/sqlx"

type Session struct {
	Id       string
	Db       *sqlx.DB
	Driver   string
	Database string
}
