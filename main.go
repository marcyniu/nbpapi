package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	resp, err := getTable("A")

	if (err != nil) {
		fmt.Println(err)
	}

	fmt.Println("Works")

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
}

func getTable(tableName string) (*http.Response, error) {
	resp, err := http.Get("http://api.nbp.pl/api/exchangerates/tables/" + tableName + "/?format=json")
	return resp, err
}
