/**
@author <a herf="mailto:yanwu0527@163.com">XuBaofeng</a>
@date 2019/11/12 14:16.
<p>
description: 分支结构（if else if）
*/
package main

import "fmt"

func main() {
	if ten := getTen(); ten > 10 {
		fmt.Println("ten > 10")
	} else {
		fmt.Println("ten <= 10")
	}
}

func getTen() int {
	return 10
}
