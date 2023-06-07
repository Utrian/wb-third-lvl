package main

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

func main() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

	h, m, s := time.Clock()
	fmt.Printf("Current time from NTP server: %d:%d:%d\n", h, m, s)
}
