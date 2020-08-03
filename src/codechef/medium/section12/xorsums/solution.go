package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	A := readNNums(reader, n)

	fmt.Println(solve(n, A))
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

func solve(n int, A []int) int {
	var X int
	for i := 0; i < n; i++ {
		X += A[i]
	}

	arr := make([]int, 2*X+3)

	update := func(p int, v int) {
		p += X
		arr[p] += v
		for p > 1 {
			arr[p>>1] = arr[p] + arr[p^1]
			p >>= 1
		}
	}

	query := func(l, r int) int {
		l += X
		r += X

		var res int

		for l < r {
			if l&1 == 1 {
				res += arr[l]
				l++
			}
			if r&1 == 1 {
				r--
				res += arr[r]
			}
			l >>= 1
			r >>= 1
		}
		return res
	}
	var res int
	S := make([]int, n+1)
	for b := 0; 1<<uint(b) <= X; b++ {
		for i := 0; i < n; i++ {
			S[i+1] = (S[i] + A[i]) % (1 << uint(b+1))
		}

		var count int
		update(S[0], 1)

		for i := 1; i <= n; i++ {
			l, r := S[i]+1, S[i]+(1<<uint(b))
			if r >= X {
				r = X
			}
			count += query(l, r+1)
			r = S[i] - (1 << uint(b))
			if r >= 0 {
				count += query(0, r+1)
			}
			update(S[i], 1)
		}

		if count&1 == 1 {
			res += 1 << uint(b)
		}

		for i := 0; i <= n; i++ {
			update(S[i], -1)
		}
	}

	return res
}
