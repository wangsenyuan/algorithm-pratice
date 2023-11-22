package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		a := make([][]int, n)
		for i := 0; i < n; i++ {
			a[i] = readNNums(reader, m)
		}
		res := solve(n, a)
		s := fmt.Sprintf("%v", res)
		s = s[1 : len(s)-1]
		buf.WriteString(s)
		buf.WriteByte('\n')
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

func solve(n int, a [][]int) []int {
	m := len(a[0])
	// r = p * q => r[j] = q[p[j]]
	// 这里重要的是1的位置，2的位置，3的位置....
	// 对于p来说，如果
	// 	r[1] = 1, 那么 q[p[1]] = 1
	//  r[2] = 2,      q[p[2]] = 2
	// 所以对于p来说 p[1] = q[pos[1]]
	// 所以想要反过来
	// 所以这里要hash吗
	freq := make(map[Key]int)
	pos := make([]int, m+1)

	for i := 0; i < n; i++ {

		for j, x := range a[i] {
			pos[x] = j + 1
		}

		var key Key
		for j := 1; j <= m; j++ {
			key = key.Add(pos[j])
			freq[key]++
		}
	}

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		// 没说p != q, 不用排除自己
		// 最长是m的 m * freq[key]
		//  (m - 1) * freq[key]
		var key Key
		for j := 1; j <= m; j++ {
			// 假设目前有x个长度 >= j, 之前是有prev个 >= j - 1
			// 那么共有 prev - x个的长度肯定是 j - 1
			key = key.Add(a[i][j-1])
			// 这里还有个问题，就是最大的k怎么计算呢？
			if freq[key] > 0 {
				ans[i] = j
			}
		}
	}

	return ans
}

const M = 100000000007

const B1 = 29
const B2 = 31

type Key struct {
	first  int
	second int
}

func (this Key) Add(v int) Key {
	first := (this.first*B1%M + v) % M
	second := (this.second*B2%M + v) % M
	return Key{first, second}
}

func (this Key) Mul(a, b int) Key {
	first := this.first * a % M
	second := this.second * b % M
	return Key{first, second}
}

func (this Key) Mul2(that Key) Key {
	first := this.first * that.first % M
	second := this.second * that.second % M
	return Key{first, second}
}

func (this Key) Sub(that Key) Key {
	first := (this.first + M - that.first) % M
	second := (this.second + M - that.second) % M
	return Key{first, second}
}
