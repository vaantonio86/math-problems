package main

import "fmt"

func main() {
    num := 5
    sum := 0
    for i := 1; i <= num; i++ {
        if i % 2 == 0 {
            sum += i
        }
    }
    fmt.Println("The sum of all even numbers up to", num, "is:", sum)
}
