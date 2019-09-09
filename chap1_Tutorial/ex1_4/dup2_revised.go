package main

import (
	"bufio"
	"fmt"
	"os"
)

type Record struct {
	num int
	filenames map[string]bool
}

func countLines(file *os.File, counts map[string]Record, filename string) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		if _, ok := counts[input.Text()]; !ok {
			counts[input.Text()] = Record{0, make(map[string]bool)}
		}
		record := counts[input.Text()]
		record.num++
		record.filenames[filename] = true
		counts[input.Text()] = record
	}
}

func main() {
	counts := make(map[string]Record)
	filenames := os.Args[1:]
	if len(filenames) == 0 {
		countLines(os.Stdin, counts, "keyboard input")
	} else {
		for _, filename := range filenames {
			file, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(file, counts, filename)
			file.Close()
		}
	}
	for line, record := range counts {
		if record.num > 1 {
			fmt.Printf("%s\t occurs inside file: ", line)
			for filename, _ := range record.filenames {
				fmt.Printf("\t%s", filename)
			}
			fmt.Println()
		}
	}
}
