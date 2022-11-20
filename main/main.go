package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var name = flag.String("name", "liang1", "usage")
	var name2 string
	flag.StringVar(&name2, "name2", "yaoming", "usage of name")
	flag.Parse()
	fmt.Println("name: ", *name)
	fmt.Println("name2: ", name2)
	for index, arg := range os.Args {
		fmt.Printf("%d -> %s\n", index, arg)
	}
}
