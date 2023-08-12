package main

import "fmt"

func xpto(a *int) int {
	*a = 100 + 1
	fmt.Println("Valor de a  = ", a)
	return *a
}

func main() {

	b := 10
	fmt.Println(xpto(&b))
	fmt.Println("Valor de b = ", &b)

	x := 10
	fmt.Println(x)  // 10
	fmt.Println(&x) // 0xc00001a078

	y := &x
	fmt.Println(y)  // 0xc00001a078
	fmt.Println(*y) // 10

	*y = 20
	fmt.Println(x) // 20

	var z *int = &x
	fmt.Println(z) // 0xc00001a078
	fmt.Println(*z)
	*z = 42
	fmt.Println(*z)
	fmt.Println(x)

}
