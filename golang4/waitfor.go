package main

import (
	"fmt"
	"time"
)

func main() {
	const timeout = 1 * time.Minute
	fmt.Println(timeout)
	fmt.Println(time.Now())
	deadline := time.Now().Add(timeout)
	fmt.Println(deadline)

	var sleeptime time.Duration
	for tries := 0; time.Now().Before(deadline); tries++ {
		sleeptime = time.Second << uint(tries)
		fmt.Println("now = ", time.Now(), "\tsleep=", sleeptime)
		time.Sleep(sleeptime)
	}
	fmt.Println(time.Now())
}
