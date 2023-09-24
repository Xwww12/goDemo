// 每个可执行程序都应该是 `main` 包的一部分。
package main

// 导入的包必须使用
import (
	"fmt"
)

// 程序入口
func main() {
	fmt.Print("hello world")
}

// 指令
// go run 编译执行
// go build 编译生成可执行二进制文件
