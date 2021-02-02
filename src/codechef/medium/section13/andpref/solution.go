package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(n, A)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
}

func solve(n int, A []int) int64 {
	var bits int
	for 1<<uint(bits) < n {
		bits++
	}

	freq := make([]int, 1<<uint(bits))

	for _, x := range A {
		freq[x]++

	}

	for bit := 0; bit < bits; bit++ {
		for mask := 0; mask < 1<<uint(bits); mask++ {
			if mask&(1<<uint(bit)) == 0 {
				freq[mask] += freq[mask^1<<uint(bit)]
			}
		}
	}

	dp := make([]int64, 1<<uint(bits))
	dp[(1<<uint(bits))-1] = int64(freq[(1<<uint(bits))-1]) * int64((1<<uint(bits))-1)

	for mask := (1 << uint(bits)) - 1; mask >= 0; mask-- {
		for bit := 0; bit < bits; bit++ {
			if mask&(1<<uint(bit)) > 0 {
				dp[mask^(1<<uint(bit))] = max(dp[mask^(1<<uint(bit))], dp[mask]+int64(freq[mask^(1<<uint(bit))]-freq[mask])*int64(mask^(1<<uint(bit))))
			}
		}
	}

	return dp[0]
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
