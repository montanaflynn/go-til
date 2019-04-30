package main

import (
	"fmt"
	"os/exec"
	"sync"
)

func execCommand(input []string, wg *sync.WaitGroup) {
	defer wg.Done()
	cmd := exec.Command(input[0], input[1:]...)
	fmt.Println("Starting", input)

	err := cmd.Start()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Println("Started", input)

	err = cmd.Wait()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Println("Finished", input)
}

func main() {
	wg := new(sync.WaitGroup)
	commands := [][]string{[]string{"sleep", "1"}, []string{"sleep", "2"}, []string{"sleep", "3"}}
	for _, command := range commands {
		wg.Add(1)
		go execCommand(command, wg)
	}
	wg.Wait()
}
