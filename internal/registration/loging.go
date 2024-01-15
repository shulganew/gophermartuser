package registration

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"gophermarketuser/internal/model"
	"io"
	"net/http"

	"github.com/jackc/pgx/v5"
)

func UserLogin(ctx context.Context, conn *pgx.Conn) {

	var users []model.User

	rows, err := conn.Query(ctx, "SELECT jwt, login, password FROM users")
	if err != nil {
		panic(fmt.Errorf("Query errror: %w", err))
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		err = rows.Scan(&user.JWT, &user.Login, &user.Password)
		if err != nil {
			panic(fmt.Errorf("Scan error %w", err))
		}
		users = append(users, user)
	}

	reqBodyDel := bytes.NewBuffer([]byte{})

	nUser := 0

	err = json.NewEncoder(reqBodyDel).Encode(&users[nUser])
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	request, err := http.NewRequest(http.MethodPost, "http://localhost:8090/api/user/login", reqBodyDel)
	if err != nil {
		fmt.Println(err)
	}

	// add jwt
	request.Header.Add("Authorization", "Bearer "+users[nUser].JWT.String)

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

	jwt := res.Header.Get("Authorization")[len("Bearer "):]

	_, err = conn.Exec(ctx, "UPDATE users SET jwt = $1", jwt)
	if err != nil {
		panic(err)
	}

}
