/**
@author <a herf="mailto:yanwu0527@163.com">XuBaofeng</a>
@date 2019/11/18 19:25.
<p>
description: 函数
	在函数中，实参通过值传递的形式进行传递，因此函数的形参是实参的拷贝，对形参的变更不会影响到实参
	返回值中的最后一个返回参数通常情况下返回函数执行中可能发生的错误
*/
package main

import "fmt"

func main() {
	i1, i2 := funcTest01("1", "2")
	fmt.Println(i1)
	fmt.Println(i2)

	i1, i2 = funcTest02()
	fmt.Println(i1)
	fmt.Println(i2)
}

func funcTest01(str1, str2 string) (int, int) {
	return 1, 2
}

func funcTest02() (a, b int) {
	a = 3
	b = 4
	return
}
