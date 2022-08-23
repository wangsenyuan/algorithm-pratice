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
		s := readString(reader)
		res := solve(s)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

const MOD = 1000000007

const N = 5010

var C [][]int

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

func modMul(a, b int) int {
	r := int64(a) * int64(b)
	return int(r % MOD)
}

func init() {
	C = make([][]int, N)
	for i := 0; i < N; i++ {
		C[i] = make([]int, N)
		C[i][0] = 1
		for j := 1; j <= i; j++ {
			C[i][j] = modAdd(C[i-1][j-1], C[i-1][j])
		}
	}
}

func solve(s string) int {
	n := len(s)

	if n == 1 {
		return 1
	}
	cnt := make([]int, 2)
	for i := 0; i < len(s); i++ {
		if s[i] == '4' {
			cnt[0]++
		} else {
			cnt[1]++
		}
	}

	if cnt[0] == 0 {
		return 1
	}

	res := C[n][cnt[0]]
	res = modSub(res, C[n-2][cnt[0]-1])
	return res
}

func solve1(s string) int {
	cnt := make([]int, 2)
	for i := 0; i < len(s); i++ {
		if s[i] == '4' {
			cnt[0]++
		} else {
			cnt[1]++
		}
	}

	// 如何将两次操作中的 74  和 74 区分开呢？
	// 74 既可以是中间位置是4，且 x = 1, 或者是中间位置是7，且x = 0
	var res int
	// (cnt[1]) 4 (cnt[0] - 1)
	for x := 0; x <= cnt[0]-1 && x <= cnt[1]; x++ {
		// put x 4 in prefix
		// 如果 x = 0, prefix 没有4，prefix = cnt[1]个7， 结构就是 7774...
		// 如果 x = cnt[1], prefix就没有7, 只有4  (4)_cnt1 4 cnt1(7) + (4)(cnt0 - 1 - cnt1)
		// y := cnt[1] - x
		// then there need y 7 in the prefix
		a := C[cnt[1]][x]
		b := C[cnt[0]-1][x]
		res = modAdd(res, modMul(a, b))
	}

	for x := 0; x <= cnt[0] && x <= cnt[1]-1; x++ {
		// put x 4 in prefix
		// 如果 x = 0, prefix 没有4，prefix = cnt[1]个7， 结构就是 7774...
		// 如果 x = cnt[1], prefix就没有7, 只有4  (4)_cnt1 4 cnt1(7) + (4)(cnt0 - 1 - cnt1)
		// y := cnt[1] - x
		// then there need y 7 in the prefix
		a := C[cnt[1]-1][x]
		b := C[cnt[0]][x]
		res = modAdd(res, modMul(a, b))
	}

	if cnt[0] > 0 && cnt[1] > 0 {
		// put 74 at mid
		for x := 0; x <= cnt[0]-1 && x <= cnt[1]-1; x++ {
			a := C[cnt[1]-1][x]
			b := C[cnt[0]-1][x]
			res = modSub(res, modMul(a, b))
		}
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
