/**
@author <a herf="mailto:yanwu0527@163.com">XuBaofeng</a>
@date 2019/11/7 17:32.
<p>
description: 字符串
	一个字符串是一个不可改变的字节序列，字符串可以包含任意的数据，但通常用来包含刻度的文本
	字符串的值不可变，即字符串是字节的定长数组
	常用的转义字符：（\n：换行符、\r：回车符、\t：制表符、\u：Unicode字符、\\：反斜杠）
	一般的比较运算符（==、!=、<、<=、>、>=）是通过在内存中按字节比较来实现字符串比较的
	字符串的内容可以通过标准索引法来获取，在方括号[]内写入索引，索引从0开始计数
	如果要使用多行字符串时，就必须使用`符号，`符号间中所有的转义字符均无效，文本会将原样输出
*/
package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	str := "Java\nGo"
	fmt.Println(str)
	stringTest1()
	stringTest2()
	stringTest3()
	stringTest4()
	stringTest5()
	stringTest6()
	stringTest7()
}

func stringTest1() {
	tip1, tip2 := "genji is a ninja", "忍者"
	fmt.Println(len(tip1))
	fmt.Println(len(tip2))
	fmt.Println(utf8.RuneCountInString(tip1))
	fmt.Println(utf8.RuneCountInString(tip2))
}

/**
字符串遍历
*/
func stringTest2() {
	theme := "狙击 start"
	fmt.Println("---------- 字符串遍历：ascii ----------")
	for i := 0; i < len(theme); i++ {
		fmt.Printf("ascii: %c %d\n", theme[i], theme[i])
	}
	fmt.Println("---------- 字符串遍历：Unicode ----------")
	for _, s := range theme {
		fmt.Printf("Unicode: %c %d\n", s, s)
	}
}

/**
字符串截取
*/
func stringTest3() {
	tracer := "死神来了,死神bye"
	comma := strings.Index(tracer, ",")
	pos := strings.Index(tracer[comma:], "死神")
	fmt.Println(comma, pos, tracer[comma:])
	split := strings.Split(tracer, ",")
	fmt.Println(split)
}

/**
字符串修改
*/
func stringTest4() {
	angel := "hello go"
	angelBytes := []byte(angel)
	for i := 1; i <= 2; i++ {
		angelBytes[i] = ' '
	}
	fmt.Println(string(angelBytes))
}

/**
字符串拼接
*/
func stringTest5() {
	hammer, sickle := "hello ", "java"
	var sb bytes.Buffer
	sb.WriteString(hammer)
	sb.WriteString(sickle)
	fmt.Println(sb.String())
}

/**
字符串格式化输出
动词		功能
%v		按值的本来值输出
%+v		在 %v 基础上，对结构体字段名和值进行展开
%#v		输出 Go 语言语法格式的值
%T		输出 Go 语言语法格式的类型和值
%%		输出 % 本体
%b		整型以二进制方式显示
%o		整型以八进制方式显示
%d		整型以十进制方式显示
%x		整型以十六进制方式显示
%X		整型以十六进制、字母大写方式显示
%U		Unicode 字符
%f	浮点数
%p	指针，十六进制方式显示
*/
func stringTest6() {
	progress, target := 2, 8
	title := fmt.Sprintf("以采集%d个药草，还需要%d个完成任务", progress, target)
	fmt.Println(title)

	pi := 3.1415926727
	title = fmt.Sprintf("%v %v %v", "月球基地", pi, true)
	fmt.Println(title)

	profile := &struct {
		name string
		hp   int
	}{
		name: "yanwu",
		hp:   100,
	}
	fmt.Printf("使用'%%+v' %+v\n", profile)
	fmt.Printf("使用'%%#v' %#v\n", profile)
	fmt.Printf("使用'%%T' %T\n", profile)
}

/**
base64编码
*/
func stringTest7() {
	message := "Away from keyboard. https://golang.org/"
	// 编码
	encodeToString := base64.StdEncoding.EncodeToString([]byte(message))
	fmt.Println(encodeToString)
	// 解码
	data, err := base64.StdEncoding.DecodeString(encodeToString)
	if err == nil {
		fmt.Println(string(data))
	} else {
		fmt.Println(err)
	}
}
