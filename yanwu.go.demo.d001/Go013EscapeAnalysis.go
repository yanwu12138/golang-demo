/**
@author <a herf="mailto:yanwu0527@163.com">XuBaofeng</a>
@date 2019/11/8 10:18.
<p>
description: 变量逃逸分析：通过编译器分析代码的特征合代码的生命周期，决定使用堆还是栈类进行内存分配
*/
package main

import "fmt"

func main() {
	analysisTest1()
	analysisTest2()
}

/**
逃逸分析
go run -gcflags "-m -l" Go013EscapeAnalysis.go
*/
func analysisTest1() {
	var a int
	void()
	fmt.Println(a, dummy(0))
}

func dummy(arg int) int {
	var result int
	result = arg
	return result
}

func void() {

}

func analysisTest2() {
	fmt.Println(dummy1())
}

type Data struct {
}

func dummy1() *Data {
	var c Data
	return &c
}
