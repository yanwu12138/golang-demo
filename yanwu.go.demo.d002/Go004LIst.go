/**
@author <a herf="mailto:yanwu0527@163.com">XuBaofeng</a>
@date 2019/11/11 20:08.
<p>
description: 列表
	列表时一种非连续的存储容器，由多个节点组成，节点通过一些变量记录彼此之间的关系，由单链表、多链表等实现方式
	声明方式：
		name := list.New()
		var name list.List
	列表并没有具体元素类型的限制，它的元素可以是任意类型
	迭代：
		for item := listTest.Front(); item != nil; item = item.Next() {}
*/
package main

import (
	"container/list"
	"fmt"
)

func main() {
	listInit()
	listDel()
	listRange()
}

func listInit() {
	fmt.Println("--------------------------")
	var listTest list.List
	listTest.PushBack(28)
	element := listTest.PushFront("yanwu")
	listTest.InsertBefore("lotus", element)
	fmt.Println("list size: ", listTest.Len(), "    list value: ", listTest)
}

func listDel() {
	fmt.Println("--------------------------")
	listTest := list.New()
	listTest.PushBack("yuwen")
	listTest.PushBack("shuxue")
	element := listTest.PushBack("yingyu")
	listTest.PushBack("huaxue")
	fmt.Println("list size: ", listTest.Len(), "    list value: ", listTest)
	listTest.Remove(element)
	fmt.Println("list size: ", listTest.Len(), "    list value: ", listTest)
}

func listRange() {
	fmt.Println("--------------------------")
	listTest := list.New()
	listTest.PushBack("yuwen")
	listTest.PushBack("shuxue")
	element := listTest.PushBack("yingyu")
	listTest.InsertAfter("huaxue", element)
	for item := listTest.Front(); item != nil; item = item.Next() {
		fmt.Println(item.Value)
	}
}
