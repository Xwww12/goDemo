package main

import "fmt"

// 练习在Go中使用控制流
func main() {
	// method1()

	//for number := 1; number <= 20; number++ {
	//	if findprimes(number) {
	//		fmt.Printf("%v ", number)
	//	}
	//}

	method3()
}

func method1() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func findprimes(number int) bool {
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}

	if number > 1 {
		return true
	} else {
		return false
	}
}

func method3() {
	for {
		val := 0
		fmt.Print("Enter number: ")
		fmt.Scanf("%d", &val)
		fmt.Scanf("%s")
		if val < 0 {
			panic("")
		} else if val == 0 {
			fmt.Println("0 is neither negative nor positive")
		} else {
			fmt.Println("You entered:", val)
		}
	}
}
