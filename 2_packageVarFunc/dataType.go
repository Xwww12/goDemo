package main

import (
	"fmt"
	"math"
	"strconv"
)

/*
基本类型：数字、字符串和布尔值
聚合类型：数组和结构
引用类型：指针、切片、映射、函数和通道
接口类型：接口
*/
func integer() {
	// 整数
	var integer8 int8 = 127
	var integer16 int16 = 32767
	var integer32 int32 = 2147483647
	var integer64 int64 = 9223372036854775807
	var i int = 1 // 大小至少为32位以上
	fmt.Println(integer8, integer16, integer32, integer64, i)

	// 通过math获取各种类型值的最大最小值
	fmt.Println(math.MaxFloat32)
}

func boolean() {
	var featureFlag bool = true

	fmt.Println(featureFlag)
}

func String() {
	var firstName string = "John"
	lastName := "Doe"
	fmt.Println(firstName, lastName)
}

func convert() {
	// 类型转换
	i, _ := strconv.Atoi("-42") // ASCII to Int
	s := strconv.Itoa(-42)      // Int to ASCII
	fmt.Println(i, s)
}
