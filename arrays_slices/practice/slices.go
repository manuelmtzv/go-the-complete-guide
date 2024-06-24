package practice

import "fmt"

func SlicesPractice() {
	// 1) Create a new array (!) that contains three hobbies you have
	hobbies := []string{"Playing videogames", "Watch One Piece", "Coding"}

	fmt.Println(hobbies)

	// 2) Also output more data about that array:
	fmt.Println("First element: ", hobbies[0])
	fmt.Println("Second and third element: ", hobbies[1:])

	// 3) Create a slice based on the first element that contains the first and second elements.
	hobbiesReduced := hobbies[:1]
	fmt.Println(hobbiesReduced)

	// 4 Re-slice the slice from (3) and change it to contain the second and last element of the original array.
	hobbiesReduced = hobbies[0:2]
	fmt.Println(hobbiesReduced)

	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	courseGoals := []string{"Learn Go", "Understand Go", "Start Go projects journey"}

	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	courseGoals[1] = "Understand Go better"
	courseGoals = append(courseGoals, "Master go in future")
	fmt.Println(courseGoals)

	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	type Product struct {
		Id    string
		Title string
		Price float64
	}

	products := []Product{
		{Id: "1", Title: "Product 1", Price: 10.0},
		{Id: "2", Title: "Product 2", Price: 20.0},
	}

	products = append(products, Product{Id: "3", Title: "Product 3", Price: 30.0})
	fmt.Println(products)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
