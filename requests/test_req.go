package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type UserDat struct {
	UserID int `json:"num1"`
}

func post() {
	const url = "http://127.0.0.1:8080"

	data := UserDat{
		UserID: 5,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Ошибка JSON:", err)
		return
	}
	resp, err := http.Post(
		"http://127.0.0.1:8080/get_token",
		"application/json",
		bytes.NewBuffer(jsonData),
	)

	if err != nil {
		fmt.Println("Ошибка запроса:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения:", err)
		return
	}

	fmt.Println("Status:", resp.StatusCode)
	fmt.Println("Body:", string(body))
}

func main() {
	post()
}
