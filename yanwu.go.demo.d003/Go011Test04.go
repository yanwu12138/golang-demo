/**
@author <a herf="mailto:yanwu0527@163.com">XuBaofeng</a>
@date 2019/11/18 14:31.
<p>
description:
*/
package main

import "fmt"

func main() {
	arr := [...]int{1, 3, 7, 3, 87, 8, 2, 4, 67, 8, 5}
	fmt.Println("begin:", arr)
	len := len(arr)
	for i := 0; i < len; i++ {
		for j := i; j < len; j++ {
			if arr[i] > arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	fmt.Println("end", arr)
}
