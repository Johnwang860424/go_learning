package main

import "fmt"

func main() {
	// 整數陣列
	var numbers [3]int
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	fmt.Println(numbers)
	fmt.Println(numbers[1] * 2)
	// 字串陣列
	var names [3]string = [3]string{"John", "Jane", "Alice"}
	fmt.Println(names)
	// 取得陣列長度
	fmt.Println(len(numbers))
	fmt.Println(len(names))
	// 逐一輸入陣列中的資料
	var grades [3]int
	for index := 0; index < len(grades); index++ {
		fmt.Scanln(&grades[index])
	}
	// 巡迴陣列中的資料
	var sum int
	for index := 0; index < len(grades); index++ {
		sum += grades[index]
	}
	fmt.Println(sum / len(grades))
}
