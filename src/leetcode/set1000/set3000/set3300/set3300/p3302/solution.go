package p3302

func validSequence(word1 string, word2 string) []int {
	n := len(word2)
	// dp[i] 表示word2的后缀能够匹配的最右端的位置
	m := len(word1)

	if n > m {
		return nil
	}

	dp := make([]int, n+1)
	dp[n] = m

	for i, j := n-1, m-1; i >= 0; i-- {
		dp[i] = -1
		for j >= 0 && word2[i] != word1[j] {
			j--
		}
		if j < 0 {
			break
		}
		dp[i] = j
		j--
	}

	getLast := func(arr []int) int {
		if len(arr) == 0 {
			return -1
		}
		return arr[len(arr)-1]
	}

	process := func(arr []int, i int, j int) []int {
		for ; i < n; i++ {
			for j < m && word2[i] != word1[j] {
				j++
			}
			arr = append(arr, j)
			j++
		}
		return arr
	}

	var arr []int
	for i, j := 0, 0; i < n; i++ {
		for j < m && word2[i] != word1[j] {
			j++
		}
		p := getLast(arr)
		if p+1 < j && dp[i+1] > p+1 {
			// 这个时候进行处理是能够获得一个更好的结果的
			arr = append(arr, p+1)
			arr = process(arr, i+1, p+2)
			return arr
		}

		if j == m {
			return nil
		}

		arr = append(arr, j)
		j++

	}

	return arr
}
