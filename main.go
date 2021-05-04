package main

import (
	"fmt"
	"io/ioutil"

	"github.com/kmwamasali/hellogopher/stringutils"
)

const state = "Maryland"

var name string

var filepath = "proverbs.txt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	name = stringutils.Upper("Kevin")
	from := `Kampala`
	var n int = 5
	dat, err := ioutil.ReadFile(filepath)
	check(err)
	fmt.Print(string(dat))

	fmt.Printf("Hello, my fellow %s Gophers!\n", state)
	fmt.Printf("My name is %s and I'm from %s.\n", name, from)
	fmt.Printf("By the time %d o'clock comes around we shall know how to code in Go\n", n)
	fmt.Println("Let's get started!")
}
