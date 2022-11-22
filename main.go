package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Printf("Current OS: %s\n", getOS())

	channel := make(chan bool)

	// Wait for goruine to end
	go simpleWait(channel)
	fmt.Println("Launched gorutine, waiting...")
	<-channel
	fmt.Println("Finished gorutine")
}

func getOS() string {
	return runtime.GOOS
}

func simpleWait(finish chan bool) {
	time.Sleep(5 * time.Second)

	finish <- true
}
