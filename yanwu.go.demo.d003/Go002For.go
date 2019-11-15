/**
@author <a herf="mailto:yanwu0527@163.com">XuBaofeng</a>
@date 2019/11/12 14:59.
<p>
description: 循环
*/
package main

import (
	"fmt"
)

func main() {
	forTest1()
	forTest2()
	forTest3()
	forTest4()
	forTest5(0)
	forTest6()
}

func forTest1() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func forTest2() {
	sum := 0
	for {
		sum++
		if sum > 10 {
			break
		}
	}
	fmt.Println(sum)
}

func forTest3() {
	for j := 0; j < 5; j++ {
		if j == 2 {
			continue
		}
		for i := 0; i < 10; i++ {
			if i == 5 {
				break
			}
			fmt.Println(j, i)
		}
	}
}

func forTest4() {
	step := 2
	for ; step > 0; step-- {
		fmt.Println(step)
	}
}

func forTest5(i int) {
	for i <= 10 {
		i++
		fmt.Println(i)
	}
}

func forTest6() {
	// 遍历, 决定处理第几行
	for y := 1; y <= 9; y++ {
		// 遍历, 决定这一行有多少列
		for x := 1; x <= y; x++ {
			fmt.Printf("%d*%d=%d ", x, y, x*y)
		}
		// 手动生成回车
		fmt.Println()
	}
}
