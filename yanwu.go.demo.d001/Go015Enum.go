/**
@author <a herf="mailto:yanwu0527@163.com">XuBaofeng</a>
@date 2019/11/8 11:14.
<p>
description: 枚举
	go现阶段只能通过iota来模型枚举类型

*/
package main

import "fmt"

func main() {
	enumTest1()
	enumTest2()
	enumTest3()
}

type Weapon int

func enumTest1() {
	const (
		arrow Weapon = iota
		shu
		sniper
		rifle
		blower
	)
	fmt.Println(arrow, shu, sniper, rifle, blower)
	weapon := rifle
	fmt.Println(weapon)
}

func enumTest2() {
	const (
		none = 1 << iota // 左移一位
		red
		green
		blue
	)
	fmt.Printf("%d %d %d %d\n", none, red, green, blue)
	fmt.Printf("%b %b %b %b\n", none, red, green, blue)
}

type Chip int

const (
	none Chip = iota
	cpu
	gpu
)

func enumTest3() {
	fmt.Printf("%s %d", cpu, gpu)
}

func (c Chip) String() string {
	switch c {
	case none:
		return "None"
	case cpu:
		return "CPU"
	case gpu:
		return "GPU"
	default:
		return "N/A"
	}
}
