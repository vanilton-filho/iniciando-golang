package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	response, _ := http.Get("http://google.com")
	body, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()

	fmt.Printf("%s", body)
}
