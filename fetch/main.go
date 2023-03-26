package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// USING ioutil.Readall() METHOD

// func main() {
// 	for _, url := range os.Args[1:] {
// 		resp, err := http.Get(url)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "fetch: %v ", err)
// 			os.Exit(1)
// 		}

// 		b, err := ioutil.ReadAll(resp.Body)

// 		resp.Body.Close()

// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v ", url, err)
// 			os.Exit(1)
// 		}

// 		fmt.Println(string(b))
// 	}
// }

// USING io.Copy() METHOD

func main() {
	for _, url := range os.Args[1:] {
		url = addHttpPrefix(url)

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v ", err)
			os.Exit(1)
		}

		dst := os.Stdout
		_, respError := io.Copy(dst, resp.Body)

		resp.Body.Close()

		if respError != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v ", url, err)
			os.Exit(1)
		}

		fmt.Println("STATUS CODE:", resp.StatusCode)

	}
}

func addHttpPrefix(url string) string {
	if strings.HasPrefix(url, "http://") {
		return url
	}
	if strings.HasPrefix(url, "https://") {
		return url
	}
	return "https://" + url
}
