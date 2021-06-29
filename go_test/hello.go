package main

import "fmt"

func hello(user string) string {
    if len( user ) == 0 {
        return "Hello Empty"
    } else {
        return fmt.Sprintf("Hello %v!", user)
    }
}