package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	v1, v2, v3 := readThreeNums(reader)

	res := solve(v1, v2, v3)

	if len(res) == 0 {
		fmt.Println("No")
		return
	}
	var buf bytes.Buffer
	buf.WriteString("Yes\n")
	for i := 0; i < len(res); i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
	}
	buf.WriteByte('\n')

	fmt.Print(buf.String())
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

const V = 7 * 7 * 7

var pos []int

func init() {
	for i := -7; i <= 7; i++ {
		pos = append(pos, i)
	}
}

func solve(v1, v2, v3 int) []int {
	// 这个是只要在一个区域内
	sum := v1 + v2 + v3

	if v1+2*v2+3*v3 != 3*V || sum < V {
		return nil
	}

	get2 := func(first []int, second []int) int {
		res := 1
		for i := 0; i < 3; i++ {
			res *= max(0, min(first[i], second[i])+7-max(first[i], second[i]))
		}
		return res
	}

	get3 := func(first []int, second []int, third []int) int {
		res := 1
		for i := 0; i < 3; i++ {
			res *= max(0, min3(first[i], second[i], third[i])+7-max3(first[i], second[i], third[i]))
		}

		return res
	}

	first := []int{0, 0, 0}

	for _, a2 := range pos {
		for _, b2 := range pos {
			for _, c2 := range pos {
				second := []int{a2, b2, c2}
				for _, a3 := range pos {
					for _, b3 := range pos {
						for _, c3 := range pos {
							third := []int{a3, b3, c3}
							nv3 := get3(first, second, third)
							if nv3 != v3 {
								continue
							}
							nv2 := get2(first, second) + get2(first, third) + get2(second, third) - 3*nv3
							if nv2 != v2 {
								continue
							}
							return []int{0, 0, 0, a2, b2, c2, a3, b3, c3}
						}
					}
				}
			}
		}
	}

	return nil
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func max3(a, b, c int) int {
	return max(a, max(b, c))
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func min3(a, b, c int) int {
	return min(a, min(b, c))
}
