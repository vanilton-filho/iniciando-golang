package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	x := 0
	for x < 10 {
		fmt.Println(x)
		x++
	}

	// loop infinito
	for {
		fmt.Println(x)
		x++

		// stop
		if x == 100 {
			break
		}
	}

}
