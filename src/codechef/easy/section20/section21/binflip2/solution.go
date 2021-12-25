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
		S, _ := reader.ReadString('\n')
		steps, first := solve(len(S)-1, S)
		if first < 0 {
			buf.WriteString("NO\n")
		} else {
			buf.WriteString("YES\n")
			buf.WriteString(fmt.Sprintf("%d\n", len(steps)))
			if len(steps) > 0 {
				buf.WriteString(fmt.Sprintf("%d\n", first))
				for _, step := range steps {
					buf.WriteString(fmt.Sprintf("%d\n", step))
				}
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

func solve(n int, S string) ([]int, int) {
	S = S[:n]
	c := count(S, '1')
	if c == 0 {
		return nil, 0
	}

	if c == n && n%2 == 0 {
		return nil, -1
	}

	var steps []int
	cur := 1 + (c % 2)
	buf := []byte(S)

	flip := func(p int) {
		if buf[p] == '1' {
			buf[p] = '0'
		} else {
			buf[p] = '1'
		}
	}
	// even numbers start by op 1, odd by 2
	var pos int
	for pos < n {
		if cur == 2 {
			i := find(buf[pos:], '1') + pos
			if i == n {
				break
			}
			buf[i] = '0'
			steps = append(steps, i)
			cur ^= 3
		}
		x := find(buf[pos:], '0') + pos
		if x == n {
			flip(pos - 1)
			flip(pos)
			flip(pos + 1)
			steps = append(steps, pos-1)
			flip(pos - 1)
			steps = append(steps, pos-1)
			pos += 2
			continue
		}
		for i := get(x, pos); i >= pos; i -= 2 {
			flip(i)
			flip(i + 1)
			flip(i + 2)
			steps = append(steps, i)
			if buf[i+1] == '1' {
				flip(i + 1)
				steps = append(steps, i+1)
			} else {
				flip(i + 2)
				steps = append(steps, i+2)
			}
		}
		pos = x + 1
	}

	return steps, 1 + (c % 2)
}

func get(x int, pos int) int {
	if (x-pos)%2 == 1 {
		return x - 1
	}
	return x - 2
}

func count(S string, x byte) int {
	var res int
	for i := 0; i < len(S); i++ {
		if S[i] == x {
			res++
		}
	}
	return res
}

func find(S []byte, x byte) int {
	for i := 0; i < len(S); i++ {
		if S[i] == x {
			return i
		}
	}
	return len(S)
}
