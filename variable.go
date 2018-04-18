package main

func test() (int, string) {
	return 1, "abcd"
}

func main() {

	println("------------------------")
	// 多变量赋值时，先计算所有相关值，然后再从左到右依次赋值。
	data, i := [3]int{0, 1, 2}, 0
	i, data[i] = 2, 100 // (i = 0) -> (i = 2), (data[0] = 100)
	println(data[0])
	println(i)

	println("------------------------")
	s := "abc"
	println(&s)
	s, y := "hello", 20 // 重新赋值: 与前 s 在同 层次的代码块中，且有新的变量被定义。
	println(&s, y)      // 通常函数多返回值 err 会被重复使 。
	{
		s, z := 1000, 30 // 定义新同名变量: 不在同 层次代码块。
		println(&s, z)
	}

	println("------------------------")
	// 特殊只写变量 "_"， 于忽略值占位。
	_, e := test()
	println(e)
}
