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
		A := readNNums(reader, n)
		B := readNNums(reader, n)
		res := solve(n, A, B)
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
		if s[i] == '\n' {
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

const INF = 1 << 60

func solve(n int, C []int, V []int) int64 {
	// dp[i] if Alice reach i first, the max A - B can get
	// fp[i] if Bob reach i first, the min A - B can get
	var dp int64
	var fp int64

	if C[n-1]%2 == 1 {
		dp = int64(V[n-1])
		fp = -int64(V[n-1])
	}

	for i := n - 2; i >= 0; i-- {
		//dp[i] = min of
		// V[i] + min value (fp[i+1...n-1]) Bob take another pile after i
		// if C[i] > 1, 0 + max(dp[i+1...n-1]) Bob pick next from pile i, and alice reach another pile after i
		// buf if min(fp) > max(dp), (?) when C[i] % 2 == 0, Alice still can get it
		pd := int64(V[i]) + fp
		if C[i] > 1 {
			// Bob take another from ith pile, cancel the value of A, and Alice take another chance
			pd = min(pd, dp)
		}
		if C[i]&1 == 1 && fp+int64(V[i]) > dp {
			// even let Bob to choose after i first, the result still better than Alice taking first, and Alice can force Bob to take the turn
			pd = max(pd, int64(V[i])+fp)
		}
		pf := -int64(V[i]) + dp

		if C[i] > 1 {
			// alice cancel it by taking aother from ith pile, and let Bob taking after i first
			pf = max(pf, fp)
		}
		if C[i]&1 == 1 && dp-int64(V[i]) < fp {
			pf = min(pf, dp-int64(V[i]))
		}
		dp = max(dp, pd)
		fp = min(fp, pf)
	}

	return dp
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
