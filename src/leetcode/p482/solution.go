package main

import "fmt"

func main() {
	fmt.Println(licenseKeyFormatting("2-4A0r7-4k", 4))
	fmt.Println(licenseKeyFormatting("2-4A0r7-4k", 3))
}

func licenseKeyFormatting(S string, K int) string {
	n := len(S)
	var res string
	g := 0
	for i := n - 1; i >= 0; i-- {
		if S[i] == '-' {
			continue
		}

		res = string(convert(S[i])) + res
		g++

		if g == K {
			res = "-" + res
			g = 0
		}
	}

	if g == 0 && len(res) > 0 {
		return res[1:]
	}
	return res
}

func convert(a byte) byte {
	if a >= 'a' && a <= 'z' {
		return a - 'a' + 'A'
	}

	return a
}
