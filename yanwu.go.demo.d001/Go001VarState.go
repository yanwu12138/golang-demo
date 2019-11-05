/**
@author <a herf="mailto:yanwu0527@163.com">XuBaofeng</a>
@date 2019/11/5 17:17.
<p>
description: go的变量声明的几种方式：
1. 标准格式
	var 变量名 变量类型 [ = 变量值 ]
2. 批量格式
	var (
		变量名 变量类型 [ = 变量值 ]
		变量名 变量类型 [ = 变量值 ]
	)
3. 简短格式
	变量名 := 变量值
	变量名, 变量名 := 变量值, 变量值
*/
package main

import "fmt"

func main() {
	// women
	fmt.Println("========== var state ==========")
	stateTest1()
	stateTest2()
	stateTest3()
}

func stateTest1() {
	fmt.Println("---------- 方式1：标准格式 ----------")
	var str1 string = "hello world"
	var str2 string = "str2"
	fmt.Println("str1: -> ", str1)
	fmt.Println("str2: -> ", str2)
}

func stateTest2() {
	fmt.Println("---------- 方式2：批量格式 ----------")
	var (
		a int
		b string = "str"
		c []float32
		d func() bool
		e struct {
			x int
		}
	)
	fmt.Println("a: -> ", a)
	fmt.Println("b: -> ", b)
	fmt.Println("c: -> ", c)
	fmt.Println("d: -> ", d)
	fmt.Println("e: -> ", e)
}

func stateTest3() {
	fmt.Println("---------- 方式3：简短格式 ----------")
	boo1 := false
	fmt.Println("boo1: -> ", boo1)
	boo2, boo3 := true, false
	fmt.Println("boo2: -> ", boo2)
	fmt.Println("boo3: -> ", boo3)
}
