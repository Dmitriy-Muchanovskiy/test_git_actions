package main

import (
    "fmt"
    "github.com/logrusorgru/aurora"
    "go_test/test_1"
)

func main() {
    greetEmpty := test_1.Hello("")
    fmt.Println(aurora.Yellow(greetEmpty))

    greetMessage := test_1.Hello("Tester")
    fmt.Println(aurora.Yellow(greetMessage))

    any_var := test_1.First()
    fmt.Println(aurora.Red(any_var))
}