/**
@author <a herf="mailto:yanwu0527@163.com">XuBaofeng</a>
@date 2019/11/11 19:18.
<p>
description: 字典 <K, V>
	字典是一种能够快速寻找之的理想结构
	声明方式：
		var name map[key_type]value_type
		name := make(map[key_type]value_type, [ cap ])
		name := map[key_type]value_type{}
	map可以根据K-V动态伸缩，因此它不存在固定长度或者最大限制
	map迭代
		for key, value := range map {}
		如果只需要迭代key：for key := range map {}
		如果只需要迭代value：for _, value := range map {}
	range创建了每个元素value的副本，而不是直接返回对该元素的value，这一特性与slice一致
	在并发的对map进行读和写会导致竟态问题，map内部会对这种并发操作进行检查并提前发现
	需要并发读写时，一般的做法是加锁，但这样的性能不高。所以GO提供了一种效率较高的并发安全的sync.Map
	sync.Map和map不同，它不是语言原生形态提供，而是sync报下的特殊结构
	sync.Map特性：
		1. 无需初始化，声明即可用
		2. sync.Map不能使用map的方式进行取值和设置等操作
		3. 使用range配合一个回调函数进行便利操作，通过回调函数内部遍历出来的值，range参数中回调函数的返回值在需要继续迭代时返回true，反之则返回false
		4. sync.Map的操作：Store（存储）、Load（获取）、Delete（删除）
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	mapInit()
	mapRange()
	mapDel()
	mapSync()
}

func mapInit() {
	fmt.Println("--------------------------")
	var mapTest1 map[string]int
	mapTest1 = map[string]int{"yanwu": 28, "litus": 26, "wenxin": 2, "wenfu": 1}
	fmt.Println("map size: ", len(mapTest1), "    map value: ", mapTest1)
	fmt.Println("yanwu: ", mapTest1["yanwu"])
	fmt.Println("litus: ", mapTest1["litus"])
	mapTest2 := map[string][]int{"yanwu": {1, 2}, "litus": {3, 4}, "wenxin": {5, 6}, "wenfu": {7}}
	fmt.Println("map size: ", len(mapTest2), "    map value: ", mapTest2)
}

func mapRange() {
	fmt.Println("--------------------------")
	mapTest := map[string]float32{"C0": 16.35, "D0": 18.35, "E0": 20.60, "F0": 21.83, "G0": 24.50, "A0": 27.50, "B0": 30.87, "A4": 440}
	fmt.Println("map size: ", len(mapTest), "    map value: ", mapTest)
	for key, value := range mapTest {
		fmt.Println("key: ", key, "    value: ", value)
		v := mapTest[key]
		fmt.Printf("Value: %d Value-Addr: %X ElemAddr: %X\n", value, &value, &v)
		delete(mapTest, key)
	}
	fmt.Println("map size: ", len(mapTest), "    map value: ", mapTest)
}

func mapDel() {
	fmt.Println("--------------------------")
	mapTest := map[string]float32{"C0": 16.35, "D0": 18.35, "E0": 20.60, "F0": 21.83, "G0": 24.50, "A0": 27.50, "B0": 30.87, "A4": 440}
	fmt.Println("map size: ", len(mapTest), "    map value: ", mapTest)
	delete(mapTest, "D0")
	fmt.Println("map size: ", len(mapTest), "    map value: ", mapTest)
}

func mapSync() {
	fmt.Println("--------------------------")
	mapTest := sync.Map{}
	mapTest.Store("yanwu", 28)
	mapTest.Store("lotus", 26)
	mapTest.Store("wenxin", 2)
	mapTest.Store("wenfu", 1)
	fmt.Println("map value: ", mapTest)
	mapTest.Range(func(key, value interface{}) bool {
		value2, _ := mapTest.Load(key)
		fmt.Printf("Value: %d Value-Addr: %X ElemAddr: %X\n", value, &value, &value2)
		if key == "wenxin" {
			return false
		} else {
			return true
		}
	})
}
