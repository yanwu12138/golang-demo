/**
@author <a herf="mailto:yanwu0527@163.com">XuBaofeng</a>
@date 2019/11/12 13:44.
<p>
description: 空值/零值
	nil是一个预先定义好的标识符，它是许多类型的零值表示
	nil标识符不能比较
	nil没有默认类型
	nil是map、slice、pointer、channel、func、interface的零值
	不同类型nil的指针是一样的
	不同类型nil的值占用内存大小可能是不一样的
*/
package main

import "fmt"

func main() {
	//fmt.Println(nil == nil)
	//fmt.Printf("%T", nil)
	//print(nil)
	var arr []int
	var num *int
	fmt.Printf("%p\n", arr)
	fmt.Printf("%p\n", num)
	yanwu := Yanwu{name: "yanwu", pawd: "123456", age: 28, sex: true}
	fmt.Printf("yanwu type: %T, yanwu value: %v", yanwu, yanwu)
}

type Yanwu struct {
	name string
	pawd string
	age  int
	sex  bool
}
