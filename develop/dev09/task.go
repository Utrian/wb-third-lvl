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
		return
	}

	url := os.Args[1]

	if err := urlValid(url); err != nil {
		fmt.Println("error:", err)
		return
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	defer resp.Body.Close()

	base := filepath.Base(url)
	file, err := os.Create(base)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("error:", err)
		return
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
