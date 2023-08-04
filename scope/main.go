package main

import "fmt"

func main() {
	x := 10
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	myPrint()
}

func printX() {
	fmt.Println(y)
}

var y int = 20 // no escopo do pacote
