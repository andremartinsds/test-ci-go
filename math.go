package main

import "fmt"

func main() {
	fmt.Println(soma(112, 10))
	fmt.Println(subtracao(100, 50))
	fmt.Println(multiplicacao(10, 10))
}

func soma(a int, b int) int {
	return a + b
}

func subtracao(a int, b int) int {
	return a - b
}

func multiplicacao(a int, b int) int {
	return a * b
}
