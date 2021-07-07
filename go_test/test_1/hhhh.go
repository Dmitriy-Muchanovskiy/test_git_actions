package test_1
import "fmt"

func Hello(user string) string {
    if len( user ) == 0 {
        return "Hello Empty"
    } else {
        return fmt.Sprintf("Hello %v!", user)
    }
}