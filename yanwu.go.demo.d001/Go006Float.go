/**
@author <a herf="mailto:yanwu0527@163.com">XuBaofeng</a>
@date 2019/11/7 17:06.
<p>
description: 浮点数
	浮点数分为两种精度：float32、float64
	通常应该优先使用float64类型，因为float32类型的累计计算误差很容易扩散，并且float32等精确表示的正整数并不是很大
	浮点数在声明的时候可以只写整数部分或小数部分
	很小或者很大的最好使用科学计数法书写（e\E）
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	var f float32 = 16777216
	fmt.Println(f == f+1)

	const e = .71828
	f = 1.

	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.2f\n", math.Pi)
}
