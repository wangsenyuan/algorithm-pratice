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
		nums := readNNums(reader, 4)
		x, y := solve(nums)
		buf.WriteString(fmt.Sprintf("%d %d\n", x, y))
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

const X = 2_000_000

var primes []int

func init() {
	vis := make([]bool, X)

	for i := 2; i < X; i++ {
		if !vis[i] {
			primes = append(primes, i)
			vis[i] = true
		}
		for j := 0; j < len(primes); j++ {
			if i*primes[j] >= X {
				break
			}
			vis[i*primes[j]] = true
			if i%primes[j] == 0 {
				break
			}
		}
	}
}

func solve(nums []int) (int, int) {
	a, b, c, d := nums[0], nums[1], nums[2], nums[3]

	fa := factors(a)
	fb := factors(b)

	for k, v := range fb {
		fa[k] += v
	}

	fs := make([]Pair, 0, len(fa))
	for k, v := range fa {
		fs = append(fs, Pair{k, v})
	}

	var dfs func(x int, now int64) bool

	aa := int64(a)
	bb := int64(b)
	cc := int64(c)
	dd := int64(d)
	tmp := aa * bb

	var ans1, ans2 int64

	dfs = func(x int, now int64) bool {
		t1 := (aa + 1 + now - 1) / now * now
		t2 := (bb + 1 + tmp/now - 1) / (tmp / now) * (tmp / now)
		if t1 <= cc && t2 <= dd {
			ans1 = t1
			ans2 = t2
			return true
		}
		t1 = (aa + 1 + tmp/now - 1) / (tmp / now) * (tmp / now)
		t2 = (bb + 1 + now - 1) / now * now
		if t1 <= cc && t2 <= dd {
			ans1 = t1
			ans2 = t2
			return true
		}
		if x >= len(fs) {
			return false
		}
		l := fs[x].second
		v := fs[x].first
		for i := 0; i <= l; i++ {
			if dfs(x+1, now) {
				return true
			}
			now *= int64(v)
		}
		return false
	}

	if dfs(0, 1) {
		return int(ans1), int(ans2)
	}
	return -1, -1
}

type Pair struct {
	first  int
	second int
}

func factors(num int) map[int]int {
	res := make(map[int]int)

	for j := 0; j < len(primes) && primes[j]*primes[j] <= num; j++ {
		for num%primes[j] == 0 {
			res[primes[j]]++
			num /= primes[j]
		}
	}

	if num > 1 {
		res[num]++
	}

	return res
}

func gcd(a, b int64) int64 {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
