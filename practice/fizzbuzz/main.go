package pratice

import "fmt"

func main() {
	var n int

	if n%15 == 0 {
		fmt.Println("FizzBuzz")
	} else if n%3 == 0 {
		fmt.Println("Fizz")
	} else if n%5 == 0 {
		fmt.Println("Buzz")
	}
}
