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

	n := readNum(reader)

	ask := func(on []int) int {
		var buf bytes.Buffer

		buf.WriteString(fmt.Sprintf("%d ", len(on)))
		for i := 0; i < len(on); i++ {
			buf.WriteString(fmt.Sprintf("%d ", on[i]))
		}
		fmt.Println(buf.String())
		return readNum(reader)
	}

	solve(n, ask)

	fmt.Println("0")
}

func solve(n int, ask func(on []int) int) {
	// find k that maximizes n - k - [n/k]

	var k = 1
	best := n - k - n/k
	for i := 2; i <= n; i++ {
		tmp := n - i - (n+i-1)/i
		if tmp > best {
			best = tmp
			k = i
		}
	}
	if k == 1 {
		// no matter what I do, it will be reverted
		return
	}
	lamps := make([]int, n)
	for i := 0; i < n; {
		lamps[i] = -1
		i += k
	}
	on := make([]int, k)
	for {
		var p int
		for i := 0; i < n && p < k; i++ {
			if lamps[i] == 0 {
				on[p] = i + 1
				p++
				lamps[i] = 1
			}
		}

		if p < k {
			break
		}

		x := ask(on)
		x--

		for i := 0; i < k; i++ {
			j := (x + i) % n
			if lamps[j] > 0 {
				lamps[j] = 0
			}
		}
	}
}
