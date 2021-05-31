package main

import (
	"fmt"

	"github.com/ivaroliRU/KassAPI/service"
)

func main() {
	client := service.New(false, "", "")

	err := client.CreateCharge()
	fmt.Println(err)
}
