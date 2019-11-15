/**
@author <a herf="mailto:yanwu0527@163.com">XuBaofeng</a>
@date 2019/11/12 15:56.
<p>
description: 示例：聊天机器人
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入你的名称：")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Printf("系统异常：%s \n", err)
	} else {
		name := input[:len(input)-2]
		fmt.Printf("你好, %s! 有什么可以帮你?\n", name)
	}
	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("系统异常：%s \n", err)
			continue
		}
		input = input[:len(input)-2]
		input = strings.ToLower(input)
		switch input {
		case "":
			continue
		case "nothing", "bye":
			fmt.Println("再见")
			os.Exit(0)
		default:
			fmt.Println("对不起，我不知道该怎么做")
		}
	}
}
