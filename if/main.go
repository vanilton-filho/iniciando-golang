package main

import "fmt"

func main() {

	a := 10

	if a > 10 {
		fmt.Println("a > 10")
	} else if a > 5 {
		fmt.Println("a > 5")
	} else {
		fmt.Println(":(")
	}

	b := true

	if lang := "golang"; b {
		fmt.Println(lang)
	}
}
