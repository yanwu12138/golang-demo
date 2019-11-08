/**
@author <a herf="mailto:yanwu0527@163.com">XuBaofeng</a>
@date 2019/11/8 15:14.
<p>
description: 数组：由一个固定长度的特定类型元素组成的序列
	一个数组可以由零个或多个元素组成
	声明格式：
		var 变量名 [长度]类型 >> 数组的每个元素都会被初始化为元素对应的零值
		var 变量名 = [长度]类型{值、值...}
		变量名 := [...]类型{值、值...}
*/
package main

import "fmt"

func main() {
	// ===== 格式1
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])
	// ----- 打印索引和元素
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
	// ----- 仅打印元素
	for _, v := range a {
		fmt.Printf("%d \n", v)
	}
	// ===== 格式2
	var q = [3]int{1, 2}
	fmt.Println(q[0])
	fmt.Println(q[len(q)-1])
	// ===== 格式3
	m := [...]string{"yanwu", "lotus", "wenxin", "wenfu"}
	fmt.Println(len(m))
	// ===== 判断两个数组是否相等
	x, y, z := [2]int{1, 2}, [...]int{1, 2}, [2]int{1, 3}
	fmt.Println(x == y)
	fmt.Println(x == z)
	// ===== 数组遍历
	team := [...]string{"hammer", "soldier", "mum"}
	for index, value := range team {
		fmt.Println(index, value)
	}
	// ===== 二维数组
	var array [4][2]int
	array = [4][2]int{{10, 11}, {1, 2}, {30, 21}, {5, 17}}
	fmt.Println(array)
	array = [4][2]int{1: {2, 7}}
	fmt.Println(array)
	array = [4][2]int{1: {0: 20}, 3: {1: 41}}
	fmt.Println(array)
}
