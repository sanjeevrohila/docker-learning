package main

import (
	"log"
	"time"
)

func main() {
	for {
		log.Printf("Test log to see on the console")
                time.Sleep(10 * time.Second)
	}
}
