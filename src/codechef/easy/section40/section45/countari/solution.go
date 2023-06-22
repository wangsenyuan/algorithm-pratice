package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	A := readNNums(reader, n)
	res := solve(A)

	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
}

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')
	if len(s) == 0 || len(s) == 1 && s[0] == '\n' {
		return readNInt64s(reader, n)
	}
	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
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

func solve(A []int) int64 {
	// A[i] <= 30000
	// n <= 100000
	// A[j] - A[i] = A[k] - A[j] => A[j] * 2 = A[i] + A[k]
	n := len(A)
	for i := 0; i < n; i++ {
		A[i]--
	}
	bef := make([]int, M)
	aft := make([]int, M)
	for _, a := range A {
		aft[a]++
	}

	per := (n + BLOCK_SIZE - 1) / BLOCK_SIZE

	p1 := make([]complex128, N2)
	p2 := make([]complex128, N2)

	inside := make([]int, M)

	var res int64

	for block := 0; block < BLOCK_SIZE; block++ {
		st := block * per
		ed := min((block+1)*per, n)

		for i := st; i < ed; i++ {
			aft[A[i]]--
		}

		for i := 0; i < M; i++ {
			inside[i] = 0
		}

		for i := st; i < ed; i++ {
			for j := i + 1; j < ed; j++ {
				if A[i] != A[j] {
					k := A[i] - (A[j] - A[i])
					if 0 <= k && k < M {
						res += int64(bef[k] + inside[k])
					}
					k = A[j] + A[j] - A[i]
					if 0 <= k && k < M {
						res += int64(aft[k])
					}
				}
			}
			inside[A[i]]++
		}

		for i := 0; i < M; i++ {
			// 中间两个在inside里面，另外一个在前面或者后面
			res += int64(inside[i]) * int64(inside[i]-1) / 2 * int64(bef[i]+aft[i])
			// 都在inside
			res += int64(inside[i]) * int64(inside[i]-1) * int64(inside[i]-2) / 6
		}

		if block > 0 && block+1 < BLOCK_SIZE {
			for i := 0; i < N2; i++ {
				p1[i] = complex(0, 0)
				p2[i] = complex(0, 0)
				if i < M {
					p1[i] = complex(float64(bef[i]), 0)
					p2[i] = complex(float64(aft[i]), 0)
				}
			}

			fft(N2, 2*math.Pi*INV_N2, p1)
			fft(N2, 2*math.Pi*INV_N2, p2)
			for i := 0; i < N2; i++ {
				p1[i] *= p2[i]
			}
			fft(N2, -2*math.Pi*INV_N2, p1)
			for i := 0; i < N2; i++ {
				p1[i] = complex(real(p1[i])*INV_N2, imag(p1[i]))
			}

			for i := 0; i < M; i++ {
				res += int64(inside[i]) * int64(real(p1[2*i])+0.5)
			}
		}

		for i := st; i < ed; i++ {
			bef[A[i]]++
		}
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

const M = 30000
const BLOCK_SIZE = 30

const N2 = 65536
const INV_N2 = 1.0 / N2

func fft(n int, theta float64, a []complex128) {
	for m := n; m > 1; m /= 2 {
		half := m / 2
		for i := 0; i < half; i++ {
			w := polar(1, float64(i)*theta)
			for j := i; j < n; j += m {
				k := j + half
				x := a[j] - a[k]
				a[j] += a[k]
				a[k] = w * x
			}
		}
		theta *= 2
	}

	for i, j := 0, 1; j < n-1; j++ {
		for k := n / 2; ; k /= 2 {
			i ^= k
			if k <= i {
				break
			}
		}
		if j < i {
			a[i], a[j] = a[j], a[i]
		}
	}
}

func polar(r float64, t float64) complex128 {
	return complex(r*math.Cos(t), r*math.Sin(t))
}
