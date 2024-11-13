package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		res := process(reader)
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

func process(reader *bufio.Reader) int {
	n, x := readTwoNums(reader)
	blows := make([][]int, n)
	for i := range n {
		blows[i] = readNNums(reader, 2)
	}
	return solve(x, blows)
}

func solve(x int, blows [][]int) int {
	var damage int

	for _, cur := range blows {
		damage = max(damage, cur[0])
	}
	// 最后一次用damage

	if damage >= x {
		// 最需要一次
		return 1
	}

	var best = -1

	for _, cur := range blows {
		d, h := cur[0], cur[1]
		if d <= h {
			// no usage
			continue
		}
		per := d - h
		// x - n * per <= damage
		// x - damage <= n * per
		n := (x - damage) / per
		for x-n*per > damage {
			n++
		}
		if best < 0 || n+1 < best {
			best = n + 1
		}
	}

	return best
}
