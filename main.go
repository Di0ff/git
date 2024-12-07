package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	var n string
	fmt.Print("Введите целое число: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatal(err)
	}

	var data interface{}

	err = json.Unmarshal([]byte(n), &data)
	if err != nil {
		data = n
	}

	fmt.Printf("Вы ввели число: %d\n", n)
}
