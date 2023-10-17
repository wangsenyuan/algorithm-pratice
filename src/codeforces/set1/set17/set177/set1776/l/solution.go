package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	s := readString(reader)[:n]
	m := readNum(reader)
	queries := make([][]int, m)
	for i := 0; i < m; i++ {
		queries[i] = readNNums(reader, 2)
	}
	res := solve(s, queries)
	var buf bytes.Buffer

	for _, cur := range res {
		if cur {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
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

func solve(s string, queries [][]int) []bool {
	// 对于a, b
	// 假设一种有x次+， y次-
	// 其中 按了xa次+a, xb次+b, ya次-a, yb次-b
	// xa + xb = x
	// ya + yb = y
	// xa * a + xb * b - ya * a - yb * b = 0
	// (xa - ya) * a = (yb - xb) * b
	// gcd(a, b) = g
	// 这里迭代xa
	// 其他的都可以算出来
	cnt := make([]int, 2)
	for i := 0; i < len(s); i++ {
		if s[i] == '+' {
			cnt[1]++
		} else {
			cnt[0]++
		}
	}
	ans := make([]bool, len(queries))

	if cnt[0] == cnt[1] {
		for i := 0; i < len(queries); i++ {
			ans[i] = true
		}
		return ans
	}

	qs := make(map[Pair][]int)

	for i, cur := range queries {
		a, b := cur[0], cur[1]
		if a > b {
			a, b = b, a
		}
		// a <= b
		c := gcd(a, b)
		p := Pair{a / c, b / c}

		if len(qs[p]) == 0 {
			qs[p] = make([]int, 0, 1)
		}
		qs[p] = append(qs[p], i)
	}

	// 假设xa是a的净使用量 （可正，可负)
	// xa * a + xb * b = 0
	// 假设xa是正数
	// -cnt[0] <= xa <= cnt[1]
	// 对于确定的某个xa
	// xb似乎不是确定的
	// -1 <= xa <= 3
	// 就是最多+++-
	//  如果xa = 2, 有两种可能性 一种是 ++ / +-  +++-/
	// 所以这里将抵消的部分给确定下来，不管是a抵消还是b抵消，都是抵消了
	// 假设 z ( <= min(cnt[0], cnt[1]))个抵消了，剩余的就是全部是 +a, -b

	for z := 0; z <= cnt[0] && z <= cnt[1]; z++ {
		// z个全部抵消了
		x := cnt[0] - z
		y := cnt[1] - z
		// x个a，y个b
		g := gcd(x, y)

		p := Pair{x / g, y / g}
		for _, i := range qs[p] {
			ans[i] = true
		}
		delete(qs, p)
		if x != y {
			p = Pair{y / g, x / g}
			for _, i := range qs[p] {
				ans[i] = true
			}
			delete(qs, p)
		}
	}

	return ans
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

type Pair struct {
	first  int
	second int
}

type Query struct {
	id  int
	val Pair
}
