package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	// create context with cancel
	ctx1, _ := context.WithCancel(context.Background())

	// create secondary context
	ctx2, cancel2 := context.WithCancel(ctx1)

	// create third context
	ctx3, _ := context.WithCancel(ctx2)

	go func() {
		// time.Sleep(time.Second)
		cancel2()
		// time.Sleep(time.Second)
		// cancel1()
		// time.Sleep(time.Second)
		// cancel3()
	}()

	// select {
	// case <-ctx1.Done():
	// 	fmt.Println("ctx1 done", ctx1.Err())
	// case <-ctx2.Done():
	// 	fmt.Println("ctx2 done", ctx2.Err())
	// case <-ctx3.Done():
	// 	fmt.Println("ctx3 done", ctx3.Err())
	// }

	time.Sleep(time.Second)
	fmt.Println("ctx1 done", ctx1.Err())
	fmt.Println("ctx2 done", ctx2.Err())
	fmt.Println("ctx3 done", ctx3.Err())

	if ctx1.Err() != nil && ctx2.Err() != nil && ctx3.Err() != nil {
		return
	}

}
