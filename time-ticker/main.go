package main

import (
	"fmt"
	"time"
)

func sendResults(results <-chan string) {
	for result := range results {
		fmt.Println(result)
	}
}

func tickTock(results chan<- string) {
	ticker := time.NewTicker(time.Second)

	count := 0
	for {
		select {
		case tick := <-ticker.C:
			count++
			// fmt.Println(tick.String(), count)
			results <- fmt.Sprintln(tick.String(), count)
			time.Sleep(time.Millisecond * 1100)
			if count > 5 {
				ticker.Stop()
				close(results)
				return
			}
		}
	}
}

func main() {
	// create ticker that ticks every second

	results := make(chan string)
	go sendResults(results)
	go tickTock(results)

	for {
		time.Sleep(time.Second)
		fmt.Println(time.Now())
	}
}
