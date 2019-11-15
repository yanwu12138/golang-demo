/**
@author <a herf="mailto:yanwu0527@163.com">XuBaofeng</a>
@date 2019/11/12 15:53.
<p>
description: continue
	continue语句可以结束当前循环，开始下一次循环
*/
package main

import "fmt"

func main() {
outerLoop:
	for i := 0; i < 10; i++ {
		switch i {
		case 5:
			continue outerLoop
		}
		fmt.Println(i)
	}
}
