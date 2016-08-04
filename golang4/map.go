package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}

	for k, v := range ages {
		fmt.Printf("%s = %d\n", k, v)
	}

	ages["hoge"] = 123
	fmt.Println(ages["hoge"])

	for k, v := range ages {
		fmt.Printf("%s = %d\n", k, v)
	}

	delete(ages, "hoge")

	for k, v := range ages {
		fmt.Printf("%s = %d\n", k, v)
	}

	fmt.Printf("%s = %d\n", "bob", ages["bob"])
	ages["bob"] = ages["bob"] + 1
	fmt.Printf("%s = %d\n", "bob", ages["bob"])
	ages["bob"] += 1
	fmt.Printf("%s = %d\n", "bob", ages["bob"])
	ages["bob"]++
	fmt.Printf("%s = %d\n", "bob", ages["bob"])

	var names []string
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
}
