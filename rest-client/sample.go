package rest_client

import (
	"fmt"
	"io"
	"net/http"
)

func SampleRestClientBuiltIn() {
	JsonPlaceholderBaseUrl := "https://jsonplaceholder.typicode.com/todos/1"
	req, err := http.NewRequest(http.MethodGet, JsonPlaceholderBaseUrl, nil)
	if err != nil {
		panic(err)
	}

	// simulasi request
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("resp.body:", string(body))
}
