package database

import (
	"context"
	"database/sql"
	"sso/database/driver"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Conn struct {
	DB *sql.DB
}

func NewConn(db *sql.DB) *Conn {
	return &Conn{DB: db}
}

func Open() (*Conn, error) {
	db, err := driver.Connect("mysql")
	if err != nil {
		fmt.Println(err)
	}
	return NewConn(db), nil
}

func DbExec(db *sql.DB, ctx context.Context, query string, args ...interface{}) error {
	_, err := db.ExecContext(ctx, query, args...)
	return err
}

func QueryDb(db *sql.DB, ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return db.QueryContext(ctx, query, args...)
}

func QueryRowFromDb(db *sql.DB, ctx context.Context, query string, args ...interface{}) *sql.Row {
	return db.QueryRowContext(ctx, query, args...)
}

func (conn *Conn) Close() {
	conn.DB.Close()
}

func (conn *Conn) Insert(ctx context.Context, query string, args ...interface{}) error {
	return DbExec(conn.DB, ctx, query, args...)
}

func (conn *Conn) Delete(ctx context.Context, query string, args ...interface{}) error {
	return DbExec(conn.DB, ctx, query, args...)
}

func (conn *Conn) Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return QueryDb(conn.DB, ctx, query, args...)
}

func (conn *Conn) QueryRow(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return QueryRowFromDb(conn.DB, ctx, query, args...)
}
