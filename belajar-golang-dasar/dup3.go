package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	
	for _, filename := range os.Args[1:] {
		// Check file size first
		info, err := os.Stat(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		
		// Skip if file too large (> 100MB)
		if info.Size() > 100*1024*1024 {
			fmt.Fprintf(os.Stderr, "dup3: file %s too large\n", filename)
			continue
		}
		
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}