package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	// Take actual time
	startTime := time.Now()

	// Create channel
	ch1 := make(chan string)

	for _, url := range os.Args[1:] {
		url := addHttpPrefix(url)

		// Start goroutine
		go fetch(url, ch1)
	}

	for range os.Args[1:] {
		// Process data from channel 'ch1'
		WriteInFile(<-ch1)
	}

	data := fmt.Sprintf("%.2fs elapsed", time.Since(startTime).Seconds())
	fmt.Println(data)
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

func fetch(url string, ch chan<- string) {
	// Take actual time
	startTime := time.Now()

	// Fetch data from url
	resp, err := http.Get(url)

	// If error, return the error data to the channel and return (exit 'fetch' function)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	// Give number of bytes transferred
	nbytes, err := io.Copy(io.Discard, resp.Body)

	// Close body
	resp.Body.Close()

	// If error, return the error data to the channel and return (exit 'fetch' function)
	if err != nil {
		ch <- fmt.Sprintf("error while reading %s: %v", url, err)
		return
	}

	// Calculate response time
	latency := time.Since(startTime).Seconds()

	// Return data to the channel
	ch <- fmt.Sprintf("%2fs %7d %s", latency, nbytes, url)
}

func WriteInFile(newline string) {
	fmt.Println(newline)

	file := "./data.log"
	lines, _ := FileTolines(file)

	fileContent := ""
	for _, line := range lines[0:] {
		fileContent += line + "\n"
	}
	fileContent += newline + "\n"

	os.WriteFile(file, []byte(fileContent), 0644)
}

func FileTolines(filePath string) ([]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return LinesFromReader(f)
}

func LinesFromReader(r io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
