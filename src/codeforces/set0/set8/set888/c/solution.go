package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	s := readString(reader)

	res := solve(s)

	fmt.Println(res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
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

func solve(s string) int {
	// only need to check one letter
	// n / k = 1, k > n / 2, 这时候肯定是成立的，因为中间部分，都会出现
	//  就是找第一类的答案
	var best int = (len(s)+1)/2 + 1
	for i := 0; i < 26; i++ {
		best = min(best, solve1(s, i))
	}
	return best
}

func solve1(s string, x int) int {
	n := len(s)

	check := func(k int) bool {
		var last int
		for i := 0; i+k <= n; i++ {
			if last < i {
				// already invalid
				last = i
			}
			for last < i+k && int(s[last]-'a') != x {
				last++
			}
			if last == i+k {
				return false
			}
			// else last < i + k  && s[last] == x
		}
		return true
	}

	l, r := 1, n
	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
