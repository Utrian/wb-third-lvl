package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

// Реализовать аналог утилиты wget

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <url>")
		os.Exit(1)
	}

	url := os.Args[1]

	if err := urlValid(url); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	base := filepath.Base(url)
	file, err := os.Create(base)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	fmt.Println("Downloaded:", base)
}

func urlValid(rawURL string) error {
	_, err := url.Parse(rawURL)
	if err != nil {
		return err
	}
	return nil
}
