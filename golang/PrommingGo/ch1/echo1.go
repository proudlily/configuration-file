package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i] //s=s+sep+os.Args[i]
		sep = ""
		fmt.Println("sep", sep)
		fmt.Println("os.Arg[i]", os.Args[i])
	}
	fmt.Println(s)

}
