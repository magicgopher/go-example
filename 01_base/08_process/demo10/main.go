package main

import "fmt"

func main() {
	// fallthrough 关键字用于强制执行下一个 case 的代码块
	// 即使下一个 case 的条件不满足。需要注意的是，fallthrough 只能穿透一层
	// 即只能跳转到下一个 case，不能连续穿透多个 case

	// 使用 fallthrough 的注意事项
	// fallthrough 关键字必须放在 case 语句块的最后一行，否则会导致编译错误。
	// fallthrough 关键字会无条件地执行下一个 case 的代码块，即使下一个 case 的条件不满足。
	// fallthrough 可能会导致代码逻辑混乱，因此应该谨慎使用。
	switch i := 1; i {
	case 1:
		fmt.Println("case 1")
		fallthrough // 穿透到下一个 case
	case 2:
		fmt.Println("case 2")
		fallthrough // 穿透到下一个 case
	case 3:
		fmt.Println("case 3")
	default:
		fmt.Println("default case")
	}
}
