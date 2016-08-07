package main

import (
	"fmt"
	"strings"
	"time"
)

func add1(r rune) rune { return r + 1 }

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	fmt.Println(strings.Map(add1, "HAL-9000"))
	fmt.Println(strings.Map(add1, "VMS"))
	fmt.Println(strings.Map(add1, "Admix"))

	const timeout = 1 * time.Minute
	fmt.Println(timeout)
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
