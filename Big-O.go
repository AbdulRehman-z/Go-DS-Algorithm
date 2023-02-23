package main

import "fmt"

func main() {
	// O(n) Linear- for loops, while loops through n items
	/*
		This is a simple example of O(n) time complexity. The for loop will run 100,000 times.
		The time complexity is O(n) because the number of operations is directly proportional to the size of the input.
	*/
	//OOfn()

	// O(1) Constant- no loops
	/*
		This is a simple example of O(1) time complexity. The function will run once.
		The time complexity is O(1) because the number of operations is constant.
	*/
	//OOff1()

	// O(n^2) Quadratic- every element in a collection needs to be compared to ever other element. Two nested loops!
	/*
		This is a simple example of O(n^2) time complexity. The function will run once.
		The time complexity is O(n^2) because the number of operations is proportional to the square of the size of the input(array of numbers).
	*/
	//OOfnSq()

	// Space Complexity - O(1) Constant
	/*
		This is a simple example of O(1) space complexity. The function will run once.
		The space complexity is O(1) because the amount of memory is constant.
	*/
	//OOff1Space()

	// Space Complexity - O(n) Linear
	/*
		This is a simple example of O(n) space complexity. The function will run once.
		The space complexity is O(n) because the amount of memory is proportional to the size of the input(array of numbers).
	*/
	//OOfnSpace()

}

// Time Complexity - O(n) Linear
func OOfn() {
	for i := 0; i < 100000; i++ {
		fmt.Println(i)
	}
}

// Time Complexity - O(1) Constant
func OOff1() {
	// slice of strings
	names := []string{"John", "Paul", "George", "Ringo"}
	fmt.Println(names[0])
}

// Time Complexity - O(n^2) Quadratic
func OOfnSq() {
	// slice of integers
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			fmt.Println(numbers[i], numbers[j])
		}
	}
}

// Space Complexity - O(1) Constant
func OOff1Space() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

// Space Complexity - O(n) Linear
func OOfnSpace() {
	// slice of integers
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// empty slice
	var doubled []int
	for i := 0; i < len(numbers); i++ {
		doubled = append(doubled, numbers[i])
	}
}
