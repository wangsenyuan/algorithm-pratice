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
		n := readNum(reader)
		a := readNNums(reader, n)
		res := solve(a)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

const MOD = 1e9 + 7
const MAX_A = 100_010

func modMul(a, b int) int {
	return int(int64(a) * int64(b) % MOD)
}

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func modSub(a, b int) int {
	return modAdd(a, MOD-b)
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = modMul(r, a)
		}
		a = modMul(a, a)
		b >>= 1
	}
	return r
}

func inverse(a int) int {
	return pow(a, MOD-2)
}

var F []int

func init() {
	F = make([]int, MAX_A)
	F[0] = 1
	for i := 1; i < MAX_A; i++ {
		F[i] = modMul(F[i-1], i)
	}
}

func solve(A []int) int {
	// 如果 a[i] > 0 mex(i) = 0
	// 如果 a[i] = 0, mex(i) = 1
	//  mex(A[l...r]) = mex(B[l...r])
	// 如果A[0] = 0, 那么 B[0] = ? 也必须为0
	// 如果在 l...r 中间没有0，其实对中间的数是没有限制的
	// 如果 A[r+1] = 0, 那么B[r+1] = 0
	// 0的位置要是一样的，然后考虑1的位置，是否也要是一样的，也必须是一样的，因为这样才能保证0和1之间的mex=1
	// 然后是2的位置，如果2在0， 1中间是没关系的，否则的2的位置也必须是一致的
	n := len(A)
	pos := make([]int, n)
	for i, x := range A {
		pos[x] = i
	}
	ans := 1
	l, r := pos[0], pos[0]
	for mex := 1; mex < n; mex++ {
		if pos[mex] < l {
			l = pos[mex]
		} else if pos[mex] > r {
			r = pos[mex]
		} else {
			//zai l ... r 中间，已经被确定了 mex个数 (0...mex-1)
			// mex可以放的位置在剩余的位置上
			ans = modMul(ans, r-l+1-mex)
		}
	}

	return ans
}
