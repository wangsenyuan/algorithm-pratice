package p819

import "strings"

func mostCommonWord(paragraph string, banned []string) string {
	freq := make(map[string]int)

	stop := make(map[byte]bool)
	stopLetters := "!?',;. "
	for i := 0; i < len(stopLetters); i++ {
		stop[stopLetters[i]] = true
	}

	ban := make(map[string]bool)
	for _, word := range banned {
		ban[word] = true
	}

	n := len(paragraph)
	for i, j := 0, 0; i <= n; i++ {
		if i == n || stop[paragraph[i]] {
			if j < i {
				word := strings.ToLower(paragraph[j:i])
				if !ban[word] {
					freq[word]++
				}
			}
			j = i + 1
		}
	}

	var ans string
	var maxFreq int

	for k, v := range freq {
		if v > maxFreq {
			ans = k
			maxFreq = v
		}
	}
	return ans
}
