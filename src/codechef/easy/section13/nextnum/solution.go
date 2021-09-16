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
		S, _ := reader.ReadString('\n')

		res := solve(len(S)-1, S)
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

const N = 19

var P [N]int64

func init() {
	P[0] = 1
	for i := 1; i < N; i++ {
		P[i] = int64(i) * P[i-1]
	}
}

func countNumber(cnt []int64) int64 {
	var y int
	for i := 0; i < 10; i++ {
		y += int(cnt[i])
	}
	ans := P[y]
	for i := 0; i < 10; i++ {
		ans /= P[cnt[i]]
	}
	return ans
}

func solve(n int, S string) int64 {
	cnt := make([]int64, 10)
	for i := 0; i < n; i++ {
		cnt[int(S[i]-'0')]++
	}

	var res int64

	for i := 0; i < n; i++ {
		x := int(S[i] - '0')
		for y := 0; y < x; y++ {
			if cnt[y] > 0 {
				cnt[y]--
				res += countNumber(cnt)
				cnt[y]++
			}
		}
		cnt[x]--
	}

	return res + 1
}
