package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	s := readString(reader)
	p := readString(reader)
	res := solve(s, p)
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

func solve(s string, p string) int {
	if len(s) < len(p) {
		return 0
	}
	freq := make([]int, 26)
	for i := 0; i < len(p); i++ {
		freq[int(p[i]-'a')]++
	}

	freq2 := make([]int, 27)

	add := func(i int) {
		if s[i] == '?' {
			freq2[26]++
		} else {
			freq2[int(s[i]-'a')]++
		}
	}

	rem := func(i int) {
		if s[i] == '?' {
			freq2[26]--
		} else {
			freq2[int(s[i]-'a')]--
		}
	}

	check := func() bool {
		tmp := freq2[26]
		for i := 0; i < 26; i++ {
			if freq2[i] > freq[i] {
				// too much
				return false
			}
			if freq2[i] == freq[i] {
				continue
			}
			tmp -= freq[i] - freq2[i]
			if tmp < 0 {
				return false
			}
		}
		return true
	}

	k := len(p)
	var res int
	for i := 0; i < len(s); i++ {
		add(i)
		if i-k >= 0 {
			rem(i - k)
		}
		if i >= k-1 && check() {
			res++
		}
	}

	return res
}
