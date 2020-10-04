package p1544

// si = s(i - 1) + '1' + reverse(invert(s(i - 1)))
// len(s[i]) = 2 * len(s[i-1]) + 1
const N = 21

var length [N]int

func init() {
	length[1] = 1
	for i := 2; i < N; i++ {
		length[i] = 2*length[i-1] + 1
	}
}

func findKthBit(n int, k int) byte {
	if k == 1 || n == 1 {
		return '0'
	}

	l := length[n]
	h := l / 2

	if h+1 == k {
		return '1'
	}

	if h+1 > k {
		// in the first part
		return findKthBit(n-1, k)
	}
	// if only reverse
	r := findKthBit(n-1, length[n-1]+1-(k-h-1))
	if r == '0' {
		return '1'
	}
	return '0'
}
