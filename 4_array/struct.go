package main

import "fmt"

// Employee 声明结构体
type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

func method() {
	// 使用结构体
	var john Employee
	fmt.Println(john)

	// 声明并初始化结构体
	employee := Employee{1001, "a", "b", "c"}
	fmt.Println(employee)
	employee.ID = 1002
	fmt.Println(employee)
}
