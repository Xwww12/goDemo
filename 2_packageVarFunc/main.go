package main

import (
	"fmt"
	"goDemo/calculator"
)

func main() {
	// 给main函数添加参数
	// go run main.go 3 5
	//number1, _ := strconv.Atoi(os.Args[1])
	//number2, _ := strconv.Atoi(os.Args[2])
	//fmt.Println("Sum:", number1+number2)

	// 传入指针
	firstName := "John"
	updateName(&firstName)
	fmt.Println(firstName)

	// 导入自定义模块
	fmt.Println(calculator.Sum(1, 2))
}
