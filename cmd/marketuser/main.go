package main

import (
	"context"
	"flag"
	"fmt"
	"gophermarketuser/internal/balance"
	"gophermarketuser/internal/db"
	"gophermarketuser/internal/orders"
	"gophermarketuser/internal/registration"
	"os"
)

func main() {
	// go run ./... -m 2
	mode := flag.Int("m", 1, "Client work mode")
	help := flag.Bool("h", false, "help")
	flag.Parse()

	if *help {
		fmt.Println(`
		go run -v ./... -m 1 - Register new user
		go run -v ./... -m 2 - Login user
		go run -v ./... -m 3 - Get users order list
		go run -v ./... -m 4 - Set user order
		go run -v ./... -m 5 - 
		
		`)
		os.Exit(0)
	}
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
	case 3:
		orders.GetOrders(ctx, conn)
	case 4:
		orders.SetOrders(ctx, conn)
	case 5:
		balance.GetBalance(ctx, conn)
	case 6:
		balance.SetWithdraw(ctx, conn)
	case 7:
		balance.GetWithdrawals(ctx, conn)
	}

}
