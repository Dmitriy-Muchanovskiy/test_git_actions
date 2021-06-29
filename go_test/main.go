package main

import (
    "fmt"
    "github.com/logrusorgru/aurora"
)

func main() {
    greetEmpty := hello("")
    fmt.Println(aurora.Yellow(greetEmpty))

    greetMessage := hello("Tester")
    fmt.Println(aurora.Yellow(greetMessage))

    any_var := first()
    fmt.Println(aurora.Red(any_var))
}