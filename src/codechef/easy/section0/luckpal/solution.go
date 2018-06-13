package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bs []byte, from int, val *int) int {
	tmp := 0
	i := from
	for i < len(bs) && bs[i] != ' ' {
		tmp = tmp*10 + int(bs[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var t int
	readInt(scanner.Bytes(), 0, &t)

	for i := 0; i < t; i++ {
		scanner.Scan()
		bs := scanner.Text()
		pal, times, can := solve(bs)
		if !can {
			fmt.Println("unlucky")
		} else {
			fmt.Printf("%s %d\n", pal, times)
		}
	}
}

func solve(str string) (string, int, bool) {
	n := len(str)
	if n < 9 {
		return "", -1, false
	}

	lucky := []byte("lucky")
	var ans string
	changes := -1
	for i := 0; i+4 < n; i++ {
		bs := []byte(str)
		cnt := 0
		for j := i; j <= i+4; j++ {
			if bs[j] != lucky[j-i] {
				cnt++
			}
			bs[j] = lucky[j-i]
		}

		j, k := 0, n-1

		for ; j < k; j, k = j+1, k-1 {
			if bs[j] != bs[k] {
				if j >= i && j < i+5 {
					if k >= i && k < i+5 {
						// in the same range, can't mirror
						cnt = 100000
						break
					}
					bs[k] = bs[j]
				} else if k >= i && k < i+5 {
					if j >= i && j < i+5 {
						cnt = 100000
						break
					}
					bs[j] = bs[k]
				} else if bs[j] < bs[k] {
					bs[k] = bs[j]
				} else {
					bs[j] = bs[k]
				}
				cnt++
			}
		}

		tmp := string(bs)
		if changes < 0 || changes > cnt || (changes == cnt && tmp < ans) {
			ans = tmp
			changes = cnt
		}
	}

	return ans, changes, true
}
