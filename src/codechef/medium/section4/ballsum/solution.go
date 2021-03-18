package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, k := readTwoNums(reader)
	A := readNNums(reader, n)
	fmt.Printf("%d\n", solve(n, k, A))
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

func solve(n int, k int, A []int) int {
	cnt := make([]int, n+1)
	for i := 0; i < n; i++ {
		cnt[A[i]]++
	}
	pos := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		pos[i] = make([]int, 0, cnt[i])
	}

	for i := 0; i < n; i++ {
		pos[A[i]] = append(pos[A[i]], i)
	}

	var ans int

	for _, cur := range pos {
		if len(cur) == 0 {
			continue
		}
		ans += count(n, int64(k), cur)
		if ans >= MOD {
			ans -= MOD
		}
	}
	return ans
}

func count(n int, k int64, arr []int) int {
	freq := make([]int64, len(arr)+1)
	freq[0] = int64(arr[0] + 1)
	for i := 1; i < len(arr); i++ {
		freq[i] = int64(arr[i] - arr[i-1])
	}
	freq[len(arr)] = int64(n - arr[len(arr)-1])
	rf := make([]int64, len(freq))
	for i := 0; i < len(freq); i++ {
		rf[i] = freq[len(freq)-1-i]
	}
	res := fftMul(freq, rf)

	var ans int64

	for i := 1; i <= len(arr); i++ {
		tmp := pow(int64(i), k)
		tmp *= res[i+len(arr)]
		tmp %= MOD
		ans += tmp
		ans %= MOD
	}
	return int(ans)
}

func pow(a, b int64) int64 {
	var res int64 = 1
	for b > 0 {
		if b&1 == 1 {
			res *= a
			res %= MOD
		}
		a = a * a % MOD
		b >>= 1
	}
	return res
}

const MOD = 1000000007

const MAX_N = 200010

var f1 [MAX_N]complex128
var w [MAX_N]complex128
var R [MAX_N]int64
var arr [MAX_N]int64
var rarr [MAX_N]int64
var rev [MAX_N]int

func init() {
	for i := 0; i < MAX_N; i++ {
		w[i] = 0
	}
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
