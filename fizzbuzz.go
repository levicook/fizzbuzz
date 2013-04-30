package main

import "fmt"

func fizzbuzz(i int) (result string) {
	var (
		multipleOf3 = i%3 == 0
		multipleOf5 = i%5 == 0
	)

	switch {

	case multipleOf3 && multipleOf5:
		result = "FizzBuzz"

	case multipleOf3:
		result = "Fizz"

	case multipleOf5:
		result = "Buzz"

	default:
		result = fmt.Sprintf("%v", i)

	}

	return
}
