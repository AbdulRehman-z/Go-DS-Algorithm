package main

import (
	"fmt"
)

func main() {
	// O(n) Linear- for loops, while loops through n items
	/*
		This is a simple example of O(n) time complexity. The for loop will run 100,000 times.
		The time complexity is O(n) because the number of operations is directly proportional to the size of the input.
	*/
	for i := 0; i < 100000; i++ {
		fmt.Println(i)
	}

}
