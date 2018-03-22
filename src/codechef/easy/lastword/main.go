package main

import "fmt"

func play(s string) string {
	slice := make([]rune, 0, len(s))

	for _, x := range s {
		if len(slice) == 0 {
			slice = append(slice, x)
			continue
		}

		if x >= slice[0] {
			slice = append([]rune{x}, slice...)
			continue
		}

		slice = append(slice, x)
	}

	return string(slice)
}

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for i := 1; i <= t; i++ {
		var s string
		fmt.Scanf("%s", &s)
		t := play(s)
		fmt.Printf("Case #%d: %s\n", i, t)
	}
}
