package balance

import (
	"context"
	"fmt"
	"gophermarketuser/internal/service"
	"io"
	"net/http"

	"github.com/jackc/pgx/v5"
)

func GetWithdrawals(ctx context.Context, conn *pgx.Conn) {

	fmt.Println("Make Withdrawals")

	user := service.GetUser(ctx, conn)

	client := &http.Client{}
	request, err := http.NewRequest(http.MethodGet, "http://localhost:8090/api/user/withdrawals", nil)
	if err != nil {
		fmt.Println(err)
	}

	// add jwt
	request.Header.Add("Authorization", "Bearer "+user.JWT.String)

	//reqest
	request.Header.Add("Content-Type", "text/plain")
	res, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	//=============================Response===================
	for k, v := range res.Header {

		fmt.Printf("%s: %v\r\n", k, v[0])

	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Body: ", string(body))
	fmt.Printf("Status Code: %d\r\n", res.StatusCode)

}
