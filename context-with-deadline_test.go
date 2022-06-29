package context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContextWithDeadline(t *testing.T) {
	fmt.Println("Total Goroutines", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	defer cancel() //mengirim sinyal cancel ke context

	destination := CreateCounter3(ctx)
	for n := range destination {
		fmt.Println("Counter", n)
	}

	fmt.Println("Total Goroutines", runtime.NumGoroutine())

}
