package main

import (
	"bufio"
	"fmt"
	"os"
)

var status []string

func init() {
	status = []string{
		"1000000010",
		"1101100111",
		"0010000010",
		"0001000011",
		"0000100011",
		"0000011011",
		"0000001010",
		"1001000111",
		"0000000010",
		"0000000011",
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t := toInt(scanner.Bytes())
	for i := 0; i < t; i++ {
		scanner.Scan()
		bs := scanner.Bytes()
		i := 0
		for bs[i] != ' ' {
			i++
		}
		s := bs[:i]
		m := bs[i+1:]
		res := solve(s, m)
		fmt.Println(res)
	}
}

func valid(b byte, d int) bool {
	x := int(b - '0')
	return status[x][d] == '1'
}

func solve(S, M []byte) int {
	best := toInt(S)
	for k := 0; k <= len(M)-len(S); k++ {
		fit, mx := 0, -1
		for i := 0; i < len(M); i++ {
			nextFit, nextMx := -1, -1
			for j := 0; j < 10; j++ {
				if 0 <= i-k && i-k < len(S) && !valid(S[i-k], j) {
					continue
				}
				if nextMx < mx*10+j {
					nextMx = mx*10 + j
				}
				if j < int(M[i]-'0') && nextMx < fit*10+j {
					nextMx = fit*10 + j
				}
				if j == int(M[i]-'0') && nextFit < fit*10+j {
					nextFit = fit*10 + j
				}
			}
			fit, mx = nextFit, nextMx
			if i-k >= len(S)-1 && best < fit {
				best = fit
			}
			if i-k >= len(S)-1 && best < mx {
				best = mx
			}
		}
	}
	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func toInt(s []byte) int {
	var res int
	for i := 0; i < len(s); i++ {
		res = res*10 + int(s[i]-'0')
	}
	return res
}
