package main

import "fmt"

func main() {
	//arrayOfIntegers()
	//iceCreamFlavors()
	//creamPrices()
	//iceCreamAvailability()
	//iceCreamToppings()
	//modifySlice2()
	task01()
}

/* /*Questions 1: Array of Integers

Part 1: Write a Go program to declare an array of 5 integers and print the first element of the array.

Part 2: Modify the program from Question 1 to print the fourth element of the array.

Part 3: Extend your program to iterate over the entire array and print each element on a new line.*/

func arrayOfIntegers() {
	numbers := [5]int{1, 2, 3, 4, 5}

	// Accessing and printing elements of the array
		fmt.Println(numbers[3]) // Output: 4
		fmt.Println("")
 
	// Iterating over the array using a range-based loop
	for _, value := range numbers {
	 fmt.Println(value)
	}
	fmt.Println("")
}


/*Questions 2: Array of Ice Cream Flavors (Strings)

Part 1: Create a Go program that declares an array with three ice cream flavors ("Chocolate", "Vanilla", "Strawberry") and prints the first flavor in the array.

Part 2: Modify your program to print the last flavor in the array.

Part 3: Update the program to loop over the array and print each ice cream flavor.
*/

func iceCreamFlavors() {
	flavors := [3]string{"Chocolate", "Vanilla", "Strawberry"}

 fmt.Println(flavors[2])
 fmt.Println("")

 for _, name := range flavors {

  fmt.Println(name)

 }
}

/*Questions 3: Array of Ice Cream Prices (Floating-Point Numbers)

Part 1: Write a Go program that initializes an array with the prices of four ice cream flavors (e.g., 2.99, 3.50, 4.25, 3.75) and prints the price of the second flavor.

Part 2: Modify your program to print the price of the last flavor in the array.

Part 2: Alter the program to iterate over the array using a range-based loop and print each price.*/

func creamPrices() {
	creamPrices := [4]float64{2.99, 3.50, 4.25, 3.75}

	// Accessing and printing elements of the array
		fmt.Println(creamPrices[3]) 
		fmt.Println("")
 
	// Iterating over the array using a range-based loop
	for _, value := range creamPrices {
	 fmt.Println(value)
	}
	fmt.Println("")
}

/*
Questions 4: Array of Ice Cream Availability (Booleans)

Part 1: Create a Go program that declares an array of booleans representing the availability of six ice cream flavors (true or false) and prints the availability of the first flavor.

Part 2: Update your program to print the availability of the fifth flavor.

Part 3: Extend the program to loop over the array and print the availability of each flavor.
*/

func iceCreamAvailability() {
	available := [6]bool{true, false, true, false, false, true}

 fmt.Println(available[0])
 fmt.Println(available[4])
 fmt.Println("")

 for _, x := range available {

  fmt.Println(x)

	}
}

/*Questions 5: Array of Ice Cream Topping Choices (Characters)
Part 1: Write a Go program that initializes an array with various ice cream topping choices (represented by characters like 'C' for chocolate chips, 'S' for sprinkles) and prints the first topping choice.

Part 2: Modify your program to print the last topping choice in the array.

Part 3: Update the program to loop over the array and print each topping choice, separated by a space*/

func iceCreamToppings() {
	toppings := [3]string{"C","D","E"}

  fmt.Println("First topping choice:", toppings[0])
	fmt.Println("")


 for _, name := range toppings {
  fmt.Println(name)
 }

}



func modifySlice() {
 // Creating a slice of integers
 numbers := []int{1, 2, 3, 4, 5}

 // Accessing and printing elements of the slice
 fmt.Println(numbers[0]) // Output: 1
 fmt.Println(numbers[3]) // Output: 4

 // Modifying an element of the slice
 numbers[1] = 10
 fmt.Println(numbers) // Output: [1 10 3 4 5]

 // Creating a sub-slice
 slice := numbers[1:4]
 fmt.Println(slice) // Output: [10 3 4]
}

func modifySlice2() {

	//Creating a slice of strings
 
	names := []string{"Alice", "Swathi", "Dayana"}
 
	//Accessing and printing elements of the slice
 
	fmt.Println(names[1])
	fmt.Println(names[2])
 
	//Modifying an element of the slice
 
	names[1] = "Charlie"
 
	fmt.Println(names)
 
	//Creating a sub slice
 
	slice := names[1:]
	fmt.Println(slice)
 
 }

 func scliceOfFloating() {

	//Creating a slice of floating point numbers
 
	grades := []float64{92.5, 89.0, 72.9, 97.5}
 
	//Accessing and print elements of the slice
 
	fmt.Println(grades[1])
	fmt.Println(grades[3])
 
	//Modifying an element of the slice
 
	grades[2] = 88.2
 
	fmt.Println(grades)
 
	//Creating a sub-slice
	slice := grades[:3]
	fmt.Println(slice)
 
 }

 func sliceofIntegers() {

	//Creating a slice of integers
 
	numbers := []int{1, 2, 3, 4, 5}
 
	//Slicing the slice to create new slices
 
	slice1 := numbers[1:3]
	slice2 := numbers[:3]
	slice3 := numbers[2:]
 
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
 
 }

 func task01() {
	names := []string{"Alice", "Bob", "Eve"}
	fmt.Println("Initial slice: ", names)
 
	//Accessing and printing elements of the slice
 
	fmt.Println(names[1])
	fmt.Println(names[2])
 
	//Modifying an element of the slice
 
	names[2] = "Charlie"
 
	fmt.Println("modified slice: ", names)
 
	//Creating a sub slice
 
	slice := names[2:]
	fmt.Println(slice)
 }

 func CREATING() {
	// Creating a map with string keys and integer values
	ages := map[string]int{
	 "Alice": 30,
	 "Bob":   25,
	 "Eve":   32,
	}
 
	// Accessing and printing values in the map
	fmt.Println(ages["Alice"]) // Output: 30
	fmt.Println(ages["Eve"])   // Output: 32
 
	// Modifying a value in the map
	ages["Bob"] = 26
	fmt.Println(ages["Bob"]) // Output: 26
 
	// Adding a new entry to the map
	ages["Carl"] = 28
	fmt.Println(ages["Carl"]) // Output: 28
 
	// Deleting an entry from the map
	delete(ages, "Eve")
	fmt.Println(ages) // Output: map[Alice:30 Bob:26 Carl:28]
 }

 func mapWithString() {
	// Creating a map with string keys and integer values
	ages := map[string]int{
	 "Alice": 30,
	 "Bob":   25,
	 "Eve":   32,
	}
 
	// Accessing and printing values in the map
	fmt.Println(ages["Alice"]) // Output: 30
	fmt.Println(ages["Eve"])   // Output: 32
 
	// Modifying a value in the map
	ages["Bob"] = 26
	fmt.Println(ages["Bob"]) // Output: 26
 
	// Adding a new entry to the map
	ages["Carl"] = 28
	fmt.Println(ages["Carl"]) // Output: 28
 
	// Deleting an entry from the map
	delete(ages, "Eve")
	fmt.Println(ages) // Output: map[Alice:30 Bob:26 Carl:28]
 }

 func mapIthIntegerKeys() {

	//Creating a map with integer keys and slice values
 
	grades := map[int][]float64{
 
	 1: {90.5, 85.2, 78.3},
	 2: {78.2, 87.5, 92.1, 81.7},
	}
 
	//Accessing and printing values in the map
 
	fmt.Println(grades[1])
	fmt.Println(grades[2])
 
	//Modify a value in the map
	grades[1][2] = 88.0
	fmt.Println(grades[1])
 
 }

 func mapWithMap() {

	//Creating a map with string keys and map values
 
	scores := map[string]map[string]int{
 
	 "Alice": {"Math": 95, "English:": 88, "Science": 92},
	 "Bob":   {"Math": 89, "English:": 78, "Science": 85},
	}
 
	//Accessing and printing values in a map
	fmt.Println(scores["Alice"])
	fmt.Println(scores["Bob"])
 
	//Modifying a value in a map
 
	scores["Alice"]["Math"] = 97
	fmt.Println(scores["Alice"])
 
 }

 func greet() {
	fmt.Println("Hello everyone")
 }
 

//greetByName greets a person by their name
func greetByName(name string) {
 fmt.Printf("Hello, %s!\n", name)
}

func getGreeting() string {

	return "Hello everyone!"
 }
 
/*func main() {
	message := getGreeting()
	fmt.Println(message)
 }*/


//add  two integers and returns the sum

func add(a int, b int) int {
	return a + b
 }
 
 /*func main() {
 
	sum := add(5, 7)
 
	fmt.Println(sum)
 
 }*/

 func checkIceCreamAvailability(flavor string) string {

	switch flavor {
	case "vanilla", "chocolate":
	 return fmt.Sprintf("%s ice cream is available!", flavor)
	default:
	 return fmt.Sprintf("Sorry %s ice cream is not available", flavor)
	}
 
 }
 
 /*func main() {
 
	fmt.Println(checkIceCreamAvailability("vanilla"))
 
	fmt.Println(checkIceCreamAvailability("strawberry"))
 }*/