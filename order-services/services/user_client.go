package services

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func UserExits(UserID int) bool {
	resp, err := http.Get("http://localhost:8081/users")
	if err != nil {
		fmt.Println("Error Calling the User Services", err.Error())
		return false
	}
	resp.Body.Close()
	var user []User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		fmt.Println("Error decoding the user", err)
		return false
	}
	for _, u := range user {
		if u.ID == UserID {
			return true
		}
	}
	return false
}
