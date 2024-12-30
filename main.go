package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Todo struct {
	//we can have private fields by using smallcase , it wont be visible to anything outside the package
	UserId    int    `json:"userId"` //if u dont want to show this then use -
	ID        int    `json:"id":`
	Title     string `json:"title,omitempty"` //use this json flag to ignore empty srings
	Completed bool   `json:"completed"`
}

func main() {
	url := "https://jsonplaceholder.typicode.com/todos/1"

	//we will send a get request to the api
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err) //execution will stop here
	}

	defer response.Body.Close() //close the body after reading to prevent the memory leaks

	//simplest way to do this is
	if response.StatusCode == http.StatusOK {
		todoItem := Todo{}
		decoder := json.NewDecoder(response.Body)
		//when you have an unknown field in response ,stop execution
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&todoItem); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("To do item is: %+v", todoItem)

		//convert data back to json
		todo, err := json.MarshalIndent(todoItem, "", "\t")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println()
		fmt.Println(string(todo))
	}
}
