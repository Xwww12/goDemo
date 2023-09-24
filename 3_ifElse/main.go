package main

import (
	"fmt"
)

func main() {

	// if 语句中赋值，作用域为if块中
	/*x := 10
	if x%2 == 0 {
		fmt.Println(x)
	}
	if num := givemeanumber(); num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has only one digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}*/

	// switch，不需要break
	/*sec := time.Now().Unix()
	rand.Seed(sec)
	i := rand.Int31n(10)
	switch i {
	case 0:
		fmt.Print("zero...")
	case 1:
		fmt.Print("one...")
	case 2:
		fmt.Print("two...")
	case 3, 4:
		fmt.Println("three or four...")
	default:
		fmt.Println("no match...")
	}
	fmt.Println("ok")*/

	// for
	/*sum := 0
	for i := 0; i < 100; i++ {
		sum += i
	}
	fmt.Println(sum)*/

	// defer：在for里使用到for结束后再执行，倒序执行
	// 在普通方法里使用，会在方法执行结束后执行，可用来关闭资源等
	for i := 0; i < 4; i++ {
		defer fmt.Println(-i)
		fmt.Println(i)
	}
	/*
		打印顺序
		0
		1
		2
		3
		-3
		-2
		-1
		0
	*/

	// panic 会中断方法执行，但defer的代码会继续执行，
	// 到执行完defer后panic("")里的内容会打印，最后打印错误栈
}

func givemeanumber() int {
	return -1
}
