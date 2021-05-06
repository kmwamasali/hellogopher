package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kmwamasali/hellogopher/stringutils"
)

const state = "Maryland"

var name string

func check(e error) {
	if e != nil {
		os.Exit(2)
	}
}

func main() {

	name = stringutils.Upper("Kevin")
	from := `Kampala`
	var n int = 5

	filepath := os.Getenv("FILEPATH")

	if filepath == "" {
		check(fmt.Errorf("filepath not found"))
	}

	dat, err := ioutil.ReadFile(filepath)
	check(err)

	fmt.Printf("Hello, my fellow %s Gophers!\n", state)
	fmt.Printf("My name is %s and I'm from %s.\n", name, from)
	fmt.Printf("By the time %d o'clock comes around we shall know how to code in Go\n", n)
	fmt.Println("Let's get started!")
	fmt.Print(string(dat))
}
