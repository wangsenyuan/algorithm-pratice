package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"time"
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
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

	ask := func(l, r int) int {
		fmt.Printf("1 %d %d\n", l, r)
		return readNum(reader)
	}

	res := solve(ask)

	var buf bytes.Buffer
	buf.WriteString("2")
	for i := 0; i < len(res); i++ {
		buf.WriteString(fmt.Sprintf(" %d", res[i]))
	}
	fmt.Println(buf.String())
}

const N = 100000

var arr [N + 1]int

func solve(ask func(l, r int) int) []int {
	s1 := rand.NewSource(time.Now().UnixNano())
	rd := rand.New(s1)

	for i := 1; i < N; i += 3 {
		ans := ask(i, i+2)
		if ans == 0 {
			continue
		}
		if ans == 3 {
			arr[i] = 1
			arr[i+1] = 1
			arr[i+2] = 1
			continue
		}
		if ans == 2 {
			// 0, 1, 2
			j := i + rd.Intn(3)
			x := ask(j, j)
			// x == 0 or x == 1
			if x == 0 {
				for k := i; k < i+3; k++ {
					if k != j {
						arr[k] = 1
					}
				}
				continue
			} else {
				arr[j] = 1
				jj := j + 1
				if jj == i+3 {
					jj = j - 1
				}
				y := ask(jj, jj)
				if y == 1 {
					arr[jj] = 1
				} else {
					arr[i^(i+1)^(i+2)^j^jj] = 1
				}
			}
			continue
		}
		// ans == 1
		j := i + rd.Intn(3)
		x := ask(j, j)
		if x == 1 {
			arr[j] = 1
			continue
		}
		jj := j + 1
		if jj == i+3 {
			jj = j - 1
		}
		x = ask(jj, jj)
		if x == 1 {
			arr[jj] = 1
			continue
		}
		arr[i^(i+1)^(i+2)^j^jj] = 1
	}

	arr[N] = ask(N, N)

	return arr[1:]
}
