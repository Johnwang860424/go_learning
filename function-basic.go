package main

import "fmt"

func show(msg string) {
	fmt.Println(msg)
}

func add(num ...int) {
	var result int
	for _, v := range num {
		result += v
	}
	fmt.Println(result)
}

func main() {
	show("Hello, World!")
	add(1, 2, 3, 4, 5)
}
