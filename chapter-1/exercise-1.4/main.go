package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	fileNames := make(map[string]map[string]struct{})
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, fileNames)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, fileNames)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			for fileName := range fileNames[line] {
				fmt.Println(fileName)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int, fileNames map[string]map[string]struct{}) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		if fileNames[line] == nil {
			fileNames[line] = make(map[string]struct{})
		}
		fileNames[line][f.Name()] = struct{}{}
	}
}
