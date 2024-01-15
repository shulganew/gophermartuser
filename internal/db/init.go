package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// Init Database
func InitDB(ctx context.Context, dsn string) (db *pgx.Conn, err error) {

	db, err = pgx.Connect(ctx, dsn)
	if err != nil {
		return nil, err
	}

	return
}
