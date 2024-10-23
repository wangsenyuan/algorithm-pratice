package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	boxes := make([][]int, n)

	for i := 0; i < n; i++ {
		s, _ := reader.ReadBytes('\n')
		var m int
		pos := readInt(s, 0, &m) + 1
		boxes[i] = make([]int, m)
		for j := 0; j < m; j++ {
			pos = readInt(s, pos, &boxes[i][j]) + 1
		}
	}

	res := solve(boxes)

	if len(res) == 0 {
		fmt.Println("NO")
		return
	}

	var buf bytes.Buffer

	buf.WriteString("YES\n")
	for _, cur := range res {
		buf.WriteString(fmt.Sprintf("%d %d\n", cur[0], cur[1]))
	}

	fmt.Print(buf.String())
}

func solve(boxes [][]int) [][]int {
	n := len(boxes)
	var sum int

	box_sum := make([]int, n)

	for i, box := range boxes {
		for _, num := range box {
			sum += num
			box_sum[i] += num
		}
	}
	if sum%n != 0 {
		return nil
	}
	sum /= n

	comp := make(map[int]int)
	for i, box := range boxes {
		for _, num := range box {
			comp[num] = i
		}
	}

	give := make([]int, n)
	to := make([]int, n)
	vis := make([]bool, 1<<n)

	var f func(s int, p int) bool

	f = func(s int, p int) bool {
		if p == n {
			return true
		}

		if vis[s] {
			return false
		}

		if (s>>p)&1 == 1 {
			return f(s, p+1)
		}

		s |= 1 << p

	outer:
		for _, x := range boxes[p] {
			give[p] = x
			q, t, cur := p, s, x
			for {
				y := sum - box_sum[q] + cur
				if y == x {
					to[p] = q + 1
					break
				}
				if j, ok := comp[y]; !ok || (t>>j)&1 == 1 {
					continue outer
				} else {
					give[j] = y
					to[j] = comp[cur] + 1

					t |= 1 << j
					cur = y
					q = j
				}
			}

			if f(t, p+1) {
				return true
			}
			vis[t] = true
		}
		return false
	}

	if !f(0, 0) {
		return nil
	}

	ans := make([][]int, n)

	for i := 0; i < n; i++ {
		ans[i] = []int{give[i], to[i]}
	}

	return ans
}
