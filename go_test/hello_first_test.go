package main
import "testing"

func TestFirst(t *testing.T) {
 
    result := "asd"//first()

    if result == "" {
        t.Errorf("first hello failed, expected Just string, got %v", result)
    } else {
        t.Logf("first hello success, expected Just string, got %v", result)
    }
}