package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
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

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)

		res := solve(n, A)

		for i := 0; i < len(res); i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
}

const MAX_N = 1 << 21

var f1 [MAX_N]complex128
var w [MAX_N]complex128
var R [MAX_N]int64
var arr [MAX_N]int64
var rarr [MAX_N]int64
var rev [MAX_N]int

const LIMIT = 3000

func init() {
	for i := 0; i < MAX_N; i++ {
		w[i] = 0
	}
}

func solve(n int, A []int) []int64 {
	cnt := make(map[int]int)

	var pre int

	ans := make([]int64, n+1)

	for i := 0; i < n; i++ {
		pre ^= A[i]
		if pre == 0 {
			ans[i+1]++
		}
		cnt[pre]++
	}
	mp := make(map[int][]int)

	pre = 0

	for i := 0; i < n; i++ {
		pre ^= A[i]
		if mp[pre] == nil {
			mp[pre] = make([]int, 0, cnt[pre])
		}
		mp[pre] = append(mp[pre], i)
	}

	for _, v := range mp {
		if len(v) < LIMIT {
			for i := 0; i < len(v); i++ {
				for j := i + 1; j < len(v); j++ {
					ans[v[j]-v[i]]++
				}
			}
		} else {
			for i := 0; i < n; i++ {
				arr[i] = 0
				rarr[i] = 0
			}
			for _, j := range v {
				arr[j] = 1
				rarr[n-1-j] = 1
			}
			//reverse(rev[:n])

			res := fftMul(arr[:n], rarr[:n])

			for i := 0; i < n; i++ {
				ans[n-1-i] += res[i]
			}
		}
	}

	return ans[1:]
}

func fftMul(a, b []int64) []int64 {
	var k = 1
	for (1 << uint(k)) < len(a)+len(b) {
		k++
	}

	buildRev(k)

	n := 1 << uint(k)
	for i := 0; i < n; i++ {
		f1[i] = 0
	}
	for i := 0; i < len(a); i++ {
		f1[i] += complex(float64(a[i]), 0)
	}
	for i := 0; i < len(b); i++ {
		f1[i] += complex(0, float64(b[i]))
	}

	fft(f1[:], k)

	for i := 0; i < 1+n/2; i++ {
		p := f1[i] + conj(f1[(n-i)%n])
		_q := f1[(n-i)%n] - conj(f1[i])
		q := complex(imag(_q), real(_q))
		f1[i] = p * q * 0.25
		if i > 0 {
			f1[n-i] = conj(f1[i])
		}
	}

	for i := 0; i < n; i++ {
		f1[i] = conj(f1[i])
	}

	fft(f1[:], k)

	for i := 0; i < len(a)+len(b); i++ {
		R[i] = int64(real(f1[i])/float64(n) + 0.5)
	}

	return R[:len(a)+len(b)]
}

func buildRev(k int) {
	K := 1 << uint(k)
	for i := 1; i <= K; i++ {
		j := rev[i-1]
		t := k - 1
		for t >= 0 && ((j>>uint(t))&1) == 1 {
			j ^= 1 << uint(t)
			t--
		}
		if t >= 0 {
			j ^= 1 << uint(t)
			t--
		}
		rev[i] = j
	}
}

func fft(a []complex128, k int) {
	n := 1 << uint(k)

	for i := 0; i < n; i++ {
		if rev[i] > i {
			a[i], a[rev[i]] = a[rev[i]], a[i]
		}
	}
	for l, l1 := 2, 1; l <= n; l, l1 = l*2, 2*l1 {
		if real(w[l1]) == 0 && imag(w[l1]) == 0 {
			if l1 > 1 {
				angle := math.Pi / float64(l1)
				ww := complex(math.Cos(angle), math.Sin(angle))
				for j := 0; j < l1; j++ {
					if j&1 == 1 {
						w[l1+j] = w[(l1+j)/2] * ww
					} else {
						w[l1+j] = w[(l1+j)/2]
					}
				}
			} else {
				w[l1] = 1
			}
		}
		for i := 0; i < n; i += l {
			for j := 0; j < l1; j++ {
				v := a[i+j]
				u := a[i+j+l1] * w[l1+j]
				a[i+j] = v + u
				a[i+j+l1] = v - u
			}
		}
	}
}

func conj(cur complex128) complex128 {
	return complex(real(cur), -imag(cur))
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
