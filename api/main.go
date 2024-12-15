package main

import (
    "fmt"
)

func main() {
    var thing1 = [5]float64{1,2,3,4,5}
    fmt.Printf("\nThe memory location of the thing1 array is: %p", &thing1)
    
    var result [5]float64 = square(&thing1)
    fmt.Printf("\nthe result is: %v", result)
    fmt.Printf("\nThe value of thing1 is: %v\n", thing1)

}

func square(thing2 *[5]float64) [5]float64 {

    fmt.Printf("\nThe memory location of the thing2 array is: %p", thing2)
    for i := range thing2 {
        thing2[i] = thing2[i]*thing2[i]
    }

    return *thing2
}