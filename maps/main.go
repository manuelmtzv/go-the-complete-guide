package main

import "fmt"

type floatMap map[string]float64

func main() {
	names := make([]string, 2)

	names[0] = "John"
	names[1] = "Doe"

	for index, value := range names {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}

	courseRatings := make(floatMap, 3)

	courseRatings["Go Fundamentals"] = 5
	courseRatings["Docker Deep Dive"] = 4
	courseRatings["Kubernetes"] = 3

	for key, value := range courseRatings {
		fmt.Printf("Course: %s, Rating: %f\n", key, value)
	}
}
