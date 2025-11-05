package services

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func OrderExits(orderId int) bool {
	url := fmt.Sprintf("http://localhost:8082/orders/%d", orderId)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error calling the services", err)
		return false
	}
	resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return false
	}

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println("Error decoding the response", err)
		return false
	}
	return true
}
