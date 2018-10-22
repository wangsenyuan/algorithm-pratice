package p927

func threeEqualParts(A []int) []int {
	n := len(A)
	var ones int

	for i := 0; i < n; i++ {
		ones += A[i]
	}
	if ones == 0 {
		return []int{0, 2}
	}

	if ones%3 != 0 {
		return []int{-1, -1}
	}

	l := ones / 3

	k := n - 1
	for l > 0 {
		l -= A[k]
		k--
	}
	k = n - 1 - k
	//A[k...n) is the pattern after one [1...]
	var i int
	for A[i] == 0 {
		i++
	}
	// j is the start position of the second part
	first := i + k

	j := first

	for j < n && A[j] == 0 {
		j++
	}
	second := j + k

	cmp := func(i, j, k int) bool {
		for k < n {
			if A[i] != A[j] {
				return false
			}
			if A[i] != A[k] {
				return false
			}
			if A[j] != A[k] {
				return false
			}
			i++
			j++
			k++
		}
		return k == n
	}

	if !cmp(i, j, n-k) {
		return []int{-1, -1}
	}
	return []int{first - 1, second}
}

func threeEqualParts2(A []int) []int {
	n := len(A)

	var i int
	for i < n && A[i] == 0 {
		i++
	}

	if i == n {
		return []int{0, 2}
	}

	cmp := func(i, j, k, l int) bool {
		for a := 0; a < l; a++ {
			if A[i+a] != A[j+a] {
				return false
			}
			if A[i+a] != A[k+a] {
				return false
			}

			if A[j+a] != A[k+a] {
				return false
			}
		}
		return true
	}

	for k := 1; k <= (n-i)/3; k++ {
		j := i + k
		for j < n && A[j] == 0 {
			j++
		}
		if j == n {
			break
		}

		x := j + k
		for x < n && A[x] == 0 {
			x++
		}

		if x+k == n {
			if cmp(i, j, x, k) {
				return []int{i + k - 1, j + k}
			}
		}

	}

	return []int{-1, -1}
}

const MOD = 1000000007

func threeEqualParts1(A []int) []int {
	f := make(map[int64]int)
	n := len(A)
	var prod int64
	for i := 0; i < n; i++ {
		prod = prod*2 + int64(A[i])
		prod %= MOD
		if _, found := f[prod]; !found {
			f[prod] = i
		}
	}
	prod = 0
	factor := int64(1)
	for i := n - 1; i >= 0; i-- {
		prod = factor*int64(A[i]) + prod
		prod %= MOD
		factor = (factor * 2) % MOD
		if j, found := f[prod]; found {
			// end at j
			//the need to check elements [j+1, i)
			k := j + 1
			var tmp int64
			for k < i {
				tmp = (tmp*2 + int64(A[k])) % MOD
				k++
			}
			if tmp == prod {
				return []int{j, i}
			}
		}
	}
	return []int{-1, -1}
}
