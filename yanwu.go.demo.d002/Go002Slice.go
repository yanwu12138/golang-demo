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
	当迭代切片时，range关键字会返回两个值，第一个值是当前迭代到的索引位置，第二个值是对应元素值的一份副本
	range创建了每个元素的副本，而不是直接返回对该元素的引用
	由于range迭代返回的变量是在迭代过程中根据切片以此赋值的新变量，所以value的地址总是相同的
	要想获取每个元素的地址，需要使用索引：arr[index]
*/
package main

import "fmt"

func main() {
	sliceInit()
	sliceAppend()
	sliceCopy()
	sliceDel()
	superSlice()
	sliceRange()
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
	// ----- 从中间开始删除
	a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	a = append(a[:3], a[5:]...)
	fmt.Println("a len: ", len(a), "a cap: ", cap(a), "a value: ", a)
	// ----- 从中间开始删除
	a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	a = a[:3+copy(a[3:], a[5:])]
	fmt.Println("a len: ", len(a), "a cap: ", cap(a), "a value: ", a)
	// ----- 从尾部删除
	a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	a = a[:len(a)-3]
	fmt.Println("a len: ", len(a), "a cap: ", cap(a), "a value: ", a)
}

func superSlice() {
	fmt.Println("--------------------------")
	a := [][]int{{101, 102}, {1, 2, 3, 4, 5, 6, 7, 8, 9, 0}}
	// ----- 为第一个切片追加元素
	a[0] = append(a[0], 103, 104)
	fmt.Println("a len: ", len(a), "a cap: ", cap(a), "a value: ", a)
}

func sliceRange() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	for index, value := range a {
		fmt.Println("index: ", index, "    value: ", value)
		fmt.Printf("Value: %d Value-Addr: %X ElemAddr: %X\n", value, &value, &a[index])
	}
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
