package main

import "fmt"

func main() {
	age := 21

	incrementAge(&age, 5)

	fmt.Println(age)
}

func incrementAge(age *int, years int) {
	*age += years
}
