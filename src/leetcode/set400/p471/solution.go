package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(encode("aaa"))
	fmt.Println(encode("aaaaa"))
	fmt.Println(encode("aaaaaaaaaa"))
	fmt.Println(encode("aabcaabcd"))
	fmt.Println(encode("abbbabbbcabbbabbbc"))
}

func encode(s string) string {
	mem := make([]map[int]string, len(s))

	for i := 0; i < len(s); i++ {
		mem[i] = make(map[int]string)
	}

	var doEncode func(i, j int) string

	doEncode = func(i, j int) string {
		if len(mem[i][j]) > 0 {
			return mem[i][j]
		}
		mem[i][j] = s[i : i+j]
		for k := 1; k < j; k++ {
			if j%k != 0 {
				continue
			}

			ok := true
			for l := 1; l*k < j; l++ {
				if s[i:i+k] != s[i+l*k:i+l*k+k] {
					ok = false
					break
				}
			}

			if !ok {
				continue
			}

			ik := doEncode(i, k)

			for l := 1; l*k < j; l++ {
				mem[i+l*k][k] = ik
			}

			if len(mem[i][j]) > len(Int2String(j/k))+2+len(ik) {
				mem[i][j] = Int2String(j/k) + "[" + ik + "]"
			}
		}

		for k := 1; k < j; k++ {
			a := doEncode(i, k)
			b := doEncode(i+k, j-k)

			if len(mem[i][j]) > len(a)+len(b) {
				mem[i][j] = a + b
			}
		}
		return mem[i][j]
	}

	return doEncode(0, len(s))
}

func Int2String(i int) string {
	return strconv.Itoa(i)
}
