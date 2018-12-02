package p952

const MAX_NUM = 100010

var primes []int
var smf []int
var set []bool

func init() {
	set = make([]bool, MAX_NUM)
	smf = make([]int, MAX_NUM)

	for i := 0; i < MAX_NUM; i++ {
		smf[i] = -1
	}

	primes = make([]int, MAX_NUM)
	var j int
	for num := 2; num < MAX_NUM; num++ {
		if !set[num] {
			primes[j] = num
			smf[num] = j
			j++
			for x := num * num; x < MAX_NUM && x > 0; x += num {
				if smf[x] < 0 {
					smf[x] = j - 1
				}
				set[x] = true
			}
		}
	}
	primes = primes[:j]
}

func largestComponentSize(A []int) int {
	n := len(A)
	xs := make([][]int, n)
	ys := make([]map[int]bool, len(primes))
	for i := 0; i < len(primes); i++ {
		ys[i] = make(map[int]bool)
	}
	for i := 0; i < len(A); i++ {
		xs[i] = make([]int, 0, 10)
		num := A[i]
		fs := make(map[int]bool)
		for num > 1 {
			j := smf[num]
			f := primes[j]
			if !fs[f] {
				xs[i] = append(xs[i], j)
				ys[j][i] = true
			}
			fs[f] = true
			num /= f
		}
	}
	uf := NewUF(len(primes))
	for i := 0; i < n; i++ {
		x := xs[i]
		for j := 0; j < len(x); j++ {
			for k := j + 1; k < len(x); k++ {
				uf.Union(x[j], x[k])
			}
		}
	}

	var ans int

	for i := 0; i < len(primes); i++ {
		if len(ys[i]) > 0 {
			// number have factor primes[i]
			pi := uf.Find(i)
			ys[pi] = merge(ys[pi], ys[i])
			if len(ys[pi]) > ans {
				ans = len(ys[pi])
			}
		}
	}
	return ans
}

func merge(a, b map[int]bool) map[int]bool {
	if len(a) < len(b) {
		a, b = b, a
	}
	for k := range b {
		a[k] = true
	}
	return a
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type UF struct {
	set  []int
	size []int
}

func NewUF(n int) *UF {
	uf := new(UF)
	uf.set = make([]int, n)
	uf.size = make([]int, n)
	for i := 0; i < n; i++ {
		uf.set[i] = i
		uf.size[i] = 1
	}
	return uf
}

func (uf *UF) Find(x int) int {
	var loop func(x int) int
	loop = func(x int) int {
		if uf.set[x] != x {
			uf.set[x] = loop(uf.set[x])
		}
		return uf.set[x]
	}
	return loop(x)
}

func (uf *UF) Union(a, b int) bool {
	pa := uf.Find(a)
	pb := uf.Find(b)
	if pa == pb {
		return false
	}
	if uf.size[pa] >= uf.size[pb] {
		uf.set[pb] = pa
		uf.size[pa] += uf.size[pb]
	} else {
		uf.set[pa] = pb
		uf.size[pb] += uf.size[pa]
	}
	return true
}

func (uf *UF) Keys() []int {
	res := make([]int, len(uf.set))
	var j int
	for i := 0; i < len(uf.set); i++ {
		if uf.set[i] == i {
			res[j] = i
			j++
		}
	}
	return res[:j]
}
