package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	a := readNNums(reader, n)

	res := solve(a)

	if len(res) == 0 {
		fmt.Println(-1)
		return
	}
	fmt.Println(len(res))
	s := fmt.Sprintf("%v", res)
	s = s[1 : len(s)-1]
	fmt.Println(s)
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

func solve(a []int) []int {
	n := len(a)
	next := make([]int, n)
	deg := make([]int, n)
	for i := 0; i < n; i++ {
		next[i] = a[i] - 1
		deg[next[i]]++
	}
	color := make([]int, n)
	que := make([]int, n)
	var front, tail int
	for i := 0; i < n; i++ {
		if deg[i] == 0 {
			que[front] = i
			front++
		}
	}
	for tail < front {
		u := que[tail]
		tail++
		v := next[u]
		if color[u] == 0 {
			color[v] = 1
		}
		deg[v]--
		if deg[v] == 0 {
			que[front] = v
			front++
		}
	}
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		if deg[i] > 0 {
			u := i
			var p int
			var check bool
			for {
				if color[u] == 1 {
					check = true
				}
				arr[p] = u
				p++
				deg[u] = 0
				u = next[u]
				if u == i {
					break
				}
			}

			if !check && p%2 == 1 {
				return nil
			}
			if !check {
				// 偶数长度的圈，且没有设置
				for j := 0; j < p; j += 2 {
					color[arr[j]] = 1
				}
				continue
			}
			prev := -1
			for j := 0; j < 2*p; j++ {
				k := j % p
				if arr[k] < 0 {
					break
				}
				if color[arr[k]] == 1 {
					prev = k
				} else {
					if prev >= 0 {
						if (j-prev)%2 == 0 {
							color[arr[k]] = 1
						}
						arr[k] = -1
					}
				}
			}
		}
	}
	var res []int

	for i := 0; i < n; i++ {
		if color[i] == 0 {
			res = append(res, next[i]+1)
		}
	}
	return res
}

type Pair struct {
	first  int
	second int
}
