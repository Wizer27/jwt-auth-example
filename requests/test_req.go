package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type UserDat struct {
	UserID int `json:"user_id"`
}

func post() {
	url := "http://127.0.0.1:8080/get_token"

	data := UserDat{
		UserID: 426783647826378,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Ошибка JSON:", err)
		return
	}

	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		fmt.Println("Ошибка создания запроса:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer my_token")
	req.Header.Set("X-API-KEY", "secret")

	client := &http.Client{}
	resp, err := client.Do(req)
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

func get() {
	const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo0MjY3ODM2NDc4MjYzNzh9.hKYSNMGutAA6EgSxUFDtJrlQH9nEtyuTn3H63mYy3ns"

	const url = "http://127.0.0.1:8080/protected"

	req, err := http.NewRequest(
		"GET",
		url,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("X-API-KEY", "secret")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения:", err)
		return
	}

	fmt.Println("Status:", resp.StatusCode)
	fmt.Println("Body:", string(body))

}

func main() {
	get()
}
