package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"golang.org/x/oauth2"
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

	conf := &oauth2.Config{
		ClientID:     credentials.ClientID,
		ClientSecret: credentials.ClientSecret,
		RedirectURL:  "http://127.0.0.1:3456/catch",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email", // You have to select your own scope from here -> https://developers.google.com/identity/protocols/googlescopes#google_sign-in
		},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.google.com/o/oauth2/auth",
			TokenURL: "https://accounts.google.com/o/oauth2/token",
		},
	}
	fmt.Println(conf.AuthCodeURL("rabbits"))

}

func check(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
