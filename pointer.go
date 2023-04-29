package main

import "fmt"

func main() {
	// 建立存放資料的變數
	var num int = 10
	fmt.Println("num 是", num)
	// 取得 num 的記憶體位址
	fmt.Println("num 的記憶體位址是", &num)
	// 存放到指標變數
	var ptr *int = &num
	// 反解指標變數
	fmt.Println("ptr 指向的值是", *ptr)

	var word string = "Hello"
	fmt.Println("word 的記憶體位址", &word)
	var wordptr *string = &word
	fmt.Println("wordptr 指向的值是", *wordptr)
}
