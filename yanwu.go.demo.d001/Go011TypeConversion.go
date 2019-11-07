/**
@author <a herf="mailto:yanwu0527@163.com">XuBaofeng</a>
@date 2019/11/7 19:22.
<p>
description: 类型转换
	在可行的情况下，一个类型的值是可以被转换成另一种类型的值
	转换公式：变量B := 类型(变量A)
	类型转换只能在定义正确的情况下转换成，例如从一个取值范围小的类型转换到一个取值范围较大的类型，否则有可能会丢失精度
	只有相同地层类型的变量之间可以相互转换
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("int8 range: ", math.MinInt8, math.MaxInt8)
	fmt.Println("int16 range: ", math.MinInt16, math.MaxInt16)
	fmt.Println("int32 range: ", math.MinInt32, math.MaxInt32)
	fmt.Println("int64 range: ", math.MinInt64, math.MaxInt64)

	var a int32 = 1047483647
	fmt.Printf("int32: 0x%x %d\n", a, a)
	b := int16(a)
	fmt.Printf("int32: 0x%x %d\n", b, b)

	var c float32 = math.Pi
	fmt.Println(int(c))
}
