package rest_client

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type Todo struct {
	Id        int    `json:"id"`
	UserId    int    `json:"userId"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func SampleRestClientResty() {
	client := resty.New()

	URL := "https://jsonplaceholder.typicode.com/todos"

	res, err := client.R().Get(URL)
	if err != nil {
		panic(err)
	}

	var todos []Todo
	err = json.Unmarshal(res.Body(), &todos)
	if err != nil {
		panic(err)
	}

	// rumus pagination
	page := 2
	size := 20

	startIndex := (page - 1) * size
	endIndex := startIndex + size

	for _, todo := range todos[startIndex:endIndex] {
		fmt.Println(todo)
	}
}
