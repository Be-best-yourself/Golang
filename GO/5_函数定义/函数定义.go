package main

import "fmt"

func math1(num1 int, num2 int) int {
	var result int
	if num1 < num2 {
		result = num2
		if num1 > num2 {
			result = num1
		}
	} else {
		result = num1 + num2
	}

	return result
}
func main() {
	var num1, num2 int
	num1 = 10
	num2 = 20
	var result int
	result = math1(num1, num2)
	fmt.Printf("%d", result)

}
	return result
}
func main() {
	var num1, num2 int
	num1 = 10
	num2 = 20
	var result int
	result = math1(num1, num2)
	fmt.Printf("%d", result)

}
