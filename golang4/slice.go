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

func appendInt(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
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
	copy(z[len(x):], y)
	return z
}

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func remove2(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
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
		x = appendInt(x, i, i)
		fmt.Printf("%d\tcap=%d\t%v\n", i, cap(x), x)
	}

	var y []int
	y = append(y, 1)
	fmt.Printf("cap=%d\t%v\n", cap(y), y)
	y = append(y, 2, 3)
	fmt.Printf("cap=%d\t%v\n", cap(y), y)
	y = append(y, 4, 5, 6)
	fmt.Printf("cap=%d\t%v\n", cap(y), y)
	y = append(y, y...)
	fmt.Printf("cap=%d\t%v\n", cap(y), y)

	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", data)
	fmt.Printf("%q\n", nonempty(data))
	fmt.Printf("%q\n", data)

	data = []string{"one", "", "three"}
	fmt.Printf("%q\n", data)
	fmt.Printf("%q\n", nonempty2(data))
	fmt.Printf("%q\n", data)

	z := []int{5, 6, 7, 8, 9}
	fmt.Println(z)
	fmt.Println(remove(z, 2))

	z = []int{5, 6, 7, 8, 9}
	fmt.Println(z)
	fmt.Println(remove2(z, 2))
}
