package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Print("Running")
		time.Sleep(5 * time.Second)
	}
}
