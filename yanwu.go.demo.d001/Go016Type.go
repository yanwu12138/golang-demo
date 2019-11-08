/**
@author <a herf="mailto:yanwu0527@163.com">XuBaofeng</a>
@date 2019/11/8 11:57.
<p>
description: type关键字
	类型定义：type NewInt Type		[NewInt会形成一种新的类型，NewInt本身依然具备int类型的特性]
	类型别名：type TypeAlias = Type	[TypeAlias只是Type的别名，本质上是与Type同一个类型]
	非本地类型不能定义方法
*/
package main

import (
	"fmt"
	"reflect"
)

func main() {
	typeTest1()
	typeTest2()
}

// 将NewInt定义为int类型
type NewInt int

// 将int取一个别名叫IntAlias
type IntAlias = int

func typeTest1() {
	var a NewInt
	fmt.Printf("a type: %T\n", a)
	var b IntAlias
	fmt.Printf("b type: %T\n", b)
}

type Brand struct {
}

func (t Brand) show() {
	fmt.Println("---")
}

type FakeBrand = Brand

type Vehicle struct {
	FakeBrand
	Brand
}

func typeTest2() {
	var a Vehicle
	a.FakeBrand.show()
	ta := reflect.TypeOf(a)
	for i := 0; i < ta.NumField(); i++ {
		field := ta.Field(i)
		fmt.Printf("field: %v, type: %v\n", field.Name, field.Type.Name())
	}
}
