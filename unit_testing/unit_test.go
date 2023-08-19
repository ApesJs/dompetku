package unit_testing

import (
	"fmt"
	"testing"
)

func DisplayNumber(n int, chanN chan int) {
	chanN <- n
}
func TestGorutin(t *testing.T) {
	chanN := make(chan int)
	for i := 1; i <= 100000; i++ {
		go DisplayNumber(i, chanN)
		fmt.Println(<-chanN)
	}
}

func DisplayNumber2(n int) {
	fmt.Println(n)
}
func TestGorutin2(t *testing.T) {
	for i := 1; i <= 100000; i++ {
		go DisplayNumber2(i)
	}
}
