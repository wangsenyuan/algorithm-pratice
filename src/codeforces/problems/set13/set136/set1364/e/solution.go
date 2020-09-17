package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
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

	ask := func(i, j int) int {
		fmt.Printf("? %d %d\n", i, j)
		return readNum(reader)
	}

	res := solve(n, ask)

	var buf bytes.Buffer
	buf.WriteByte('!')
	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf(" %d", res[i]))
	}
	fmt.Println(buf.String())
}

func solve(n int, ask func(int, int) int) []int {
	r := rand.Perm(n)
	p := make([]int, n)

	for i, j := range r {
		p[j] = i + 1
	}

	z := make([]int, 11)

	for i := 0; i < len(z); i++ {
		z[i] = -1
	}

	for !allSet(z) {
		i := rand.Intn(n) + 1
		j := rand.Intn(n) + 1
		if i == j {
			continue
		}
		tmp := ask(i, j)
		for b := 0; b < 11; b++ {
			if tmp&(1<<b) == 0 {
				z[b] = i
			}
		}
	}

	get := func(i int, sup int) int {
		ans := sup

		for b := 0; b < 11; b++ {
			if ans&(1<<b) > 0 {
				if z[b] != i {
					ans &= ask(i, z[b])
				} else {
					ans ^= 1 << b
				}
			}
		}
		return ans
	}

	idx := -1
	val := 1<<11 - 1

	for _, i := range p {
		if idx < 0 || ask(i, idx) == val {
			idx = i
			val = get(i, val)
		}
	}
	// val == 0
	for i := 1; i <= n; i++ {
		if i == idx {
			p[i-1] = 0
		} else {
			p[i-1] = ask(i, idx)
		}
	}

	return p
}

func allSet(arr []int) bool {
	var res int
	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 {
			res++
		}
	}
	return res == len(arr)
}

func solve1(n int, ask func(int, int) int) []int {
	r := rand.Perm(n)
	p := make([]int, n)

	for i, j := range r {
		p[j] = i + 1
	}

	a, b := p[0], p[1]
	val := ask(a, b)

	for i := 2; i < n; i++ {
		c := p[i]
		tmp := ask(b, c)
		if tmp < val {
			a = c
			val = tmp
		} else if tmp == val {
			b = c
			val = ask(a, b)
		}
	}
	idx := -1
	//a, b are candidates for 0
	for idx < 0 {
		i := rand.Intn(n) + 1
		if i == a || i == b {
			continue
		}
		x := ask(i, a)
		y := ask(i, b)
		if x == y {
			continue
		}
		if x < y {
			idx = a
		} else {
			idx = b
		}
	}

	p[idx-1] = 0

	for i := 1; i <= n; i++ {
		if i == idx {
			continue
		}
		p[i-1] = ask(i, idx)
	}

	return p
}
