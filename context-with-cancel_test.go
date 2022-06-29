package context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

//golang goroutine league
func TestContextWithCancel(t *testing.T) {
	fmt.Println("Total Goroutines", runtime.NumGoroutine())

	destination := CreateCounter()
	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}

	fmt.Println("Total Goroutines", runtime.NumGoroutine())
}

func TestContextWithCancel2(t *testing.T) {
	fmt.Println("Total Goroutines", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounter2(ctx)
	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}
	cancel() //mengirim sinyal cancel ke context

	time.Sleep(2 * time.Second)

	fmt.Println("Total Goroutines", runtime.NumGoroutine())
}
