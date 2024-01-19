package balance

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"gophermarketuser/internal/service"
	"io"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
)

type Withdraw struct {
	Onumber   string  `json:"order"`
	Withdrawn float64 `json:"sum"`
}

func SetWithdraw(ctx context.Context, conn *pgx.Conn) {
	fmt.Println("Withdraw Order")

	user := service.GetUser(ctx, conn)

	fmt.Println("Set Order for Login", user.Login)

	// Generate luna number for order

	onumber := "0267010072"

	fmt.Println("Order Number: ", onumber)

	sum := 1000.0
	wd := Withdraw{Onumber: onumber, Withdrawn: sum}

	jsonWs, err := json.Marshal(wd)
	if err != nil {
		log.Fatalln(err)
	}

	jbody := bytes.NewReader(jsonWs)
	client := &http.Client{}
	request, err := http.NewRequest(http.MethodPost, "http://localhost:8090/api/user/balance/withdraw", jbody)
	if err != nil {
		log.Fatalln(err)
	}

	// add jwt
	request.Header.Add("Authorization", "Bearer "+user.JWT.String)

	//reqest
	request.Header.Add("Content-Type", "application/json")
	res, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	//

	//=================Market Response================
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
