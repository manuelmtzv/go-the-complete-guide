package main

import "fmt"

func main() {
	memory := map[int]int{
		1: 1,
		2: 1,
	}

	fmt.Println(fibonacci(50, memory))
}

func fibonacci(number int, memory map[int]int) int {
	if val, exists := memory[number]; exists {
		fmt.Println(memory)
		return val
	}

	memory[number] = fibonacci(number-1, memory) + fibonacci(number-2, memory)
	return memory[number]
}
