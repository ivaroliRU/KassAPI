package main

import (
	"encoding/json"
	"fmt"

	"github.com/ivaroliRU/KassAPI/service"
)

func main() {
	client := service.New(false, "kass_test_auth_token", "")
	response, _ := client.CreateCharge(2199, "Kass bolur", "https://photos.kassapi.is/kass/kass-bolur.jpg", "ABC12332öiö23iö3", "1001001", 1, 90, "asdfasdf")
	out, _ := json.Marshal(response)

	fmt.Println(string(out))
}
