/**
@author <a herf="mailto:yanwu0527@163.com">XuBaofeng</a>
@date 2019/11/8 15:34.
<p>
description: 切片，切片是对数组的一个连续片段的引用，所以切片是一个引用类型
	切片的内部结构包含地址、大小和容量
	切片一般用于快速的操作一块数据集合
	声明方式：
		1. 从数组或切片中生成
		2. 声明切片：var name []type
		3. 使用make()函数构造切片：make([]type, size, cap)
	在使用append()函数为切片追加元素时，当空间不足时，切片会以当前容量的2倍来进行扩容
*/
package main

import "fmt"

func main() {
	sliceInit()
	sliceAppend()
	sliceCopy()
	sliceDel()
	sliceTest()
}

func sliceInit() {
	// ===== 从数组生成新的切片
	a := [4]int{1, 2, 4, 8}
	fmt.Println(a, a[1:2])
	fmt.Println(a[:])
	fmt.Println(a[1:])
	fmt.Println(a[:2])
	// ===== 从切片生成新的切片
	b := a[:]
	fmt.Println(b, b[1:2])
	fmt.Println(b[:])
	fmt.Println(b[1:])
	fmt.Println(b[:2])
	fmt.Printf("%T %T \n", a, b)
	// ===== 重置切片
	c := b[0:0]
	fmt.Printf("%T %v \n", c, c)
	// ===== 声明切片：var name []type
	var m []string
	fmt.Println(m)
	// ===== 使用make()函数构造切片
	bools := make([]bool, 2, 4)
	fmt.Printf("%T %v \n", bools, bools)
}

func sliceAppend() {
	// ===== 使用append()函数为切片追加函数
	fmt.Println("--------------------------")
	var a []int
	a = append(a, 1)
	// ----- 在切片后追加元素
	fmt.Println("a len: ", len(a), "a cap: ", cap(a), "a value: ", a)
	a = append(a, 2, 3)
	fmt.Println("a len: ", len(a), "a cap: ", cap(a), "a value: ", a)
	a = append(a, []int{5, 6, 7}...)
	fmt.Println("a len: ", len(a), "a cap: ", cap(a), "a value: ", a)
	// ----- 在切片前填充元素
	a = append([]int{0}, a...)
	fmt.Println("a len: ", len(a), "a cap: ", cap(a), "a value: ", a)
	a = append([]int{-3, -2, -1}, a...)
	fmt.Println("a len: ", len(a), "a cap: ", cap(a), "a value: ", a)
	// ----- 在指定位置添加元素
	a = append(a[:8], append([]int{0}, a[8:]...)...)
	fmt.Println("a len: ", len(a), "a cap: ", cap(a), "a value: ", a)
	a = append(a[:5], append([]int{1, 2, 3}, a[5:]...)...)
	fmt.Println("a len: ", len(a), "a cap: ", cap(a), "a value: ", a)
}

func sliceCopy() {
	fmt.Println("--------------------------")
	ints1, ints2 := []int{1, 2, 3, 4, 5}, []int{6, 7, 8}
	// ----- ints2 copy into ints1
	copy(ints1, ints2)
	fmt.Println("ints1 len: ", len(ints1), "ints1 cap: ", cap(ints1), "ints1 value: ", ints1)
	fmt.Println("ints2 len: ", len(ints2), "ints2 cap: ", cap(ints2), "ints2 value: ", ints2)
	ints1, ints2 = []int{1, 2, 3, 4, 5}, []int{6, 7, 8}
	// ----- ints1 copy into ints2
	copy(ints2, ints1)
	fmt.Println("ints1 len: ", len(ints1), "ints1 cap: ", cap(ints1), "ints1 value: ", ints1)
	fmt.Println("ints2 len: ", len(ints2), "ints2 cap: ", cap(ints2), "ints2 value: ", ints2)
}

func sliceDel() {
	fmt.Println("--------------------------")
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	// ----- 删除开头N个元素
	a = a[3:]
	fmt.Println("a len: ", len(a), "a cap: ", cap(a), "a value: ", a)
	a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	// ----- 删除开头N个元素
	a = append(a[:0], a[3:]...)
	fmt.Println("a len: ", len(a), "a cap: ", cap(a), "a value: ", a)
	// ----- 删除开头N个元素
	a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	a = a[:copy(a, a[3:])]
	fmt.Println("a len: ", len(a), "a cap: ", cap(a), "a value: ", a)

}

func sliceTest() {
	fmt.Println("--------------------------")
	const elementCount = 1000
	srcData := make([]int, elementCount)
	for i := 0; i < elementCount; i++ {
		srcData[i] = i
	}
	refData := srcData
	copyData := make([]int, elementCount)
	copy(copyData, srcData)
	srcData[0] = 999
	fmt.Println(refData[0], copyData[0])
	fmt.Println(copyData[0], copyData[elementCount-1])
	// *****
	copy(copyData, srcData[4:6])
	for i := 0; i < 5; i++ {
		fmt.Println(copyData[i])
	}
}
