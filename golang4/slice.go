package main

import "fmt"

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func main() {
	months := [...]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}

	fmt.Println(months[1:])
	fmt.Println(months[1:13])
	fmt.Println(months[:])
	fmt.Println(len(months[:]))
	fmt.Println(cap(months[:]))

	Q2 := months[4:7]
	summer := months[6:9]

	fmt.Println(Q2)
	fmt.Println(len(Q2))
	fmt.Println(cap(Q2))
	fmt.Println(summer)
	fmt.Println(len(summer))
	fmt.Println(cap(summer))

	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}

	//	fmt.Println(summer[:20])
	endlessSummer := summer[:5]
	fmt.Println(endlessSummer)
	fmt.Println(len(endlessSummer))
	fmt.Println(cap(endlessSummer))

	a := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Println(a)
	reverse(a[:])
	fmt.Println(a)

	s := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(s)
	reverse(s[:2])
	fmt.Println(s)
	reverse(s[2:])
	fmt.Println(s)
	reverse(s)
	fmt.Println(s)

	Q22 := months[4:7]
	fmt.Println(equal(Q2, summer))
	fmt.Println(equal(Q2, Q22))

	var runes []rune
	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)

	var x []int
	for i := 0; i < 10; i++ {
		x = appendInt(x, i)
		fmt.Printf("%d\tcap=%d\t%v\n", i, cap(x), x)
	}
}
