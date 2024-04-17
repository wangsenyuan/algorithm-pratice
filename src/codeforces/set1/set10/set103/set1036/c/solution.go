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
		l, r := readTwoNums(reader)
		res := solve(l, r)
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

func nCr(n int, r int) int {
	if n < r {
		return 0
	}

	if r == 0 || n == r {
		return 1
	}
	if r == 1 {
		return n
	}
	if r == 2 {
		return n * (n - 1) / 2
	}
	return n * (n - 1) * (n - 2) / 6
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r *= a
		}
		a *= a
		b >>= 1
	}
	return r
}

const H = 19

var fp [H]int

func init() {
	for i := 1; i < H; i++ {
		fp[i] = fp[i-1]
		// 第一个数有非0，再选y个非0的数
		for y := 0; y < 3 && y < i; y++ {
			// y non-zeros, others are 0
			fp[i] += nCr(i-1, y) * pow(9, y+1)
		}
	}
}

func solve(L int, R int) int {
	return count(R) - count(L-1)
}

func count(num int) int {
	if num < 1000 {
		return num
	}
	// dp[i] 表示已经有i个非0的数字了
	var arr []int
	for tmp := num; tmp > 0; tmp /= 10 {
		arr = append(arr, tmp%10)
	}
	reverse(arr)
	n := len(arr)
	res := fp[n-1]

	var cnt int
	for i := 0; i < n; i++ {
		// 当前位可以选取 < arr[i]的数
		var s int
		if i == 0 {
			s = 1
		}

		for x := s; x < arr[i]; x++ {
			tmp := cnt
			if x > 0 {
				tmp++
			}
			// 后面的数里面可以选择y个非0的数
			for y := 0; y+tmp <= 3; y++ {
				res += pow(9, y) * nCr(n-i-1, y)
			}
		}

		if arr[i] != 0 {
			cnt++
		}
		if cnt == 4 {
			break
		}
	}

	if cnt <= 3 {
		res++
	}

	return res
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
