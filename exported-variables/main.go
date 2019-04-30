package main

import (
	"github.com/montanaflynn/til/exported-variables/name"
	"github.com/montanaflynn/til/exported-variables/say"
)

func main() {
	say.Hello(name.Default)

	say.Greeting = "Hola"

	say.Hello(name.Default)
}
