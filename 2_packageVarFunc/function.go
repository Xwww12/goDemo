package main

import "strconv"

// 指定一个返回值
func sum(num1 string, num2 string) int {
	i1, _ := strconv.Atoi(num1)
	i2, _ := strconv.Atoi(num2)
	return i1 + i2
}

// 指定多个返回值，返回值必须都得接住，不用的可以用 _ 忽略
func calc(number1 string, number2 string) (sum int, mul int) {
	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)
	sum = int1 + int2
	mul = int1 * int2
	return
}
