package registration

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"gophermarketuser/internal/model"
	"io"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5"
)

func UserReg(ctx context.Context, conn *pgx.Conn) {
	i := strconv.Itoa(rand.Intn(1000))
	user := model.User{Login: "Igor" + i, Password: "MySecret"}

	reqBodyDel := bytes.NewBuffer([]byte{})

	err := json.NewEncoder(reqBodyDel).Encode(&user)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	request, err := http.NewRequest(http.MethodPost, "http://localhost:8088/api/user/register", reqBodyDel)
	if err != nil {
		fmt.Println(err)
	}

	//reqest
	request.Header.Add("Content-Type", "application/json")
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

	//jwt := res.Header.Get("Authorization")[len("Bearer "):]

	_, err = conn.Exec(ctx, "INSERT INTO users (jwt, login, password) VALUES ($1, $2, $3)", user.JWT, user.Login, user.Password)
	if err != nil {
		panic(err)
	}

}
