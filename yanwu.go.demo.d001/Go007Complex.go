/**
@author <a herf="mailto:yanwu0527@163.com">XuBaofeng</a>
@date 2019/11/7 17:17.
<p>
description: 复数
	复数分为两种类型：complex64、complex128
	复数使用re + imi来表示，其中re表示实数部分，im表示虚数部分，i代表根号-1
	如果一个浮点数或者十进制整数后面跟着一个i，它将构成一个复数的虚部，而复数的实部是0
	通常情况下使用complex128，因为相关函数都是用这个类新的参数
*/
package main

import "fmt"

func main() {
	var x complex128 = complex(1, 2)
	var y complex128 = complex(3, 4)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(real(x) * real(y))
	fmt.Println(imag(x * y))
}
