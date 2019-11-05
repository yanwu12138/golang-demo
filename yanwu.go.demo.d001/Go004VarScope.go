/**
@author <a herf="mailto:yanwu0527@163.com">XuBaofeng</a>
@date 2019/11/5 18:01.
<p>
description: 变量的作用域
	1. 局部变量
		在函数体内声明的变量为局部变量，它的作用域仅在函数内
		参数和返回值变量也是局部变量
	2. 全局变量
		在函数体外声明的变量为全局变量，全局变量能在整个包与外部包（被导出后）使用
		全局变量与局部变量名称可以相同，但函数体内的局部变量会被优先考虑（就近原则）
	3. 形式参数
		形式参数会作为函数的局部变量来使用
*/
package main

import "fmt"

func main() {
	fmt.Println("========== var scope ==========")
	scopeTest1()
	scopeTest2()
	scopeTest3()
}

func scopeTest1() {
	fmt.Println("---------- 作用域1：局部变量 ----------")
	a, b := 3, 4
	c := a * b
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("c: ", c)
}

var d, e int

func scopeTest2() {
	fmt.Println("---------- 作用域2：全局变量 ----------")
	fmt.Println("e: ", e)
	a, b := 3, 4
	d = a - b
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("d: ", d)
	e := "string"
	fmt.Println("e: ", e)
}

func scopeTest3() {
	fmt.Println("---------- 作用域3：形式参数 ----------")
	a, b := 3, 4
	c := sum(a, b)
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("c: ", c)
}

func sum(a, b int) int {
	return a + b
}
