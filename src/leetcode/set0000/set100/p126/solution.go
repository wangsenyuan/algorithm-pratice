package p126

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	mem := make(map[string]bool)

	for _, word := range wordList {
		mem[word] = true
	}

	if !mem[endWord] {
		return nil
	}

	dist := make(map[string]int)
	que := make([]string, len(wordList)+1)
	var front, end int

	que[end] = beginWord
	dist[beginWord] = 0
	end++

	for front < end {
		cur := que[front]
		front++
		nexts := changeOneLetter(cur, mem)

		for _, next := range nexts {
			if _, found := dist[next]; !found {
				dist[next] = dist[cur] + 1
				que[end] = next
				end++
			}
		}
	}

	if _, found := dist[endWord]; !found {
		return nil
	}

	h := dist[endWord]

	path := make([]string, h+1)

	var ans [][]string

	var dfs func(cur string, d int)

	dfs = func(cur string, d int) {
		if d > h {
			return
		}
		path[d] = cur
		if cur == endWord {
			tmp := make([]string, h+1)
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}

		nexts := changeOneLetter(cur, mem)

		for _, next := range nexts {
			if dist[next] == d+1 {
				dfs(next, d+1)
			}
		}
	}

	dfs(beginWord, 0)

	return ans

}

func changeOneLetter(s string, mem map[string]bool) []string {
	bs := []byte(s)
	var res []string
	for i := 0; i < len(s); i++ {
		tmp := s[i]
		x := int(s[i] - 'a')
		for j := 0; j < 26; j++ {
			if x == j {
				continue
			}
			y := byte('a' + j)
			bs[i] = y
			if mem[string(bs)] {
				res = append(res, string(bs))
			}
		}
		bs[i] = tmp
	}

	return res
}
