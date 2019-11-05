/**
@author <a herf="mailto:yanwu0527@163.com">XuBaofeng</a>
@date 2019/11/5 17:50.
<p>
description: 匿名变量
	1. 匿名变量不占用内存空间
	2. 匿名变量不会分配内存
	3. 匿名变量与匿名变量之间也不会因为多次声明而无法使用
	4. 匿名变量用标识符 _ 表示
	5. 任何类型都可以赋值给 _
	6. 任何赋给标识符 _ 的值都将被抛弃，因此这些值不能在后续的代码中使用，也不可以使用这个标识符作为变量对其它变量进行赋值或运算
*/
package main

import "fmt"

func main() {
	a, _ := getData()
	_, b := getData()
	fmt.Println("a: -> ", a)
	fmt.Println("b: -> ", b)
}

func getData() (int, int) {
	return 100, 200
}
