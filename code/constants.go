package main

import "fmt"
import "math"

const s string = "constant"

func main() {
	fmt.Println(s)
	
	const n = 5_000_000_000
	fmt.Println(n)

	

	fmt.Println(math.Sin(n))

}