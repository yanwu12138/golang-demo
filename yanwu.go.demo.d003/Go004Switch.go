/**
@author <a herf="mailto:yanwu0527@163.com">XuBaofeng</a>
@date 2019/11/12 15:32.
<p>
description: 选择语句
	case按照从上到下的顺序进行求值，直到找到匹配的项
*/
package main

import "fmt"

func main() {
	switchTest1()
	switchTest2()
	switchTest3(61)
}

func switchTest1() {
	a := "ha"
	switch a {
	case "hello":
		fmt.Println(1)
	case "ha":
		fmt.Println(2)
	default:
		fmt.Println(0)
	}
}

func switchTest2() {
	a := "ha"
	switch a {
	case "hello", "ha":
		fmt.Println(1)
	default:
		fmt.Println(0)
	}
}

func switchTest3(r int) {
	switch {
	case r > 0 && r < 60:
		fmt.Println("c")
	case r >= 60 && r < 80:
		fmt.Println("b")
	case r >= 80 && r <= 100:
		fmt.Println("a")
	default:
		fmt.Println("error")
	}
}
