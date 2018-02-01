package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Credentials struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

var credentials Credentials

func main() {

	file, err := ioutil.ReadFile("./google.json")
	check(err)
	err = json.Unmarshal(file, &credentials)
	check(err)

	fmt.Println(credentials)

}

func check(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
