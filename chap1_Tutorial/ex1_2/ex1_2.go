package main

import (
	"fmt"
	"os"
)

func main() {
	for idx, str := range os.Args[1:] {
		fmt.Println(idx, str)
	}
}
