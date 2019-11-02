// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 6.
//!+

// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(os.Args[0])
	s, sep := "", ""
	for i, arg := range os.Args[1:] {
		s += sep + strconv.FormatInt(int64(i + 1), 16) + ": " + arg
		sep = "\n"
	}
	fmt.Println(s)
}

//!-
