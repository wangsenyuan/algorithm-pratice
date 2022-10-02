package p2430

func deleteString(s string) int {
	n := len(s)

	dp := make([]int, n+1)
	dp[n] = 1

	keys := make([][]Key, n+1)

	for i := n - 1; i >= 0; i-- {
		keys[i] = make([]Key, 1+(n+1)/2)
		keys[i][0] = NewKey()
		for l := 1; l < len(keys[i]); l++ {
			j := i + l - 1
			if j >= n {
				break
			}
			keys[i][l] = keys[i][l-1].Add(uint64(s[j] - 'a' + 1))
		}
		dp[i] = 1
		for l := 1; l < len(keys[i]); l++ {
			j := i + l
			if j+l > n {
				break
			}
			if keys[i][l] == keys[j][l] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
	}

	return dp[0]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

const M1 = 1000000007
const M2 = 10000000009

type Key struct {
	first  uint64
	second uint64
}

func NewKey() Key {
	return Key{0, 0}
}

func (this Key) Add(num uint64) Key {
	first := (this.first*27 + num) % M1
	second := (this.second*27 + num) % M2
	return Key{first, second}
}
