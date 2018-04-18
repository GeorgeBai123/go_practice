package main

import "unsafe"

// 常量值还可以是 len、cap、unsafe.Sizeof 等编译期可确定结果的函数返回值。
// 在常量组中，如不提供类型和初始化值，那么视作与上 常量相同。

const (
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(b)
	s = "abc"
	x // x = "abc"
)

func main() {
	// const x = "xxx" // 未使 局部常量不会引发编译错误。
	println(e)
}
