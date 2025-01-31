package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, m, k := readThreeNums(reader)
	s := readString(reader)[:n]
	res := solve(m, k, s)
	fmt.Println(res)
}

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
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

func solve(m int, k int, s string) int64 {
	// repeat s m times, and remove k x's,

	process := func(s string, k int) int64 {
		var cnt int
		var ans int
		for l, r := 0, 0; l < len(s); l++ {
			for r < len(s) && cnt <= k {
				if s[r] == 'x' {
					if cnt == k {
						// should not cross r
						break
					}
					cnt++
				}
				r++
			}
			// cnt > k
			ans = max(ans, r-l)
			if s[l] == 'x' {
				cnt--
			}
		}
		return int64(ans)
	}

	if m == 1 {
		return process(s, k)
	}

	var v int
	for i := 0; i < len(s); i++ {
		if s[i] == 'x' {
			v++
		}
	}

	n := len(s)
	a, b := k/v, k%v

	ans := int64(a) * int64(n)
	ans += process(repeat(s, min(m-a, 2)), b)

	return ans
}

func repeat(s string, cnt int) string {
	buf := make([]byte, len(s)*cnt)
	for i := 0; i < cnt; i++ {
		copy(buf[len(s)*i:], s)
	}
	return string(buf)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max2(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
