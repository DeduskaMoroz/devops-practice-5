package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"
)

var Version = "dev"

func main() {
	showVersion := flag.Bool("version", false, "print version and exit")
	flag.Parse()

	if *showVersion {
		fmt.Printf("syspulse %s\n", Version)
		return
	}

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	fmt.Printf("Syspulse v%s\n", Version)
	fmt.Printf("Host: %s\n", hostname)
	fmt.Printf("OS/Arch: %s/%s\n", runtime.GOOS, runtime.GOARCH)
	fmt.Printf("CPUs: %d\n", runtime.NumCPU())
	fmt.Printf("Time: %s\n", time.Now().Format(time.RFC3339))
}
