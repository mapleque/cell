package server

import (
	"context"
	"database/sql"
	"time"
)

// DB The sql database interface
type DB interface {
	BeginTx(context.Context, *sql.TxOptions) (*sql.Tx, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

var (
	dbCtx, _ = context.WithTimeout(context.Background(), 10*time.Second)
)
