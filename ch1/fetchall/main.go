// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 17.
//!+

// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"strings"
	"io"
	"io/ioutil"
	"bufio"
	"net/http"
	"os"
	"time"
)

func main() {
	execSplit := strings.Split(os.Args[0], string(os.PathSeparator))
	execName := execSplit[len(execSplit) - 1]
	output := "output_" + execName + "_" + time.Now().Format("060102150405")
	fo, err := os.OpenFile(output, os.O_RDWR|os.O_CREATE, 0750)
	if err != nil {
		panic(err)
	}
	defer fo.Close()
	w := bufio.NewWriter(fo)

	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Fprintln(w, <-ch) // receive from channel ch
	}
	fmt.Fprintf(w, "%.2fs elapsed\n", time.Since(start).Seconds())

	w.Flush()
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

//!-
