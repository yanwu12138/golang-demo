/**
@author <a herf="mailto:yanwu0527@163.com">XuBaofeng</a>
@date 2019/11/12 15:40.
<p>
description: goto通过标签进行代码间的无条件跳转
*/
package main

import "fmt"

func main() {
	gotoTest1()
	gotoTest2()
}

func gotoTest1() {
	fmt.Println("--------------------")
	var breakAgain bool
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			if y == 5 {
				breakAgain = true
				break
			}
			fmt.Println(x, y)
		}
		if breakAgain {
			break
		}
	}
}

func gotoTest2() {
	fmt.Println("--------------------")
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			if y == 5 {
				goto breakHere
			}
			fmt.Println(x, y)
		}
	}
	fmt.Println("--------------------")
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			if y > 500 {
				goto breakHere
			}
			fmt.Println(x, y)
		}
	}
breakHere:
	fmt.Println("done")
}
