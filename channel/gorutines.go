package main

import (
	"fmt"
	"time"
)

func task(from string) {
	for i := 0; i < 100; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	go task("1st task")
	go task("2nd task")
	time.Sleep(10 * time.Second)
}
