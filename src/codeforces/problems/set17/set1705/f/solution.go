package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	query := func(s string) int {
		fmt.Println(s)
		x := readNum(reader)
		if x == n {
			os.Exit(0)
		}
		return x
	}

	solve(n, query)
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

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
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

func solve(n int, query func(string) int) {
	all_t := make([]byte, n)
	ans := make([]byte, n)
	for i := 0; i < n; i++ {
		all_t[i] = 'T'
		ans[i] = '?'
	}
	all_tf := make([]byte, n)
	copy(all_tf, all_t)
	for i := 1; i < n; i += 2 {
		all_tf[i] = 'F'
	}

	cnt_t := query(string(all_t))
	cnt_tf := query(string(all_tf))

	l, r := 0, n-1

	s := make([]byte, n)

	for l <= r {
		if l == r {
			copy(s, all_t)
			s[l] = 'F'
			k := query(string(s))
			if k > cnt_t {
				ans[l] = 'F'
			} else {
				ans[l] = 'T'
			}
			l++
			r--
		} else {
			copy(s, all_t)
			s[l] = 'F'
			s[l+1] = 'F'
			k := query(string(s)) - cnt_t
			if k == 2 {
				ans[l] = 'F'
				ans[l+1] = 'F'
				l += 2
			} else if k == -2 {
				ans[l] = 'T'
				ans[l+1] = 'T'
				l += 2
			} else {
				if r == l+1 {
					copy(s, all_t)
					s[l] = 'F'
					k := query(string(s))
					if k > cnt_t {
						ans[l] = 'F'
						ans[l+1] = 'T'
					} else {
						ans[l] = 'T'
						ans[l+1] = 'F'
					}
					l += 2
				} else {
					copy(s, all_tf)
					s[l] = 'F'
					s[l+1] = 'T'
					if s[r] == 'F' {
						s[r] = 'T'
					} else {
						s[r] = 'F'
					}
					k := query(string(s)) - cnt_tf
					if k == 3 {
						ans[l] = 'F'
						ans[l+1] = 'T'
						ans[r] = s[r]
					} else if k == 1 {
						ans[l] = 'F'
						ans[l+1] = 'T'
						ans[r] = all_tf[r]
					} else if k == -1 {
						ans[l] = 'T'
						ans[l+1] = 'F'
						ans[r] = s[r]
					} else {
						ans[l] = 'T'
						ans[l+1] = 'F'
						ans[r] = all_tf[r]
					}
					l += 2
					r--
				}
			}
		}
	}

	query(string(ans))
}
