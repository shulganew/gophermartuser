package main

import (
	"context"
	"flag"
	"fmt"
	"gophermarketuser/internal/db"
	"gophermarketuser/internal/registration"
)

func main() {
	// go run ./... -m 2
	mode := flag.Int("m", 1, "Client work mode")
	flag.Parse()
	fmt.Println("Mode: ", *mode)
	ctx := context.Background()
	conn, err := db.InitDB(ctx, "postgres://customer:1@postgres:5433/customer")
	if err != nil {
		panic(err)
	}

	switch *mode {
	case 1:
		registration.UserReg(ctx, conn)
	case 2:
		registration.UserLogin(ctx, conn)

	}

}
