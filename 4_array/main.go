package main

import "fmt"

func main() {
	// 创建数组
	/**
	var a [3]int
	a[1] = 10
	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[len(a)-1])
	*/

	// 初始化数组
	/**
	cities := [5]string{"asd", "dsa"}
	fmt.Println(cities)
	nums := []int{99: -1}
	fmt.Println(nums[99])  // -1
	fmt.Println(len(nums)) // 100
	*/

	// 切片，左闭右开
	// 切片的长度len： 切片中的元素数目。
	// 切片的容量cap： 切片开头与基础数组结束之间的元素数目。
	/**
	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	quarter1 := months[0:3]
	quarter2 := months[3:6]
	quarter3 := months[6:9]
	quarter4 := months[9:12]
	fmt.Println(quarter1, len(quarter1), cap(quarter1))
	fmt.Println(quarter2, len(quarter2), cap(quarter2))
	fmt.Println(quarter3, len(quarter3), cap(quarter3))
	fmt.Println(quarter4, len(quarter4), cap(quarter4))
	切片是动态数组，可以追加元素，每次扩容容量翻倍
	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("%d\tcap=%d\t%v\n", i, cap(numbers), numbers)
	}
	*/

	// 切片的复制（深拷贝）
	/**
	letters := []string{"A", "B", "C", "D", "E"}
	fmt.Println("Before", letters)

	slice1 := letters[0:2]

	slice2 := make([]string, 3)
	copy(slice2, letters[1:4])

	slice1[1] = "Z"

	fmt.Println("After", letters)
	fmt.Println("Slice2", slice2)
	*/

	// 哈希表
	/**
	studentAge := map[string]int{
		"john": 32,
		"bob":  31,
	}
	fmt.Println(studentAge)
	// 创建空哈新表
	studentAge = make(map[string]int)
	fmt.Println(studentAge)
	// 添加元素
	studentAge["Tom"] = 22
	fmt.Println(studentAge)
	// 判断元素是否存在
	age, exit := studentAge["asdasd"]
	fmt.Println(age, exit)
	// 删除元素
	delete(studentAge, "bob")
	fmt.Println(studentAge)
	// 遍历
	for name, age := range studentAge {
		fmt.Printf("%s\t%d\n", name, age)
	}
	*/

	// roman("MCLX")

	// 两个切片指向同一个基础数组，一个切片扩容后会指向新的数组，另一个不变
	/*
		s1 := []int{1, 2, 3}
		s2 := s1

		fmt.Printf("s1切片地址：%p\n", &s1)
		fmt.Printf("s2切片地址：%p\n", &s2)

		s1 = append(s1, 4, 5, 6)

		fmt.Printf("扩容后s1切片地址：%p\n", &s1)
		fmt.Printf("扩容后s2切片地址：%p\n", &s2)

		fmt.Println(s1)
		fmt.Println(s2)
	*/

	method()

}

func fibo(n int) {
	if n < 2 {
		panic("must > 2")
	}

	arr := make([]int, n)
	for i := 2; i < n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}

	fmt.Println(arr)
}

func roman(numStr string) {
	romanMap := map[rune]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}

	arabicVals := make([]int, len(numStr)+1)
	for index, digit := range numStr {
		if val, present := romanMap[digit]; present {
			arabicVals[index] = val
		} else {
			panic("error")
		}
	}

	total := 0
	for index := 0; index < len(numStr); index++ {
		if arabicVals[index] < arabicVals[index+1] {
			arabicVals[index] = -arabicVals[index]
		}
		total += arabicVals[index]
	}
	fmt.Println(total)
}
