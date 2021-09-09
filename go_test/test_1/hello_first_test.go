package test_1
import "testing"

//func TestAbc(t *testing.T) {
//    t.Error() // to indicate test failed
//}

// func TestHelloEmpty(t *testing.T) {
//     //test for empty argument
//     emptyResult := Hello("") // should be return Hello Empty

//     if emptyResult != "Hello Empty" {
//         t.Errorf("hello(\"\") failed, expected %v, got %v", "Hello Empty", emptyResult)
//     } else {
//         t.Logf("hello(\"\") success, expected %v, got %v", "Hello Empty", emptyResult)
//     }
// }

func TestHelloValue(t *testing.T) {
    //test for empty argument
    emptyResult := Hello("Test") // should be return Hello Empty

    if emptyResult != "Hello Test!" {
        t.Errorf("hello(\"\") failed, expected %v, got %v", "Hello Test!", emptyResult)
    } else {
        t.Logf("hello(\"\") success, expected %v, got %v", "Hello Test!", emptyResult)
    }
}
/*
func TestFirst(t *testing.T) {
 
    result := First()

    if result == "" {
        t.Errorf("first hello failed, expected Just string, got %v", result)
    } else {
        t.Logf("first hello success, expected Just string, got %v", result)
    }
}*/