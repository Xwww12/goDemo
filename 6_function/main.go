package main

import (
	"fmt"
	"strings"
)

type triangle struct {
	size int
}

type square struct {
	size int
}

// 给已有的类型取别名
type upperstring string

// Shape 接口
type Shape interface {
	Perimeter() float64
	Area() float64
}

// 实现接口的所有方法就是实现接口
type Square struct {
	size float64
}

func (s Square) Area() float64 {
	return s.size * s.size
}

func (s Square) Perimeter() float64 {
	return s.size * 4
}
func main() {
	t := triangle{3}
	s := square{4}
	fmt.Println("Perimeter (triangle):", t.perimeter())
	fmt.Println("Perimeter (square):", s.perimeter())

	fmt.Println("Size:", t.size)
	fmt.Println("Perimeter:", t.perimeter())
}

// 使用这个方法必须通过方法体
func (t triangle) perimeter() int {
	return t.size * 3
}

// 通过指针能够改变实例的值
func (t *triangle) doubleSize() {
	t.size *= 2
}

func (s square) perimeter() int {
	return s.size * 4
}

// Upper 给已有类型添加方法
func (s upperstring) Upper() string {
	return strings.ToUpper(string(s))
}
