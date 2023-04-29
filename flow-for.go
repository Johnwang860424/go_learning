package main

import "fmt"

func main() {
	x := 0
	for x < 5 {
		fmt.Println("Value of x is: ", x)
		x++
	}
	for x := 0; x < 3; x++ {
		fmt.Println(x)
	}
	// 計算 1+2+3+...+50 的結果
	var total int
	for i := 1; i <= 50; i++ {
		total += i
		if i == 50 {
			fmt.Println("1累加到50的結果為", total)
		}
	}

	// 偶數不印出來
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	// 持續讓使用者輸入數字，計算總和。直到使用者輸入 0 為止
	var sum int
	for true {
		var input int
		fmt.Println("請輸入數字，輸入 0 結束")
		fmt.Scanln(&input)
		if input == 0 {
			break
		}
		sum += input
	}
	fmt.Println("總和為", sum)
}
