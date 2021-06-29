package main
import "testing"

//func TestAbc(t *testing.T) {
//    t.Error() // to indicate test failed
//}

func TestHelloEmpty(t *testing.T) {
    //test for empty argument
    emptyResult := hello("") // should be return Hello Empty

    if emptyResult != "Hello Empty" {
        t.Errorf("hello(\"\") failed, expected %v, got %v", "Hello Empty", emptyResult)
    } else {
        t.Logf("hello(\"\") success, expected %v, got %v", "Hello Empty", emptyResult)
    }
}

func TestHelloValue(t *testing.T) {
    //test for empty argument
    emptyResult := hello("Test") // should be return Hello Empty

    if emptyResult != "Hello Test!" {
        t.Errorf("hello(\"\") failed, expected %v, got %v", "Hello Test!", emptyResult)
    } else {
        t.Logf("hello(\"\") success, expected %v, got %v", "Hello Test!", emptyResult)
    }
}
