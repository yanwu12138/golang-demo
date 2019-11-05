/**
@author <a herf="mailto:yanwu0527@163.com">XuBaofeng</a>
@date 2019/11/5 17:17.
<p>
description: go变量初始化的几种形式：
1. 标准格式
	var 变量名 变量类型 = 变量值
2. 推导格式
	var 变量名 = 变量值
	var 变量名, 变量名 = 变量值, 变量值
3. 简短格式
	变量名 := 变量值
	变量名, 变量名 := 变量值, 变量值
4. 批量格式
	var (
		变量名 变量类型 = 变量值
		变量名 变量类型 = 变量值
	)
*/
package main

import "fmt"

func main() {
	fmt.Println("========== var init ==========")
	initTest1()
	initTest2()
	initTest3()
	initTest4()
}

func initTest1() {
	fmt.Println("---------- 方式1：标准格式 ----------")
	var str1 string = "hello world"
	fmt.Println("str1: -> ", str1)
}

func initTest2() {
	fmt.Println("---------- 方式2：推导格式 ----------")
	var hp = 100
	fmt.Println("ph: -> ", hp)
	var int1, str2 = -80, "str2"
	fmt.Println("int1: -> ", int1)
	fmt.Println("str2: -> ", str2)
}

func initTest3() {
	fmt.Println("---------- 方式3：简短格式 ----------")
	flag := false
	fmt.Println("flag: -> ", flag)
	int1, str1 := 100, "str1"
	fmt.Println("int1: -> ", int1)
	fmt.Println("str1: -> ", str1)
}

func initTest4() {
	fmt.Println("---------- 方式4：批量格式 ----------")
	var (
		a int    = 123
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
