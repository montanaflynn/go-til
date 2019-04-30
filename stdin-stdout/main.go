package main

import (
	"fmt"
	"os"
)

func main() {
	fh, err := os.Open("/dev/stdin")
	if err != nil {
		fmt.Println("os.Open err:", err.Error())
		return
	}
	defer fh.Close()

	b := make([]byte, 1)
	for {
		i, err := fh.Read(b)
		if err != nil {
			fmt.Println("fh.Read err:", err.Error())
			break
		}
		if i == 1 {
			fmt.Println("b:", b)
			continue
		}
		fmt.Println("i:", i)
	}
}
