package service

import (
	"context"
	"fmt"
	"gophermarketuser/internal/model"

	"github.com/jackc/pgx/v5"
)

func GetUser(ctx context.Context, conn *pgx.Conn) *model.User {

	row := conn.QueryRow(ctx, "SELECT jwt, login, password FROM users ORDER BY random() LIMIT 1")

	var user model.User
	err := row.Scan(&user.JWT, &user.Login, &user.Password)
	if err != nil {
		panic(fmt.Errorf("Scan error %w", err))
	}

	return &user

}
