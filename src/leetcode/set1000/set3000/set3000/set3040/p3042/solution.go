package p3042

func countPrefixSuffixPairs(words []string) int64 {

	//slices.SortFunc(words, func(a, b string) int {
	//	return len(a) - len(b)
	//})
	//n := len(words)
	var ln int
	for _, w := range words {
		ln = max(ln, len(w))
	}
	ks := make([]Key, ln+1)
	var res int64

	freq := make(map[Key]int64)

	lps := make([]int, ln)

	process := func(w string) {
		for i := 1; i < len(w); i++ {
			j := lps[i-1]
			for j > 0 && w[i] != w[j] {
				j = lps[j-1]
			}
			if w[i] == w[j] {
				j++
			}
			lps[i] = j
		}

		var key Key

		for i := 0; i < len(w); i++ {
			v := uint64(w[i] - 'a')
			key = key.Mul(BASE).AddValue(v + 1)
			ks[i] = key
		}

		j := lps[len(w)-1]

		for j > 0 {
			res += freq[ks[j-1]]
			j = lps[j-1]
		}

		res += freq[key]

		freq[key]++
	}

	for _, w := range words {
		process(w)
	}

	return res
}

const BASE = 27

const MOD1 = 100000000007
const MOD2 = 100000000009

type Key struct {
	first  uint64
	second uint64
}

func NewKey(v uint64) Key {
	return Key{v % MOD1, v % MOD2}
}

func (this Key) Add(that Key) Key {
	return Key{(this.first + that.first) % MOD1, (this.second + that.second) % MOD2}
}

func (this Key) Sub(that Key) Key {
	return this.Add(Key{MOD1 - that.first, MOD2 - that.second})
}

func (this Key) Mul(v uint64) Key {
	return Key{(this.first * v) % MOD1, (this.second * v) % MOD2}
}

func (this Key) AddValue(v uint64) Key {
	return Key{(this.first + v) % MOD1, (this.second + v) % MOD2}
}
