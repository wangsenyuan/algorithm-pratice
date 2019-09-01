package p1175

var primes []int
var fact []int

const MOD = 1e9 + 7

func init() {
	set := make([]bool, 201)

	for i := 2; i < 201; i++ {
		if set[i] {
			continue
		}
		primes = append(primes, i)
		for j := i * i; j < 201; j += i {
			set[j] = true
		}
	}
	fact = make([]int, 101)
	fact[0] = 1

	for i := 1; i <= 100; i++ {
		fact[i] = int((int64(i) * int64(fact[i-1])) % MOD)
	}
}

func numPrimeArrangements(n int) int {
	if n < 2 {
		return 1
	}

	//n positions
	i := search(len(primes), func(i int) bool {
		return primes[i] > n
	})

	// total i primes
	a := fact[i]
	b := fact[n-i]
	return int((int64(a) * int64(b)) % MOD)
}

func search(n int, fn func(int) bool) int {
	var left = 0
	var right = n

	for left < right {
		mid := (left + right) / 2
		if fn(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
