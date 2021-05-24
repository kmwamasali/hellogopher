package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kmwamasali/hellogopher/stringutils"
)

const state = "Maryland"

var name string
var filepath string

func check(e error) {
	if e != nil {
		fmt.Printf("%s\n", e)
		os.Exit(2)
	}
}

func main() {

	name = stringutils.Upper("Kevin")
	from := `Kampala`
	var n int = 5

	fileFlag := flag.String("f", "", "command line file flag")
	flag.Parse()

	envFilePath := os.Getenv("FILE")

	if *fileFlag != "" && len(os.Args) > 0 {
		filepath = *fileFlag
	} else if envFilePath != "" {
		filepath = envFilePath
	} else {
		check(fmt.Errorf("no filepath assigned"))
	}

	fmt.Printf("%s", filepath)
	dat, err := ioutil.ReadFile(filepath)
	check(err)

	fmt.Printf("Hello, my fellow %s Gophers!\n", state)
	fmt.Printf("My name is %s and I'm from %s.\n", name, from)
	fmt.Printf("By the time %d o'clock comes around we shall know how to code in Go\n", n)
	fmt.Println("Let's get started!")
	fmt.Print(string(dat))
}
