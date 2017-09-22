package main

import "fmt"

func main() {
	var a int = 20
	var ip *int
	ip = &a
	fmt.Printf("变量的地址是%x\n", &a)
	//指针的变量存储地址
	fmt.Printf("Ip变量的存储的指针地址%x\n", ip)

	//使用指针访问之
	fmt.Printf("IP变量的值:%d\n", *ip)

}
