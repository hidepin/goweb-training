package main

import "fmt"

func main() {
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}

	for k, v := range ages {
		fmt.Printf("%s = %d\n", k, v)
	}
}
