package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(A)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
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

const MOD = 1000000007

func solve(A []int) int {
	n := len(A)
	arr := make([]int, 0, n)

	for i := 0; i < n && i < 2; i++ {
		arr = append(arr, A[i])
	}

	// 如果有多于2个的连续数字，保留2个就足够了
	for i := 2; i < n; i++ {
		m := len(arr)
		if A[i] != A[m-1] || A[i] != arr[m-2] {
			arr = append(arr, A[i])
		}
	}
	n = len(arr)
	dp := make([]Record, n+1)
	dp[0] = Record{0, 0}
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1]
		seen := make(map[int]bool)
		for j := i - 1; j >= max(0, i-7); j-- {
			seen[arr[j]] = true
			if len(seen) == 2 {
				dp[i] = dp[i].Max(dp[j].IncrCnt2())
			} else if len(seen) == 3 {
				dp[i] = dp[i].Max(dp[j].IncrCnt3())
				break
			}
		}
	}

	ans := pow(2, dp[n].cnt2)
	ans *= pow(3, dp[n].cnt3)

	return int(ans % MOD)
}

func pow(a, b int) int64 {
	A := int64(a)
	R := int64(1)
	for b > 0 {
		if b&1 == 1 {
			R = (R * A) % MOD
		}
		A = (A * A) % MOD
		b >>= 1
	}
	return R
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

var ratio = 1.0 / math.Log2(3)

type Record struct {
	cnt3, cnt2 int
}

func (this Record) IncrCnt2() Record {
	return Record{this.cnt3, this.cnt2 + 1}
}

func (this Record) IncrCnt3() Record {
	return Record{this.cnt3 + 1, this.cnt2}
}

func (this Record) Max(that Record) Record {
	if this.Less(that) {
		return that
	}
	return this
}

// The product will not fit inside the numerical data types of most languages. However since the product if of form 2^x * 3^y, we can keep store the count of $x$ and $y$ instead of the product, and use logarithm to compare products.

// log(2 ^^ x * 3 ^^ y) = log(2 ^^ x) + log(3 ^^ y) = x + y * ratio
func (this Record) Less(that Record) bool {
	return float64(this.cnt3)+float64(this.cnt2)*ratio < float64(that.cnt3)+float64(that.cnt2)*ratio
}

type Records []Record

func (records Records) Len() int {
	return len(records)
}

func (records Records) Less(i, j int) bool {
	return records[i].Less(records[j])
}

func (records Records) Swap(i, j int) {
	records[i], records[j] = records[j], records[i]
}
