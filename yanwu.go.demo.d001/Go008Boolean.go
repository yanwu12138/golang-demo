/**
@author <a herf="mailto:yanwu0527@163.com">XuBaofeng</a>
@date 2019/11/7 17:25.
<p>
description: 布尔
	go对值之间的比较有非常严格的限制，只有两个相同类型的值才可以进行比较，如果值的类型式接口，那么它们也必须都实现了相同的接口
	布尔值可以和 && (AND) || (OR) 操作符结合结合使用，并且有短路型为
	逻辑运算符中 && 的优先级比 || 的优先级高
	布尔无法参与数值运算，也无法与其他类型进行转换
*/
package main

import "fmt"

func main() {
	var a = 10
	fmt.Println(a == 5)
	fmt.Println(a == 10)
	fmt.Println(a != 5)
	fmt.Println(a != 10)
}
