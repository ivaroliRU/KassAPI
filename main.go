package main

import (
	"fmt"
)

func main() {
	client := New(false, "", "")

	err := client.CreateCharge()
	fmt.Println(err)
}
