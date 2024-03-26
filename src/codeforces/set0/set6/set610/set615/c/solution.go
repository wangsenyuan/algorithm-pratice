package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	a := readString(reader)
	b := readString(reader)
	res := solve(a, b)

	if len(res) == 0 {
		fmt.Println(-1)
		return
	}
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d %d\n", x[0], x[1]))
	}

	fmt.Print(buf.String())
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

func solve(s string, t string) [][]int {
	// dp[i] 表示到t[i]为止能够获得的最少的操作数
	// dp[i] = dp[j] + 1 if t[j+1...i]是s的一个子串（反字串）
	n := len(s)
	forward := make([]map[Key]int, n+1)
	back := make([]map[Key]int, n+1)

	for i := 1; i <= n; i++ {
		forward[i] = make(map[Key]int)
		back[i] = make(map[Key]int)
	}

	for i := 0; i < n; i++ {
		var key Key
		for j := i; j < n; j++ {
			x := int(s[j] - 'a')
			key = key.Add(x)
			forward[j-i+1][key] = i
		}
		key = Key{0, 0}
		for j := i; j >= 0; j-- {
			x := int(s[j] - 'a')
			key = key.Add(x)
			back[i-j+1][key] = j
		}
	}

	m := len(t)
	dp := make([][]int, m)

	update := func(i int, j int, l, r int) {
		if j < 0 {
			dp[i] = []int{1, -1, l, r}
			return
		}
		v := dp[j]
		if v[0] < 0 {
			return
		}
		if dp[i][0] < 0 || v[0]+1 < dp[i][0] {
			dp[i] = []int{v[0] + 1, j, l, r}
		}
	}

	for i := 0; i < m; i++ {
		dp[i] = []int{-1, -1, -1, -1}
		var key Key
		for j := i; j >= 0; j-- {
			x := int(t[j] - 'a')
			key = key.Add(x)
			ln := i - j + 1
			if ln > n {
				// too much
				break
			}

			if k, ok := back[ln][key]; ok {
				update(i, j-1, k+1, k+ln)
			}
			if k, ok := forward[i-j+1][key]; ok {
				update(i, j-1, k+ln, k+1)
			}
		}
	}

	if dp[m-1][0] < 0 {
		// no answer
		return nil
	}

	var res [][]int

	for i := m - 1; i >= 0; {
		v := dp[i]
		res = append(res, []int{v[2], v[3]})
		i = v[1]
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return res
}

const M = 1000000007

const B1 = 29
const B2 = 31

type Key struct {
	first  int64
	second int64
}

func (this Key) Add(v int) Key {
	first := (this.first*B1%M + int64(v)) % M
	second := (this.second*B2%M + int64(v)) % M
	return Key{first, second}
}
