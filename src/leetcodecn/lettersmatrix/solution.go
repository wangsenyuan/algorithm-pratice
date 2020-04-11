package lettersmatrix

import "sort"

const B = 27
const MOD = 10000000007

func maxRectangle(words []string) []string {
	n := len(words)

	pref := make([][]bool, 26)
	for i := 0; i < 26; i++ {
		pref[i] = make([]bool, 26)
	}

	for _, word := range words {
		for i := 1; i < len(word); i++ {
			a, b := int(word[i-1]-'a'), int(word[i]-'a')
			pref[b][a] = true
		}
	}

	wordsByLen := make(map[int][]int)

	for i, word := range words {
		m := len(word)
		if _, found := wordsByLen[m]; !found {
			wordsByLen[m] = make([]int, 0, 10)
		}
		wordsByLen[m] = append(wordsByLen[m], i)
	}

	checkCanPutOnTop := func(x, y string) bool {
		for i := 0; i < len(y); i++ {
			a := int(x[i] - 'a')
			b := int(y[i] - 'a')
			if !pref[b][a] {
				return false
			}
		}
		return true
	}

	conn := make([][]int, n)

	for i := 0; i < n; i++ {
		conn[i] = make([]int, 0, 10)
	}

	lens := make([]int, 0, len(wordsByLen))

	for k, ws := range wordsByLen {
		for i := 0; i < len(ws); i++ {
			for j := 0; j < len(ws); j++ {
				if checkCanPutOnTop(words[ws[i]], words[ws[j]]) {
					conn[ws[i]] = append(conn[ws[i]], ws[j])
				}
			}
		}
		lens = append(lens, k)
	}

	mem := make(map[int64]bool)
	mem1 := make(map[int64]bool)
	mem2 := make(map[int64]int)
	for _, word := range words {
		var res int64

		for i := 0; i < len(word); i++ {
			res = (res*B + int64(word[i])) % MOD
			mem[res] = true
			mem2[res] = max(mem2[res], len(word))
		}
		mem1[res] = true
	}

	rows := make([]int, 100)
	colHash := make([]int64, 100)

	var extendRow func(u, r int, n int)

	var ans Answer

	createAnswer := func(r, c int) Answer {
		ws := make([]string, r)
		for i := 0; i < r; i++ {
			ws[i] = words[rows[i]]
		}
		return Answer{r, c, ws}
	}

	extendRow = func(u, r int, n int) {
		var mh int = 100
		for i := 0; i < n; i++ {
			if !mem[colHash[i]] {
				// not a prefix
				return
			}
			mh = min(mh, mem2[colHash[i]])
		}

		if mh*n < ans.h*ans.w {
			// will not get a better answer
			return
		}

		var ok = true

		for i := 0; i < n; i++ {
			if !mem1[colHash[i]] {
				ok = false
				break
			}
		}
		if ok {
			ans = ans.Max(createAnswer(r, n))
		}
		backup := make([]int64, n)

		for _, v := range conn[u] {
			for j := 0; j < n; j++ {
				backup[j] = colHash[j]
				colHash[j] = (colHash[j]*B + int64(words[v][j])) % MOD
			}
			rows[r] = v
			extendRow(v, r+1, n)

			copy(colHash, backup)
		}
	}

	sort.Ints(lens)

	for i := len(lens) - 1; i >= 0; i-- {
		k := lens[i]

		ws := wordsByLen[k]

		for _, u := range ws {
			for j := 0; j < len(words[u]); j++ {
				colHash[j] = int64(words[u][j])
			}
			rows[0] = u
			extendRow(u, 1, len(words[u]))
		}
	}

	return ans.words
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
func maxRectangle1(words []string) []string {
	prefix := make([]map[int64]int, 26)

	for i := 0; i < 26; i++ {
		prefix[i] = make(map[int64]int)
	}

	ii := make(map[int64]int)
	var ans Answer

	for id, word := range words {
		if len(word) == 1 {
			ans = Answer{1, 1, []string{word}}
		}
		hash := int64(word[0] - 'a' + 1)

		for i := 1; i < len(word); i++ {
			x := int(word[i] - 'a')
			prefix[x][hash]++
			hash = (hash*B + int64(x) + 1) % MOD
		}
		ii[hash] = id
	}

	colHash := make([]int64, 100)
	backup := make([]int64, 100)
	rows := make([]int, 100)

	allWords := func(n int) bool {
		for j := 0; j < n; j++ {
			if _, found := ii[colHash[j]]; !found {
				return false
			}
		}
		return true
	}

	createAnswer := func(r, c int) Answer {
		ws := make([]string, r)
		for i := 0; i < r; i++ {
			ws[i] = words[rows[i]]
		}
		return Answer{r, c, ws}
	}

	var extendCol func(rowHash int64, c int, n int) int

	extendCol = func(rowHash int64, c int, n int) int {
		if c == n {
			if i, found := ii[rowHash]; found {
				return i
			}
			return -1
		}

		for x := 0; x < 26; x++ {
			if _, found := prefix[x][colHash[c]]; !found {
				continue
			}

			if prefix[x][rowHash] == 0 {
				continue
			}
			i := extendCol((rowHash*B+int64(x)+1)%MOD, c+1, n)
			if i < 0 {
				continue
			}
			return i
		}
		return -1
	}

	var extendRow func(r, n int)

	extendRow = func(r, n int) {
		if r > 100 {
			return
		}

		if allWords(n) {
			ans = ans.Max(createAnswer(r, n))
		}

		for x := 0; x < 26; x++ {
			// try put x at first column
			if _, found := prefix[x][colHash[0]]; !found {
				continue
			}
			i := extendCol(int64(x)+1, 1, n)
			if i < 0 {
				continue
			}
			rows[r] = i

			for j := 0; j < n; j++ {
				backup[j] = colHash[j]
				colHash[j] = (colHash[j]*B + int64(words[i][j]-'a'+1)) % MOD
			}
			extendRow(r+1, n)
			for j := 0; j < n; j++ {
				colHash[j] = backup[j]
			}
		}
	}

	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			colHash[j] = int64(words[i][j] - 'a' + 1)
		}
		rows[0] = i
		extendRow(1, len(words[i]))
	}

	return ans.words
}

type Answer struct {
	w, h  int
	words []string
}

func (this Answer) Max(that Answer) Answer {
	if this.w*this.h >= that.w*that.h {
		return this
	}
	return that
}
