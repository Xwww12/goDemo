package main

// 传指针& 参数为指针* 取指针指向的值*
func updateName(name *string) {
	*name = "David"
}
