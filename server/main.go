package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Main started")
	var wg sync.WaitGroup
	wg.Add(2)
	// Start WebSocket room server
	go func() {
		defer wg.Done()
		InitRoom()
	}()
	// Start additional HTTP server (REST, health, etc.)
	go func() {
		defer wg.Done()
		StartServer()
	}()
	wg.Wait()
}
