
package main

import "fmt"

func randomInt(min int, max int) int {
	return min + (max - min) * rand.Float64()
}

func main() {
	fmt.Println("The random integer between 1 and 5 is", randomInt(1, 5))
}