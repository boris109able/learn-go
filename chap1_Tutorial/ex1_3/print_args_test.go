package main

import(
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func BenchmarkPrintArgsLoop(b *testing.B) {
	//oldArgs := os.Args
	//defer func() { os.Args = oldArgs }()
	//os.Args = []string
	var args []string
	for i := 0; i < 10000; i++ {
		args = append(args, strconv.Itoa(i))
	}
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	fmt.Println(s)
}

func BenchmarkPringArgsJoin(b *testing.B) {
//	oldArgs := os.Args
//	defer func() { os.Args = oldArgs }()
//	os.Args = []string
        var args []string
	for i := 0; i < 10000; i++ {
		args = append(args, strconv.Itoa(i))
	}
	fmt.Println(strings.Join(args[1:], " "))
}
