package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)

	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n := readNum(reader)
		a := readNNums(reader, n)
		res := solve(a)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

const mod = 998244353

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return a * b % mod
}

const N = 200010

var F [N]int
var I [N]int

func init() {
	F[0] = 1
	for i := 1; i < N; i++ {
		F[i] = mul(i, F[i-1])
	}

	I[N-1] = pow(F[N-1], mod-2)
	for i := N - 2; i >= 0; i-- {
		I[i] = mul(i+1, I[i+1])
	}
}

func pow(a, b int) int {
	var R int = 1
	for b > 0 {
		if b&1 == 1 {
			R = (R * a) % mod
		}
		a = (a * a) % mod
		b >>= 1
	}
	return R
}

func nCr(n, r int) int {
	if n < r {
		return 0
	}
	return mul(F[n], mul(I[r], I[n-r]))
}

func solve(a []int) int {
	sort.Ints(a)
	reverse(a)

	var cnt int

	n := len(a)
	for cnt < n && a[cnt] == a[0] {
		cnt++
	}

	if cnt == n {
		return F[cnt]
	}
	if cnt == 1 {
		if a[cnt] != a[0]-1 {
			return 0
		}
		// 只要在a[0]后面有一个 a[cnt]就可以了
		var cnt2 int
		for i := cnt; i < n && a[i] == a[cnt]; i++ {
			cnt2++
		}
		cnt += cnt2
		res := nCr(n, cnt)
		// 但是最后一个位置是给a[cnt]的
		res = mul(res, cnt2)
		// 剩下的随便排
		res = mul(res, F[cnt-1])
		res = mul(res, F[n-cnt])
		return res
	}
	// 随便排
	return F[n]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
