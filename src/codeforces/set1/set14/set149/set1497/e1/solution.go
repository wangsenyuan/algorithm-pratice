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
		n, k := readTwoNums(reader)
		a := readNNums(reader, n)
		res := solve(a, k)
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

const X = 1_000_000_9

var lps [X]int

func init() {
	var primes []int
	for x := 2; x < X; x++ {
		if lps[x] == 0 {
			lps[x] = x
			primes = append(primes, x)
		}
		for j := 0; j < len(primes); j++ {
			if x*primes[j] >= X {
				break
			}
			lps[x*primes[j]] = primes[j]
			if x%primes[j] == 0 {
				break
			}
		}
	}
}

func solve(a []int, k int) int {
	// k = 0
	n := len(a)
	// 也就是在一个segment里面， a[i] * a[j] is not a square
	//  => 它们的质数因子的和不是偶数
	// 所以假设a[i] * a[j] 是一个square，那么它们必须分成两组
	// 假设a[i]的质因子的向量是（2:x[2], 3:x[3], .....)
	// 只应该考虑那些奇数个数的质因子，且必须刚好和它相同的部分
	dp := make([]int, n)

	prev := make(map[int]int)

	freq := make(map[int]int)
	for i, num := range a {
		if i > 0 {
			// 和它一组
			dp[i] = dp[i-1]
		}
		clear(freq)
		for num > 1 {
			freq[lps[num]]++
			num /= lps[num]
		}
		for k, v := range freq {
			if v&1 == 1 {
				num *= k
			}
		}
		if j, ok := prev[num]; ok {
			dp[i] = max(dp[i], dp[j]+1)
		}
		prev[num] = i
	}

	return dp[n-1] + 1
}
