package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	var n string
	fmt.Print("Введите данные: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatal(err)
	}

	var data interface{}

	err = json.Unmarshal([]byte(n), &data)
	if err != nil {
		data = n
	}

	fmt.Printf("Вы ввели следующие данные: %s\n", n)
}
