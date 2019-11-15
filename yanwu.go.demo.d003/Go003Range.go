/**
@author <a herf="mailto:yanwu0527@163.com">XuBaofeng</a>
@date 2019/11/12 15:18.
<p>
description: 迭代
*/
package main

import (
	"fmt"
)

func main() {
	rangeArray()
	rangeSlice()
	rangeString()
	rangeMap()
	rangeChannel()
}

func rangeArray() {
	fmt.Println("--------------------")
	array := []int{1, 2, 3, 4, 5}
	for index, value := range array {
		fmt.Printf("index: %d  value: %d \n", index, value)
	}
}

func rangeSlice() {
	fmt.Println("--------------------")
	slice := make([]int, 2, 4)
	for index, value := range slice {
		fmt.Printf("index: %d  value: %d \n", index, value)
	}
}

func rangeString() {
	fmt.Println("--------------------")
	str := "hello, 你好！"
	for index, value := range str {
		fmt.Printf("index: %d value: 0x%x \n", index, value)
	}
}

func rangeMap() {
	fmt.Println("--------------------")
	m := map[string]int{"yw": 98, "sx": 99, "yy": 61}
	for key, value := range m {
		fmt.Println(key, value)
	}
}

func rangeChannel() {
	fmt.Println("--------------------")
	c := make(chan int)
	go func() {
		c <- 1
		c <- 2
		c <- 3
		c <- 4
	}()
	for v := range c {
		fmt.Println(v)
	}
}
