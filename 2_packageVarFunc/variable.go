package main

import "fmt"

// 创建变量
func createVar1() {
	// 声明一个变量
	var age int
	// 声明多个变量
	var firstName, lastName string
	// 把var提取出来
	var (
		aaa string
		bbb int
	)

	fmt.Print(age, firstName, lastName, aaa, bbb)
}

// 初始化值
func createVar2() {
	// 可以不写类型，会自动推断
	var (
		firstName = "x"
		lastName  = "x"
		age       = 10
	)
	// 省略var和类型，只能在函数里用
	ccc := "asd"
	// 声明常量，可以不指定类型
	const HTTPStatusOK = 200

	fmt.Print(age, firstName, lastName, ccc)
}
