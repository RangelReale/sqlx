package sqlx

import (
	"context"
	"database/sql"
)

func (db *DB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return db.db.Query(query, args...)
}

func (db *DB) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return db.db.QueryContext(ctx, query, args...)
}

func (db *DB) QueryRow(query string, args ...interface{}) *sql.Row {
	return db.db.QueryRow(query, args...)
}

func (db *DB) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return db.db.QueryRowContext(ctx, query, args...)
}

func (db *DB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.db.Exec(query, args...)
}

func (db *DB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return db.db.ExecContext(ctx, query, args...)
}

func (db *DB) Prepare(query string) (*sql.Stmt, error) {
	return db.db.Prepare(query)
}

func (db *DB) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	return db.db.PrepareContext(ctx, query)
}
