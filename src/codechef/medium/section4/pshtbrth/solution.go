package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	sign := int64(1)
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := int64(0)
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		N, M := readTwoNums(scanner)
		A := make([]int, N)
		for i := 0; i < N; i++ {
			S := make([]string, 4)
			for j := 0; j < 4; j++ {
				scanner.Scan()
				S[j] = scanner.Text()
			}
			A[i] = encode(S)
			scanner.Scan()
		}
		queries := make([][]int, M)
		for i := 0; i < M; i++ {
			queries[i] = make([]int, 3)
			scanner.Scan()
			var tp int
			pos := readInt(scanner.Bytes(), 0, &tp)
			queries[i][0] = tp
			if tp == 1 {
				pos = readInt(scanner.Bytes(), pos+1, &queries[i][1])
				readInt(scanner.Bytes(), pos+1, &queries[i][2])
			} else {
				readInt(scanner.Bytes(), pos+1, &queries[i][1])
				S := make([]string, 4)
				for j := 0; j < 4; j++ {
					scanner.Scan()
					S[j] = scanner.Text()
				}
				queries[i][2] = encode(S)
			}
		}
		res := solve(N, M, A, queries)
		for _, ans := range res {
			if ans == 1 {
				fmt.Println("Pishty")
			} else {
				fmt.Println("Lotsy")
			}
		}
	}
}

const MASK = 16

var dp []int

func init() {
	dp = make([]int, 1<<MASK)
	for i := 0; i < 1<<MASK; i++ {
		dp[i] = -1
	}
	dp[0] = 0

	var loop func(state int) int
	cnt := make([][]int, 4)
	for i := 0; i < 4; i++ {
		cnt[i] = make([]int, 4)
	}
	loop = func(state int) int {
		if dp[state] < 0 {
			subs := make([]int, 0, 8)
			//row := (state >> uint(4*(3-c))) & 15
			for i := 0; i < 4; i++ {
				for j := 0; j < 4; j++ {
					row := (state >> uint(4*(3-i))) & 15
					cell := row >> uint(3-j) & 1
					cnt[i][j] = cell
					if i > 0 {
						cnt[i][j] += cnt[i-1][j]
					}
					if j > 0 {
						cnt[i][j] += cnt[i][j-1]
					}
					if i > 0 && j > 0 {
						cnt[i][j] -= cnt[i-1][j-1]
					}
				}
			}
			for i := 0; i < 4; i++ {
				for j := 0; j < 4; j++ {
					for a := i; a < 4; a++ {
						for b := j; b < 4; b++ {
							c := cnt[a][b]
							if i > 0 {
								c -= cnt[i-1][b]
							}
							if j > 0 {
								c -= cnt[a][j-1]
							}
							if i > 0 && j > 0 {
								c += cnt[i-1][j-1]
							}
							if c == (b-j+1)*(a-i+1) {
								// all ones
								next := state
								for u := i; u <= a; u++ {
									for v := j; v <= b; v++ {
										next ^= (1 << uint(4*(3-u)+3-v))
									}
								}
								subs = append(subs, loop(next))
							}
						}
					}
				}
			}
			sort.Ints(subs)
			var res int
			for i := 0; i < len(subs); i++ {
				if res < subs[i] {
					break
				}
				res++
			}
			dp[state] = res
		}
		return dp[state]
	}

	for i := 0; i < 1<<MASK; i++ {
		loop(i)
	}
}

func solve(N, M int, A []int, queries [][]int) []int {
	st := NewSegTree(A)
	res := make([]int, M)
	var j int
	for _, query := range queries {
		tp := query[0]
		if tp == 1 {
			L, R := query[1]-1, query[2]-1
			tmp := st.Get(L, R)
			if tmp > 0 {
				res[j] = 1
			}
			j++
		} else {
			idx, val := query[1]-1, query[2]
			st.Update(idx, val)
		}
	}
	return res[:j]
}

func encode(S []string) int {
	var res int

	for i := 3; i >= 0; i-- {
		s := S[i]
		for j := 3; j >= 0; j-- {
			if s[j] == '1' {
				res |= 1 << uint(4*(3-i)+3-j)
			}
		}
	}

	return res
}

type SegTree struct {
	arr  []int
	size int
}

func NewSegTree(in []int) SegTree {
	n := len(in)
	arr := make([]int, 4*n)
	var loop func(i int, start, end int)

	loop = func(i int, start, end int) {
		if start == end {
			arr[i] = in[start]
			return
		}
		mid := (start + end) >> 1
		loop(2*i+1, start, mid)
		loop(2*i+2, mid+1, end)
		arr[i] = arr[2*i+1] ^ arr[2*i+2]
	}
	loop(0, 0, n-1)
	return SegTree{arr, n}
}

func (st *SegTree) Update(pos int, val int) {
	arr := st.arr
	var loop func(i int, start, end int)
	loop = func(i int, start, end int) {
		if start == end {
			arr[i] = val
			return
		}
		mid := (start + end) >> 1
		if pos <= mid {
			loop(2*i+1, start, mid)
		} else {
			loop(2*i+2, mid+1, end)
		}
		arr[i] = arr[2*i+1] ^ arr[2*i+2]
	}
	loop(0, 0, st.size-1)
}

func (st SegTree) Get(left, right int) int {
	arr := st.arr
	var loop func(i int, start, end int) int
	loop = func(i int, start, end int) int {
		if end < left || right < start {
			return 0
		}
		if left <= start && end <= right {
			return arr[i]
		}
		mid := (start + end) >> 1
		a := loop(2*i+1, start, mid)
		b := loop(2*i+2, mid+1, end)
		return a ^ b
	}
	return loop(0, 0, st.size-1)
}
