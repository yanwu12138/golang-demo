/**
@author <a herf="mailto:yanwu0527@163.com">XuBaofeng</a>
@date 2019/11/8 13:44.
<p>
description: 字符串和数值类型的相互转换
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("int to str -> 32:", intToString(32))
	result, err := stringToInt("32")
	fmt.Println("str to inr -> err:", err)
	fmt.Println("str to inr -> 32:", result)
}

func intToString(source int) string {
	return strconv.Itoa(source)
}

func stringToInt(source string) (int, error) {
	return strconv.Atoi(source)
}
