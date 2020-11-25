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
		res := solve(n, A)
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

const H = 50

func solve(n int, A []int) uint64 {
	cnt := make([]uint64, H)

	for i := 0; i < n; i++ {
		cur := A[i]
		for j := 0; j < H && cur > 0; j++ {
			cnt[j] += uint64(cur & 1)
			cur >>= 1
		}
	}
	var res uint64

	for i := 0; i < H; i++ {
		if i > 0 {
			cnt[i] += cnt[i-1] / 2
		}
		if cnt[i] > 0 {
			res |= 1 << uint64(i)
		}
	}

	return res
}

func solve1(n int, A []int) uint64 {
	cnt := make([]uint64, H)

	for i := 0; i < n; i++ {
		cur := A[i]
		for j := 0; j < H && cur > 0; j++ {
			cnt[j] += uint64(cur & 1)
			cur >>= 1
		}
	}
	var res uint64

	for j := H - 1; j >= 0; j-- {
		if cnt[j] > 0 {
			res |= 1 << uint64(j)
			continue
		}
		if j == 0 {
			continue
		}
		var ok bool
		var borrow uint64 = 1
		for k := 1; j-k >= 0; k++ {
			borrow <<= 1
			if cnt[j-k] >= borrow {
				ok = true
				break
			}
			borrow -= cnt[j-k]
		}
		if ok {
			res |= 1 << uint64(j)
		}
	}
	return res
}
