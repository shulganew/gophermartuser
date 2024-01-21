package orders

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"gophermarketuser/internal/model"
	"gophermarketuser/internal/service"
	"io"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
)

type RequestAccural struct {
	Order string        `json:"order"`
	Goods []model.Goods `json:"goods"`
}

func GetOrders(ctx context.Context, conn *pgx.Conn) {

	fmt.Println("Set Order")

	user := service.GetUser(ctx, conn)

	fmt.Println("Set Order for User", user.Login)

	client := &http.Client{}
	request, err := http.NewRequest(http.MethodGet, "http://localhost:8088/api/user/orders", nil)
	if err != nil {
		fmt.Println(err)
	}

	// add jwt
	request.Header.Add("Authorization", user.JWT.String)

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

func SetOrders(ctx context.Context, conn *pgx.Conn) {
	fmt.Println("Set Order")

	user := service.GetUser(ctx, conn)

	fmt.Println("Set Order for Login", user.Login)

	// Generate luna number for order
	//onumber := goluhn.Generate(10)
	onumber := "0265410704"

	fmt.Println("Order Number: ", onumber)

	order := bytes.NewBuffer([]byte(onumber))

	client := &http.Client{}
	request, err := http.NewRequest(http.MethodPost, "http://localhost:8088/api/user/orders", order)
	if err != nil {
		log.Fatalln(err)
	}

	// add jwt
	request.Header.Add("Authorization", user.JWT.String)

	//reqest
	request.Header.Add("Content-Type", "text/plain")
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

	fmt.Println("========================================")
	fmt.Println("=================Accural================")
	goods := LoadGoods(ctx, conn, 1)

	reqAcc := RequestAccural{Order: onumber, Goods: goods}
	jsonAccural, err := json.Marshal(reqAcc)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Send to accural system ", string(jsonAccural))

	bodyAcc := bytes.NewReader(jsonAccural)

	client = &http.Client{}
	request, err = http.NewRequest(http.MethodPost, "http://localhost:8090/api/orders", bodyAcc)
	if err != nil {
		log.Fatalln(err)
	}

	//reqest
	request.Header.Add("Content-Type", "application/json")

	res, err = client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("=================Accural Responce================")

	for k, v := range res.Header {

		fmt.Printf("%s: %v\r\n", k, v[0])

	}

	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Body: ", string(body))
	fmt.Printf("Status Code: %d\r\n", res.StatusCode)

	fmt.Println("=================Accural Check ================")

	client = &http.Client{}
	request, err = http.NewRequest(http.MethodGet, "http://localhost:8090/api/orders/"+onumber, nil)
	if err != nil {
		log.Fatalln(err)
	}

	res, err = client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("=================Accural Check Responce================")

	for k, v := range res.Header {

		fmt.Printf("%s: %v\r\n", k, v[0])

	}

	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Body: ", string(body))
	fmt.Printf("Status Code: %d\r\n", res.StatusCode)

}

func LoadGoods(ctx context.Context, conn *pgx.Conn, number int) []model.Goods {
	var goods []model.Goods

	rows, err := conn.Query(ctx, "SELECT description, price FROM goods ORDER BY random() LIMIT $1", number)
	if err != nil {
		panic(fmt.Errorf("Query errror: %w", err))
	}
	defer rows.Close()

	for rows.Next() {
		var good model.Goods
		err = rows.Scan(&good.Description, &good.Price)
		if err != nil {
			panic(fmt.Errorf("Scan error %w", err))
		}
		goods = append(goods, good)
	}

	return goods
}
