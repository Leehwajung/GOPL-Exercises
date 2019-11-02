// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.

// Echo3 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

//!+
func main() {
	args := make([]string, len(os.Args) - 1)
	for i, arg := range os.Args[1:] {
		args[i] = strconv.FormatInt(int64(i + 1), 16) + ": " + arg
	}
	fmt.Println(os.Args[0])
	fmt.Println(strings.Join(args, "\n"))
}

//!-
