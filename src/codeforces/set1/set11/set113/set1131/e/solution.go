package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	p := make([]string, n)

	for i := range n {
		p[i] = readString(reader)
	}

	res := solve(p)

	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func solve(p []string) int {
	n := len(p)

	best, _, _ := get(p[n-1])

	if best != len(p[n-1]) {
		return easy(p)
	}

	return hard(p)
}

func hard(p []string) int {
	n := len(p)
	m := len(p[n-1])
	x := p[n-1][0]

	find := func(s string) (best int, pref int, suf int) {
		var cnt int
		k := len(s)
		for i := 0; i < k; i++ {
			if s[i] == x {
				cnt++
			} else {
				cnt = 0
			}

			if i+1 == cnt {
				pref = cnt
			}
			if i == k-1 {
				suf = cnt
			}

			best = max(best, cnt)
		}

		return
	}

	best := m
	pref := m
	suf := m
	whole := true
	for i := n - 2; i >= 0; i-- {
		b1, f1, s1 := find(p[i])
		if b1 == len(p[i]) {
			// 这个全部由字符x组成
			if whole {
				best = best*(b1+1) + b1
				pref = best
				suf = best
			} else {
				best = max(best, pref+1+suf)
			}
			continue
		}
		if b1 > 0 {
			// b1 > 0
			if whole {
				best = best*(b1+1) + b1
				pref += f1
				suf += s1
			} else {
				best = max(best, pref+1+suf)
			}
		}

		whole = false
	}

	return best
}

func easy(p []string) int {
	n := len(p)

	x := p[n-1]
	m := len(x)

	best, pref, suf := get(x)

	for _, r := range p {
		for i := 0; i < len(r); i++ {
			if r[i] == x[m-1] {
				best = max(best, suf+1)
			}
			if r[i] == x[0] {
				best = max(best, pref+1)
			}
			if r[i] == x[0] && r[i] == x[m-1] {
				best = max(best, pref+suf+1)
			}
		}
	}

	return best
}

func get(s string) (best int, pref int, suf int) {
	cnt := 1
	for i := 1; i <= len(s); i++ {
		if i == len(s) || s[i] != s[i-1] {
			if pref == 0 {
				pref = cnt
			}
			if i == len(s) {
				suf = cnt
			}
			best = max(best, cnt)
			cnt = 0
		}
		cnt++
	}
	return
}
