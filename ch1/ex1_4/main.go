// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
// gopl.io/ch1/dup2 を改変したものです。

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	files := make(map[string]string)
	readFiles := os.Args[1:]
	if len(readFiles) == 0 {
		countLines(os.Stdin, counts, files)
	} else {
		for _, arg := range readFiles {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, files)
			f.Close()
		}
	}
	for line, fname := range files {
		fmt.Printf("keyword %s  in  %s\n", line, fname)
	}
}

func countLines(f *os.File, counts map[string]int, files map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if counts[input.Text()] == 1 {
			files[input.Text()] = f.Name()
		} else {
			if strings.Contains(files[input.Text()], f.Name()) {
				files[input.Text()] += " " + f.Name()
			}
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
