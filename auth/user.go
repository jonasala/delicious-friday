package auth

import (
	"encoding/json"
	"io/ioutil"
)

var Users []User

type User struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `json:"password,omitempty"`
}

func LoadUsers() error {
	fileBytes, err := ioutil.ReadFile("users.json")
	if err != nil {
		return err
	}

	if err := json.Unmarshal(fileBytes, &Users); err != nil {
		return err
	}

	return nil
}
