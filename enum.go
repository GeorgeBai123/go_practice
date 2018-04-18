package main

type Color int

const (
	Black Color = iota
	Red
	Blue
)

const (
	A = iota //0
	B        //1
	C = "c"  //c
	D        //c，与上  相同。
	E = iota //4，显式恢复。注意计数包含了 C、D 两 。
	F        //5
)

// 在同 常量组中，可以提供多个 iota，它们各 增 。
const (
	_        = iota             // iota = 0
	KB int64 = 1 << (10 * iota) //iota=1
	MB                          // 与 KB 表达式相同，但 iota = 2
	GB
	TB
)
const (
	G, H = iota, iota << 10 //0,0<<10
	L, J                    // 1, 1 << 10
)

func test(c Color) {
	println(c)
}
func main() {
	c := Black
	test(c)
	// x := 1
	// test(x) // Error: cannot use x (type int) as type Color in function argument
	test(1) // 常量会被编译器 动转换。
	println(F)
	println("-------------------")
	println(KB)
	println(MB)
	println(GB)
	println(TB)
	println("-------------------")
	println(G)
	println(H)
	println(L)
	println(J)

}
