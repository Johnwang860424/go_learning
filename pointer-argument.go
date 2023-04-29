package main

import "fmt"

func add(x int) {
	x += 10
	fmt.Println("x =", x)
}

func add2(x *int) {
	*x += 10
	fmt.Println("x =", *x)
}

func main() {
	var a int = 10
	// Pass by Value
	add(a)
	fmt.Println("a =", a)
	// Pass by Pointer
	add2(&a)
	fmt.Println("a =", a)
}
