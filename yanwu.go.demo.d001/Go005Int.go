/**
@author <a herf="mailto:yanwu0527@163.com">XuBaofeng</a>
@date 2019/11/7 16:11.
<p>
description: 整数
	整数包含int8、int16、int32、int64等，每种数值类型都决定了对应的大小范围和是否支持正负符号，其中后面的数字为整数的二进制长度
	有符号包括：int8、 int16、 int32、 int64， 有符号整数的取值范围：-2^(n-1) ~ 2^(n-1)-1
	无符号包括：uint8、uint16、uint32、uint64，无符号整数的取值范围：0 ~ 2^n-1
	通常情况下，我们只需要用到int就可以了，而且通常int类型的处理速度也是最快的

*/
package main

import "fmt"

func main() {
	intTest()
	int8Test()
	uint8Test()
	int16Test()
	uint16Test()
	int32Test()
	uint32Test()
	int64Test()
	uint64Test()
}

func intTest() {
	a, b := -4294967296, 4294967296
	fmt.Println("int a: -> ", a)
	fmt.Println("int b: -> ", b)
}

func int8Test() {
	min, max := int8(-128), int8(127)
	fmt.Println("int8 min: -> ", min)
	fmt.Println("int8 max: -> ", max)
	fmt.Println(min == max+1)
}

func uint8Test() {
	min, max := uint8(0), uint8(255)
	fmt.Println("uint8 min: -> ", min)
	fmt.Println("uint8 max: -> ", max)
	fmt.Println(min == max+1)
}

func int16Test() {
	min, max := int16(-32768), int16(32767)
	fmt.Println("int16 min: -> ", min)
	fmt.Println("int16 max: -> ", max)
	fmt.Println(min == max+1)
}

func uint16Test() {
	min, max := uint16(0), uint16(65535)
	fmt.Println("uint16 min: -> ", min)
	fmt.Println("uint16 max: -> ", max)
	fmt.Println(min == max+1)
}

func int32Test() {
	min, max := int32(-2147483648), int32(2147483647)
	fmt.Println("int32 min: -> ", min)
	fmt.Println("int32 max: -> ", max)
	fmt.Println(min == max+1)
}

func uint32Test() {
	min, max := uint32(0), uint32(4294967295)
	fmt.Println("uint32 min: -> ", min)
	fmt.Println("uint32 max: -> ", max)
	fmt.Println(min == max+1)
}

func int64Test() {
	min, max := int64(-9223372036854775808), int64(9223372036854775807)
	fmt.Println("int64 min: -> ", min)
	fmt.Println("int64 max: -> ", max)
	fmt.Println(min == max+1)
}

func uint64Test() {
	min, max := uint64(0), uint64(18446744073709551615)
	fmt.Println("uint64 min: -> ", min)
	fmt.Println("uint64 max: -> ", max)
	fmt.Println(min == max+1)
}
