package calculator

// 对应java：开头小写为default，大写为public
var logMessage = "[LOG]"

var Version = "1.0"

func internalSum(number int) int {
	return number - 1
}

func Sum(num1, num2 int) int {
	return num1 + num2
}
