package main

import (
	"fmt"

	"github.com/kmwamasali/hellogopher/stringutils"
)

const state = "Maryland"

var name string

func main() {

	name = stringutils.Upper("Kevin")
	from := `Kampala`
	var n int = 5

	fmt.Printf("Hello, my fellow %s Gophers!\n", state)
	fmt.Printf("My name is %s and I'm from %s.\n", name, from)
	fmt.Printf("By the time %d o'clock comes around we shall know how to code in Go\n", n)
	fmt.Println("Let's get started!")
}
