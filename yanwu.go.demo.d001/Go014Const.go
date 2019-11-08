/**
@author <a herf="mailto:yanwu0527@163.com">XuBaofeng</a>
@date 2019/11/8 10:40.
<p>
description: 常量
	常量使用const关键字定义，用于存储不会改变的数据
	常量是在编译时被创建的，及时定义在函数内部也是如此
	常量只能时布尔、数字（整数、浮点数、复数）、字符串
	格式：const 变量名 [变量类型] = 变量值
	变量的值或运算都必须是在编译时就能够确定的
	无类型常量：编译器为没有明确的基础类型的数字常量提供比基础类型更高经度的算术运算，可以认为至少有256bit的运算精度
	go中有六种未明确类型的常量类型：无类型布尔型、无类型整型、无类型字符、无类型浮点、无类型复数、无类型的字符串
*/
package main

import (
	"fmt"
)

func main() {
	constTest1()
	constTest2()
}

func constTest1() {
	const con1 int = 17
	const con2 = 17
	fmt.Println(con1 == con2)
	const (
		con3 = 18
		con4
		con5 = 19
	)
	fmt.Printf("%v %T \n", con3, con4)
}

type Week int

func constTest2() {
	const (
		sun Week = iota
		mon
		tue
		wed
		thu
		fri
		sat
	)
	fmt.Printf("%v %v %v %v %v %v %v \n", sun, mon, tue, wed, thu, fri, sat)
}
