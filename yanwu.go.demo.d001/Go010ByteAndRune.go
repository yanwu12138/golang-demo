/**
@author <a herf="mailto:yanwu0527@163.com">XuBaofeng</a>
@date 2019/11/7 19:08.
<p>
description: 字符，GO有两种字符：
	1. uint8类型，或者叫byte型，代表了一个ASCII码字符
	2. rune类型，代表一个UTF-8字符，当需要处理中文或其他复合字符时，则需要用到rune类型，rune类型等价于int32类型
*/
package main

import "fmt"

func main() {
	byteTest1()
	byteTest2()
}

func byteTest1() {
	var (
		a byte = 'A'
		b byte = 65
		c byte = '\x41'
	)
	fmt.Println(a == b && b == c)
}

func byteTest2() {
	var (
		ch1 int = '\u0041'
		ch2 int = '\u03B2'
		ch3 int = '\U00101234'
	)
	fmt.Printf("%d - %d - %d\n", ch1, ch2, ch3) // integer
	fmt.Printf("%c - %c - %c\n", ch1, ch2, ch3) // character
	fmt.Printf("%X - %X - %X\n", ch1, ch2, ch3) // UTF-8 bytes
	fmt.Printf("%U - %U - %U\n", ch1, ch2, ch3) // UTF-8 code point
}
