package main

import "fmt"

func main() {
	var n [1000]int // 是一个长度为10的数组
	var i, j int

	for i = 0; i < 10; i++ {
		n[i] = i + 100 //设置元素
	}
	for i = 0; j < 100; j++ {
		fmt.Printf("em[%d]=%d\n", j, n[j])
	}
}
