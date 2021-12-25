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
		ans := solve(n)

		buf.WriteString(fmt.Sprintf("%d\n", len(ans.plates)))
		for i := 0; i < len(ans.plates); i++ {
			buf.WriteString(fmt.Sprintf("%d ", ans.plates[i]))
		}
		buf.WriteByte('\n')
		for i := 0; i < n; i++ {
			cur := ans.dayPuts[i]
			buf.WriteString(fmt.Sprintf("%d\n", len(cur)))
			for j := 0; j < len(cur); j++ {
				buf.WriteString(fmt.Sprintf("%d %d\n", cur[j].fruit, cur[j].plate))
			}
		}
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

func solve(n int) Answer {
	plates := make([]int, 0, 20)

	for i := 0; ; {
		sz := 1 << uint(i)
		i++
		if sz >= n {
			plates = append(plates, n)
			break
		}
		plates = append(plates, sz)
	}

	dayPuts := make([][]Put, n)
	if n == 1 {
		dayPuts[0] = append(dayPuts[0], Put{1, 1})
	} else {
		cur := 1
		elem := 0
		filled := make([]int, len(plates))

		for i := 1; i <= n; i++ {
			if i == 1 {
				dayPuts[i-1] = append(dayPuts[i-1], Put{1, 1})
				filled[0]++
				dayPuts[i-1] = append(dayPuts[i-1], Put{2, 1})
				filled[1]++
				continue
			}
			if cur == len(plates)-1 {
				dayPuts[i-1] = append(dayPuts[i-1], Put{cur + 1, i})
			} else {
				dayPuts[i-1] = append(dayPuts[i-1], Put{cur + 1, i})
				filled[cur]++
				elem++
				dayPuts[i-1] = append(dayPuts[i-1], Put{cur + 2, elem})
				elem++
				dayPuts[i-1] = append(dayPuts[i-1], Put{cur + 2, elem})
				filled[cur+1] += 2
				if filled[cur] == plates[cur] {
					cur++
					elem = 0
				}
			}
		}

	}

	return Answer{plates, dayPuts}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type Answer struct {
	plates  []int
	dayPuts [][]Put
}

type Put struct {
	plate int
	fruit int
}
