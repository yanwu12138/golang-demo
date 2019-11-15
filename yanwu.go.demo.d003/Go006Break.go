/**
@author <a herf="mailto:yanwu0527@163.com">XuBaofeng</a>
@date 2019/11/12 15:48.
<p>
description: break语句可以结束for、switch和select的代码块
	另外break语句还可以在语句后面添加标签，表示退出某个标签对应的代码块，标签要求必须定义在对应的 for、switch 和 select 的代码块上。
*/
package main

import "fmt"

func main() {
OuterLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			switch j {
			case 2:
				fmt.Println(i, j)
				break OuterLoop // ===== 运行到这里退出
			case 3:
				fmt.Println(i, j)
				break OuterLoop
			}
		}
	}
}
