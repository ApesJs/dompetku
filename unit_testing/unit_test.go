package unit_testing

import (
	"fmt"
	"sync"
	"testing"
)

func DisplayNumber(n int, chanTes chan int) {
	chanTes <- n
}
func TestGorutin(t *testing.T) {
	// chanTest := make(chan int)
	var mutex sync.Mutex
	var x = 0
	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
		// chanTest <- x
	}
	fmt.Println(x)
}
