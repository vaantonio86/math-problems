
package main
import "math/rand"

func main() {
    const num1 = rand.Intn(10)
    const num2 = rand.Intn(10)
    var sum = num1 + num2
    fmt.Println("The solution to the math problem is:", sum)
}