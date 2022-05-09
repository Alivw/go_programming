package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	//test1_7()
	for _, url := range os.Args[1:] {
		prefix := "http://"
		if !bytes.HasPrefix([]byte(url), []byte(prefix)) {
			url = fmt.Sprintf("%s%s", prefix, url)
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:reading %s: %v \n", url, err)
			os.Exit(1)
		}
	}
}

func test1_7() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:reading %s: %v \n", url, err)
			os.Exit(1)
		}
	}
}
