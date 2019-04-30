package main

import (
	"context"
	"log"
	"os/exec"
	"time"
)

func killAfterSomeTime(cancel context.CancelFunc) {
	time.Sleep(time.Second * 4)
	cancel()
}

func runSomeCommand() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go killAfterSomeTime(cancel)

	cmd := exec.CommandContext(ctx, "cat", "/dev/urandom")

	err := cmd.Start()
	if err != nil {
		log.Println(err)
	}

	err = cmd.Wait()
	if err != nil {
		log.Println(err)
	}
}

func main() {
	log.Println("running some command")
	runSomeCommand()
	log.Println("some command finished")
}
