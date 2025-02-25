package main

import "fmt"

func main() {
	/*
		函数类型变量定义和使用

		定义：将其函数名赋值给变量，这里不需要加上()括号
			例如：var f1 = functionName
			或者 var f1 func(参数列表) (返回值列表)
			例如：var f1 func(a,b int) (res int, err error)

		函数加上() 是调用

		注意点：
			函数作为一种复合数据类型，可以看做是一种特殊的变量。
				函数名()：将函数进行调用，函数中的代码会全部执行，然后将return的结果返回给调用处
				函数名：指向函数体的内存地址
	*/
	// 定义int类型变量
	a := 10
	// 运算
	a += 5
	fmt.Println("a:", a)

	fmt.Println("===============")

	// 函数作为一个变量
	fmt.Printf("%T\n", fun1) // func(int, int)
	fmt.Println(fun1)        // 输出：0x100467070 看做函数名对应的函数体的地址

	// 函数变量
	var f1 func(a, b int)
	fmt.Println(f1) // <nil> 这里为空，也就是说堆内存的指向不存在

	//f1 = fun2 // Cannot use 'fun2' (type func(a int, b float64)) as the type func(a int, b int)

	f1 = fun1       // 将函数fun1的值（函数的地址）赋值给f1函数变量
	fmt.Println(f1) // 这里f1变量就指向fun1函数的地址

	// 调用fun1函数
	fun1(10, 20)
	// f1是函数变量，加上()也可以被当成函数使用。
	f1(15, 15)

	res1 := fun3       // 这里是将fun3函数的地址值赋值给res1，所以res1为函数变量
	res2 := fun3(1, 2) // 这里加上了()是调用函数，将函数的返回值赋值给了res2
	fmt.Println(res1)  // 输出的是fun3的地址值
	fmt.Println(res2)  // 函数的调用结果

	fmt.Println(res1(10, 2))
}

func fun1(a, b int) {
	fmt.Printf("a:%d, b:%d\n", a, b)
}

func fun2(a int, b float64) {

}

func fun3(a, b int) int {
	return a + b
}
