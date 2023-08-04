package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var url string

	fmt.Print("Type URL: ")
	fmt.Scan(&url)

	response, _ := http.Get(url)
	body, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()

	fmt.Printf("%s", body)
}
