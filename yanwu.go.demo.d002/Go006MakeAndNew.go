/**
@author <a herf="mailto:yanwu0527@163.com">XuBaofeng</a>
@date 2019/11/12 14:02.
<p>
description: make & new
	make: 初始化内置的数据结构（数组、切片、channel）
	new: 获取指向某个类型的指针
*/
package main

import (
	"fmt"
)

func main() {
	makeTest1()
	newTest1()
}

func makeTest1() {
	fmt.Println("--------------------")
	slice := make([]int, 0, 100)
	hash := make(map[int]bool, 10)
	ch := make(chan int, 5)
	fmt.Printf("slice type: %T, slice size: %v, slice cap: %v, slice value: %v \n", slice, len(slice), cap(slice), slice)
	fmt.Printf("hash type: %T, hash size: %v, hash value: %v \n", hash, len(hash), hash)
	fmt.Printf("ch type: %T, ch size: %v, ch cap: %v, ch value: %v \n", ch, len(ch), cap(ch), ch)
}

func newTest1() {
	fmt.Println("--------------------")
	i := new(int)
	var v int
	j := &v
	fmt.Printf("i type: %T, i value: %v \n", i, i)
	fmt.Printf("j type: %T, j value: %v \n", j, j)
}
