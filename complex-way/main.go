package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"-"`
	ID        int    `json:"id":`
	Title     string `json:"title,omitempty"`
	Completed bool   `json:"completed"`
}

func main() {
	url := "https://jsonplaceholder.typicode.com/todos/4"

	//we will send a get request to the api
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err) //execution will stop here
	}

	defer response.Body.Close() //close the body after reading to prevent the memory leaks

	if response.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(response.Body) //reading the response body
		if err != nil {
			log.Fatal(err)
		}

		//reading data in string directly
		data := string(bodyBytes)
		fmt.Println("The data is following:")
		fmt.Println(data)

		//reading json in a struct
		todoItem := Todo{}
		json.Unmarshal(bodyBytes, &todoItem)
		fmt.Printf("To do item is: %+v", todoItem)
	}
}
