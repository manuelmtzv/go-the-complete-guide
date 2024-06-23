package main

func main() {
	println(Add(1, 2))
	println(Add(1.1, 2.2))
	println(Add("Hello, ", "World!"))
}

func Add[T string | int | float64](a, b T) T {
	return a + b
}
