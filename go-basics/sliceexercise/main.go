package main

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func main() {

	// Time to practice what you learned!

	// 1) Create a new array (!) that contains three hobbies you have
	// 		Output (print) that array in the command line.
	hobbies := [3]string{"Disc Golf", "Paddle Board", "Ping Pong"}
	fmt.Println(hobbies)
	// 2) Also output more data about that array:
	//		- The first element (standalone)
	fmt.Println(hobbies[0])
	//		- The second and third element combined as a new list
	newList := hobbies[1:]
	fmt.Println(newList)
	// 3) Create a slice based on the first element that contains
	//		the first and second elements.
	//		Create that slice in two different ways (i.e. create two slices in the end)
	var mainHobbies = hobbies[0:2]
	fmt.Println(mainHobbies)
	mainHobbies = hobbies[0:3][0:2]
	fmt.Println(mainHobbies)

	// 4) Re-slice the slice from (3) and change it to contain the second
	//		and last element of the original array.
	mainHobbies = mainHobbies[1:3]

	fmt.Println("4", mainHobbies)

	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	courseGoals := []string{"Learn Go", "Learn all the basics"}

	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array

	courseGoals[1] = "Learn all the details"
	courseGoals = append(courseGoals, "Learn even more")
	fmt.Println("6", courseGoals)
	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	//		dynamic list of products (at least 2 products).
	//		Then add a third product to the existing list of products.
	products := []Product{
		{id: "1", title: "title 1", price: 1.00},
		{id: "2", title: "product 2", price: 2.33},
	}
	fmt.Println(products)
	products = append(products, Product{id: "3", title: "title 3", price: 3.0})
	fmt.Println(products)

}
