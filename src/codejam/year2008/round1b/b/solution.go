package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	// f, err := os.Open("./A-small-practice.in")
	f, err := os.Open("./B-large-practice.in")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	// scanner := bufio.NewScanner(f)

	// tc := readNum(scanner)
	var tc int
	fmt.Fscanf(f, "%d", &tc)
	for i := 1; i <= tc; i++ {
		var A, B, P int64

		fmt.Fscanf(f, "%d %d %d", &A, &B, &P)
		// fmt.Printf("[debug]%d %d %d\n", A, B, P)
		fmt.Printf("Case #%d: %d\n", i, solve(A, B, P))
	}
}

const MAX_N = 1000007

var primes []int

func init() {
	set := make([]bool, MAX_N+1)
	primes = make([]int, 0, 100000)
	for i := 2; i <= MAX_N; i++ {
		if !set[i] {
			primes = append(primes, i)
			for j := i * i; j > i && j < MAX_N; j += i {
				set[j] = true
			}
		}
	}
}

func solve(A, B, P int64) int {
	n := int(B-A) + 1

	start := sort.Search(len(primes), func(i int) bool {
		return int64(primes[i]) >= P
	})

	end := sort.Search(len(primes), func(i int) bool {
		// if the prime number is large than the size of the interval,
		// there will be at most one multiple of the prime
		return primes[i] > n
	})

	uf := CreateUFSet(n)

	// find the multiple of factor at least A
	binarySearch := func(factor int64) int64 {
		left, right := int64(1), B/factor+1
		for left < right {
			mid := left + (right-left)>>1
			if factor*mid >= A {
				right = mid
			} else {
				left = mid + 1
			}
		}
		return left * factor
	}
	for i := start; i < end; i++ {
		prime := int64(primes[i])
		X := binarySearch(prime)

		for Y := X + prime; Y <= B; Y += prime {
			uf.Union(int(X-A), int(Y-A))
		}
	}
	return uf.size
}

type UFSet struct {
	set  []int
	cnt  []int
	size int
}

func CreateUFSet(n int) UFSet {
	set := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		cnt[i] = 1
		set[i] = i
	}
	return UFSet{set, cnt, n}
}

func (uf *UFSet) Union(a, b int) {
	var find func(u int) int
	find = func(u int) int {
		if uf.set[u] != u {
			uf.set[u] = find(uf.set[u])
		}
		return uf.set[u]
	}

	pa := find(a)
	pb := find(b)
	if pa == pb {
		return
	}
	if uf.cnt[pa] < uf.cnt[pb] {
		uf.set[pa] = pb
		uf.cnt[pb] += uf.cnt[pa]
	} else {
		uf.set[pb] = pa
		uf.cnt[pa] += uf.cnt[pb]
	}
	uf.size--
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func readFloat64(bs []byte, from int, val *float64) int {
	i := from
	sign := 1
	if bs[i] == '-' {
		sign = -1
		i++
	}

	var dec float64
	for i < len(bs) && bs[i] != '.' && bs[i] != ' ' {
		dec = dec*10 + float64(bs[i]-'0')
		i++
	}
	*val = dec

	if i == len(bs) || bs[i] == ' ' {
		//no fraction
		return i
	}
	i++
	var frac float64
	base := 1.0
	for i < len(bs) && bs[i] != ' ' {
		frac = frac*10 + float64(bs[i]-'0')
		base *= 10
		i++
	}
	*val += frac / base
	return i * sign
}

func readNFloat64s(scanner *bufio.Scanner, n int) []float64 {
	res := make([]float64, n)

	pos := 0
	scanner.Scan()
	bs := scanner.Bytes()
	//fmt.Printf("[debug] %s\n", string(bs))
	for i := 0; i < n; i++ {
		for pos < len(bs) && bs[pos] == ' ' {
			pos++
		}

		pos = readFloat64(bs, pos, &res[i])
	}
	return res
}
