package main

import "fmt"

func main() {
	// 算術運算: + - * / %
	var x int = 8 % 3
	fmt.Println(x)
	// 指定運算: = += -= *= /= %=
	x %= 2
	fmt.Println(x)
	// 單元運算: ++ --
	x++
	fmt.Println(x)
	// 比較運算: == != > < >= <=
	var result bool = x != 3
	fmt.Println(result)
	// 邏輯運算: && || !
	fmt.Print(true || false)
}
