/**
@author <a herf="mailto:yanwu0527@163.com">XuBaofeng</a>
@date 2019/11/12 19:52.
<p>
description: 二公查找
*/
package main

import "fmt"

func main() {
	arr := []int{1, 2, 5, 7, 15, 25, 30, 36, 39, 51, 67, 78, 80, 82, 85, 91, 92, 97}
	index := binaryFind(&arr, 0, len(arr)-1, 30)
	fmt.Println("main arr: ", arr, index)
}

func binaryFind(arr *[]int, leftIndex, rightIndex, findVal int) int {
	if leftIndex > rightIndex {
		fmt.Println("error")
		return -1
	}
	middle := (leftIndex + rightIndex) / 2
	if (*arr)[middle] > findVal {
		binaryFind(arr, leftIndex, middle-1, findVal)
	} else if (*arr)[middle] < findVal {
		binaryFind(arr, middle+1, rightIndex, findVal)
	} else {
		fmt.Println("find: ", middle)
	}
	return middle
}
