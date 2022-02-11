package main

func main() {

}

// func main() {
// 	// 匿名函数
// 	f := func(x, y int) int {
// 		return x + y
// 	}

// 	r1 := f(3, 4)
// 	fmt.Println(r1)

// 	r2 := func(x, y int) int {
// 		return x + y
// 	}(4, 5)

// 	fmt.Println(r2)

// }

// func main() {
// 	f()
// }

// // 匿名函数的内存地址
// func f() {
// 	for i := 0; i < 4; i++ {
// 		g := func(i int) { fmt.Printf("%d ", i) } //此例子中只是为了演示匿名函数可分配不同的内存地址，在现实开发中，不应该把该部分信息放置到循环中。
// 		g(i)
// 		fmt.Printf(" - g is of type %T and has value %v\n", g, g)
// 	}
// }

//	0  - g is of type func(int) and has value 0x10a49a0
// 1  - g is of type func(int) and has value 0x10a49a0
// 2  - g is of type func(int) and has value 0x10a49a0
// 3  - g is of type func(int) and has value 0x10a49a0
