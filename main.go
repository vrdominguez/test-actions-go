package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Current OS: %s\n", getOS())
}

func getOS() string {
	return runtime.GOOS
}
