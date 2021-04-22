// package main is designed to solve the programming problem below:
// Write a program that prints the numbers from 1 to 100. But for multiples of three
// print “Fizz” instead of the number and for the multiples of five print “Buzz”. For
// numbers which are multiples of both three and five print “FizzBuzz”.

package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(multipleof(i))
	}
}

// here we separate things out into their own function as it's way easier to test outside
// of the main function.
// This way we have segmented the important logic out so we can test it in isolation.
func multipleof(i int) string {
	switch {
	case i%3 == 0 && i%5 == 0:
		return "fizzbuzz"
	case i%3 == 0:
		return "fizz"
	case i%5 == 0:
		return "buzz"
	default:
		return fmt.Sprint(i)
	}
}
