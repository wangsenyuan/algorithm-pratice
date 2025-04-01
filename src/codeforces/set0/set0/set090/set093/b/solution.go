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
	n, w, m := readThreeNums(reader)
	res := solve(n, w, m)

	if len(res) == 0 {
		fmt.Println("NO")
		return
	}
	var buf bytes.Buffer
	buf.WriteString("YES\n")
	for _, cur := range res {
		for _, x := range cur {
			buf.WriteString(fmt.Sprintf("%d %.8f ", x.id, x.vol))
		}
		buf.WriteByte('\n')
	}
	buf.WriteTo(os.Stdout)
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

type state struct {
	id  int
	vol float64
}

func solve(n int, w int, m int) [][]state {
	// f = n * w / m
	// f < w / 2
	if 2*n < m {
		return nil
	}
	// f >= w / 2
	var res [][]state

	avg := float64(n*w) / float64(m)
	var j int
	var cur float64

	wf := float64(w)

	for range m {
		var tmp []state
		var more float64
		if cur > avg && math.Abs(cur-avg) > eps {
			return nil
		}
		if cur > eps {
			tmp = append(tmp, state{j + 1, cur})
			j++
			more = avg - cur
			cur = 0
		} else {
			more = avg
		}
		for more > eps {
			if check(wf, more) {
				tmp = append(tmp, state{j + 1, wf})
				j++
				more -= wf
			} else {
				tmp = append(tmp, state{j + 1, more})
				cur = wf - more
				break
			}
		}
		res = append(res, tmp)
	}

	return res
}

const eps = 1e-9

func check(a float64, b float64) bool {
	return a <= b || math.Abs(a-b) < eps
}
