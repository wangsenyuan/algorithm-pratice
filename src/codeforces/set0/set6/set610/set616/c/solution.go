package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	mat := make([]string, n)
	for i := 0; i < n; i++ {
		mat[i] = readString(reader)[:m]
	}
	res := solve(mat)
	for _, row := range res {
		fmt.Println(row)
	}
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

func solve(mat []string) []string {
	n := len(mat)
	m := len(mat[0])

	set := NewUFSet(n * m)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if mat[i][j] == '.' {
				if i > 0 && mat[i-1][j] == '.' {
					set.Union((i-1)*m+j, i*m+j)
				}
				if j > 0 && mat[i][j-1] == '.' {
					set.Union(i*m+j-1, i*m+j)
				}
			}
		}
	}

	res := make([][]byte, n)
	for i := 0; i < n; i++ {
		res[i] = make([]byte, m)
		for j := 0; j < m; j++ {
			res[i][j] = '.'
			if mat[i][j] == '*' {
				cnt := make(map[int]int)
				if i > 0 && mat[i-1][j] == '.' {
					p := set.Find((i-1)*m + j)
					cnt[p] = set.size[p]
				}
				if i+1 < n && mat[i+1][j] == '.' {
					p := set.Find((i+1)*m + j)
					cnt[p] = set.size[p]
				}
				if j > 0 && mat[i][j-1] == '.' {
					p := set.Find(i*m + j - 1)
					cnt[p] = set.size[p]
				}
				if j+1 < m && mat[i][j+1] == '.' {
					p := set.Find(i*m + j + 1)
					cnt[p] = set.size[p]
				}
				tmp := 1
				for _, v := range cnt {
					tmp += v
				}
				tmp %= 10
				res[i][j] = byte(tmp + '0')
			}
		}
	}
	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = string(res[i])
	}

	return s
}

type UFSet struct {
	arr  []int
	size []int
	cnt  int
}

func NewUFSet(n int) *UFSet {
	arr := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		size[i] = 1
	}

	return &UFSet{arr, size, n}
}

func (set *UFSet) Find(x int) int {
	if set.arr[x] != x {
		set.arr[x] = set.Find(set.arr[x])
	}
	return set.arr[x]
}

func (set *UFSet) Union(a, b int) bool {
	pa := set.Find(a)
	pb := set.Find(b)
	if pa == pb {
		return false
	}
	set.cnt--
	if set.size[pa] < set.size[pb] {
		pa, pb = pb, pa
	}
	set.size[pa] += set.size[pb]
	set.arr[pb] = pa
	return true
}
