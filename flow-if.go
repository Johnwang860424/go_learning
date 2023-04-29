package main

import "fmt"

func main() {
	var num int
	fmt.Print("請輸入一個整數:")
	var _, err = fmt.Scanln(&num)
	if err != nil {
		fmt.Println("無效的輸入")
		return
	}
	if num < 0 {
		fmt.Println("負數")
	} else if num == 0 {
		fmt.Println("零")
	} else if num > 0 {
		fmt.Println("正數")
	}
}
