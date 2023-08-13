package main

import "fmt"

func doubleMe(a int) int {
	return a * a
}

func namedReturn(a string) (x string) {
	x = a
	return
}

func manyReturn(a string, b string) (string, string) {
	return a, b
}

func variadicFunction(x ...int) int {
	var res int
	for key, v := range x {
		res += v
		fmt.Println(key)
	}

	fmt.Println(res)

	return res
}

func inside() func() int {
	x := 10
	return func() int {
		return x * x
	}
}

func main() {

	x := doubleMe(5)

	fmt.Println(doubleMe(x))
	fmt.Println(doubleMe(10))

	fmt.Println(namedReturn("Vanilton"))
	hello, world := manyReturn("Hello", "World")
	fmt.Println(hello, world)

	variadicFunction(2022, 2021)

	// clojure
	z := 0
	add := func() int {
		z += 2
		return z
	}

	fmt.Println(add())
	fmt.Println(inside()())

}
