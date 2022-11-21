package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	u := User{
		ID:   1,
		Name: "Ong",
		Age:  20,
	}

	b, err := json.Marshal(u)

	fmt.Printf("string: %T\n", b)
	fmt.Println("err: %v", err)
}
