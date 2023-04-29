package main

import (
	"fmt"
	"strings"
)

func show(msg string) {
	if msg == "Hello" {
		return // 結束函式
	}
	fmt.Println(msg)
}

func add(num1 int, num2 int) int {
	result := num1 + num2
	return result
}

func splitName(fullName string) (string, string) {
	// 將全名拆分為姓和名
	names := strings.Split(fullName, " ")
	firstName := names[0]
	lastName := names[len(names)-1]

	// 返回兩個字符串
	return firstName, lastName
}

func main() {
	// 基本的 return 用法
	show("Hello")
	show("妳好，世界！")
	// 單一回傳值
	var result int = add(1, 2)
	fmt.Println(result)
	// 多個回傳值
	firstName, lastName := splitName("Steve Jobs")
	fmt.Println("First name:", firstName)
	fmt.Println("Last name:", lastName)
}
