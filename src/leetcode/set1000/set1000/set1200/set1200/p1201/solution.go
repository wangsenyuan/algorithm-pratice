package p1201

var primes []int

func init() {
	N := 200000
	set := make([]bool, N+1)

	for x := 2; x < N; x++ {
		if set[x] {
			continue
		}
		primes = append(primes, x)
		y := x * x
		if y < 0 {
			// overflow
			continue
		}
		for y <= N {
			set[y] = true
			y += x
		}
	}
}

func nthUglyNumber(n int, a int, b int, c int) int {

	// x = a * i
	// or x = b * j
	// or x = c * k
	A := int64(a)
	B := int64(b)
	C := int64(c)
	ab := findLCM([]int{a, b})
	bc := findLCM([]int{b, c})
	ac := findLCM([]int{a, c})
	abc := findLCM([]int{a, b, c})

	count := func(x int64) int64 {
		sum := x/A + x/B + x/C

		sum -= x / ab
		sum -= x / ac
		sum -= x / bc
		sum += x / abc

		return sum
	}

	N := int64(n)

	left := 1
	right := 2000000020

	for left < right {
		mid := (left + right) / 2
		x := count(int64(mid))
		if x >= N {
			right = mid
		} else {
			left = mid + 1
		}
	}
	// count(right) >= N
	x := right
	for {
		if x%a == 0 {
			return x
		}
		if x%b == 0 {
			return x
		}
		if x%c == 0 {
			return x
		}
		x++
	}
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func findLCM(arr []int) int64 {
	n := len(primes)
	cnt := make([]int, n)

	for i := 0; i < len(arr); i++ {
		for j := 0; j < n && arr[i] >= primes[j]; j++ {
			var c int
			for arr[i]%primes[j] == 0 {
				c++
				arr[i] /= primes[j]
			}
			if c > cnt[j] {
				cnt[j] = c
			}
		}
	}

	res := int64(1)

	for i := 0; i < len(arr); i++ {
		res *= int64(arr[i])
	}

	for j := 0; j < n; j++ {
		if cnt[j] > 0 {
			res *= pow(primes[j], cnt[j])
		}
	}

	return res
}

func pow(a int, b int) int64 {
	x := int64(a)
	r := int64(1)
	for b > 0 {
		if b&1 == 1 {
			r *= x
		}
		x *= x
		b >>= 1
	}
	return r
}
