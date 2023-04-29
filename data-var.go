package main // 可執行程式必須使用 main 封包

import "fmt" // 匯入內建的 fmt 封包，用來做基本輸出輸入

func main() { // main 函式是程式的進入點
	// 執行程式時，從這邊開始
	// 輸出訊息到終端: fmt.Println(資料, 資料, 資料, ...)
	/*
		fmt.Println(3)           // 整數 int
		fmt.Println(3.1415)      // 浮點數 float64
		fmt.Println("Hello, 世界") // 字串 string
		fmt.Println(true)        // 布林值 bool
		fmt.Println('a')         // 字符 rune
	*/
	var x int // 宣告變數 x 為整數
	x = 3     // 賦值
	fmt.Println(x)
	x = 10 // 指定新的資料: 把資料 10 放進變數，覆蓋原資料
	fmt.Println(x)
	var f float64 = 3.1415 // 宣告變數 f 為浮點數，並賦值
	fmt.Println(f)
	var s string = "Hello, 世界" // 宣告變數 s 為字串，並賦值
	fmt.Println(s)
	var test bool = true // 宣告變數 test 為布林值，並賦值
	fmt.Println(test)
	var c rune = 'a' // 宣告變數 c 為字符，並賦值
	fmt.Println(c)
}
