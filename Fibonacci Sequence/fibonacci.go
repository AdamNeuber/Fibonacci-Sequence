package main

import "fmt"

func fib() {
	value1, value2, i := 1, 1, 0

	fmt.Println(value1)
	for i < 10 {
		value1 = value1 + value2
		fmt.Println(value1)
		value2 = value1 + value2
		fmt.Println(value2)
		i++
	}
}

func main() {
	fib()
}
