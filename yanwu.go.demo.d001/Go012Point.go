/**
@author <a herf="mailto:yanwu0527@163.com">XuBaofeng</a>
@date 2019/11/7 19:33.
<p>
description: 指针
	指针在GO语言中可以被拆分为两个核心概念
		1. 类型指针：允许对这个指针类型的数据进行修改，传递数据可以直接使用指针，而无需拷贝数据，类型指针不能进行偏移和运算
		2. 切片：由指向起始元素的原始指针、元素数量和容量组成
	GO语言中的指针类型变量即拥有指针高效访问的特点，又不会发生指针偏移，从未避免了非法修改关键性数据的问题
	切片比原始指针具有更强大的特性，而且更为安全
	一个指针变量可以指向任何一个值的内存地址，当一个指针被定义后没有分配到任何变量时，它的默认值为nil
	指针变量通常声明为：ptr := &v，其中v代表被取地址的变量，每个变量都拥有指针，指针的值就是地址
	new() 函数可以创建一个对应类型的指针，创建过程会分配内存，被创建的指针指向默认值
*/
package main

import (
	"flag"
	"fmt"
)

func main() {
	pointTest1()
	pointTest2()
	pointTest3()
	pointTest4()
}

func pointTest1() {
	cat, str := 1, "banana"
	fmt.Printf("%p %p \n", &cat, &str)
	house := "Malibu Point 10880, 90265"
	ptr := &house
	fmt.Printf("ptr type: %T\n", ptr) // ptr类型
	fmt.Printf("address: %p\n", ptr)  // ptr指针地址
	value := *ptr
	fmt.Printf("value type: %T\n", value) //指针取值后的类型
	fmt.Printf("value: %s\n", value)      // 指针取值后就是指向变量的值
}

func pointTest2() {
	x, y := 1, 2
	swap(&x, &y)
	fmt.Println(x, y)
}

func swap(a, b *int) {
	t := *a
	*a = *b
	*b = t
}

var mode = flag.String("mode", "", "process mode")

func pointTest3() {
	flag.Parse()
	fmt.Println(*mode)
}

func pointTest4() {
	str := new(string)
	*str = "go"
	fmt.Println(*str)
	fmt.Println(str)
}
