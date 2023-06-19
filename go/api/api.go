package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	Id       int
	Username string
	Password string
	Mail     string
}

func CallApiUsers(table *[]User) {
	url := "https://service-apiforum.onrender.com/users"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	res, err2 := client.Do(req)
	if err2 != nil {
		log.Fatal(err2)
	}
	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, table)
}

func CallUser(table *User, user string) {
	urll := "http://localhost:8080/getuserv?username=" + user
	req, err := http.NewRequest("GET", urll, nil)
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	res, err2 := client.Do(req)
	if err2 != nil {
		log.Fatal(err2)
	}
	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, table)
}
