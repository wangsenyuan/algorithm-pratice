package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var t int
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &t)

	for i := 0; i < t; i++ {
		scanner.Scan()
		s := scanner.Bytes()
		can, res := solve(s)
		if !can {
			fmt.Println("-1")
		} else {
			fmt.Printf("%d", res[0]+1)
			for j := 1; j < len(s); j++ {
				fmt.Printf(" %d", res[j]+1)
			}
			fmt.Println()
		}
	}
}

func solve(s []byte) (bool, []int) {
	n := len(s)

	cnt := make([]int, 26)

	for i := 0; i < n; i++ {
		cnt[s[i]-'a']++
	}

	var odds int

	for i := 0; i < 26; i++ {
		if cnt[i]%2 == 1 {
			odds++
		}
	}
	if odds > 1 {
		return false, nil
	}

	res := make([]int, n)
	left, prev := 0, 0

	for x := 0; x < 26; x++ {
		if cnt[x] > 0 {
			for i := 0; i < n; i++ {
				if s[i] == byte(x+'a') {
					if cnt[x]%2 == 1 {
						res[n/2] = i
						cnt[x] = 0
					} else if left == 0 {
						res[prev] = i
						left = 1
						prev++
					} else {
						res[n-prev] = i
						left = 0
					}
				}
			}
		}
	}

	return true, res
}
