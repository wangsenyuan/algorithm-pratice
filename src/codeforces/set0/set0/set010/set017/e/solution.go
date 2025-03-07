package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	readString(reader)
	s := readString(reader)
	res := solve1(s)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

const mod = 51123987

func add(a, b int) int {
	a %= mod
	b %= mod
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func sub(a, b int) int {
	return add(a, mod-b)
}

func mul(nums ...int) int {
	r := 1
	for _, num := range nums {
		num %= mod
		r = r * num % mod
	}
	return r
}

func solve(s string) int {
	n := len(s)
	t := append(make([]byte, 0, n*2+3), '^')
	for _, c := range []byte(s) {
		t = append(t, '#', c)
	}
	t = append(t, '#', '$')
	halfLen := make([]int, len(t)-2)
	halfLen[1] = 1
	boxM, boxR := 0, 0
	var tot int

	for i := 2; i < len(halfLen); i++ {
		hl := 1
		if i < boxR {
			hl = min(halfLen[boxM*2-i], boxR-i)
		}
		for t[i-hl] == t[i+hl] {
			hl++
			boxM, boxR = i, i+hl
		}
		halfLen[i] = hl
		tot += hl / 2
	}
	tot %= mod
	ans := tot * (tot - 1) / 2

	diff := make([]int, n+1)
	for i, hl := range halfLen {
		l, r := (i-hl)/2, (i+hl)/2-2
		if l <= r {
			diff[(l+r+1)/2]++
			diff[r+1]--
		}
	}
	cntEnd := diff[:n]
	for i := 1; i < n; i++ {
		cntEnd[i] += cntEnd[i-1]
	}
	for i := 1; i < n; i++ {
		cntEnd[i] += cntEnd[i-1]
	}

	sum := make([]int, n+1)
	for i, v := range cntEnd {
		sum[i+1] = (sum[i] + v) % mod
	}
	for i, hl := range halfLen {
		l, r := (i-hl)/2, (i+hl)/2-2
		if l <= r {
			ans -= sum[(l+r)/2] - sum[max(l-1, 0)]
		}
	}

	ans %= mod

	if ans < 0 {
		ans += mod
	}

	return ans
}

func solve1(s string) int {
	l := fastLongestPalindromes(s)
	n := len(s)

	var tot int
	dp := make([]int, n+1)
	fp := make([]int, n+1)
	for i := range n {
		if l[2*i+1] > 0 {
			cnt := (l[2*i+1] + 1) / 2
			tot += cnt
			dp[i]++
			dp[i+cnt]--
			fp[i]++
			if i-cnt >= 0 {
				fp[i-cnt]--
			}
		}
		if l[2*i] > 0 {
			// 中点在i-1, i中间
			cnt := l[2*i] / 2
			tot += cnt
			dp[i]++
			dp[i+cnt]--
			fp[i-1]++
			if i-1-cnt >= 0 {
				fp[i-1-cnt]--
			}
		}
	}
	var res int
	if tot%2 == 0 {
		res = mul(tot/2, tot-1)
	} else {
		res = mul(tot, (tot-1)/2)
	}
	// dp[i]表示在右端在i处的计数
	for i := 1; i < n; i++ {
		dp[i] = add(dp[i], dp[i-1])
	}

	var suf int
	for i := n - 1; i > 0; i-- {
		fp[i] = add(fp[i], fp[i+1])
		suf = add(suf, fp[i])
		tmp := mul(dp[i-1], suf)
		res = sub(res, tmp)
	}

	return res
}

/*
*
*

	"""
	   Given a sequence seq, returns a list l such that l[2 * i + 1]
	   holds the length of the longest palindrome centered at seq[i]
	   (which must be odd), l[2 * i] holds the length of the longest
	   palindrome centered between seq[i - 1] and seq[i] (which must be
	   even), and l[2 * len(seq)] holds the length of the longest
	   palindrome centered past the last element of seq (which must be 0,
	   as is l[0]).

	   The actual palindrome for l[i] is seq[s:(s + l[i])] where s is i
	   // 2 - l[i] // 2. (// is integer division.)
*/
func fastLongestPalindromes(seq string) []int {
	n := len(seq)
	var l []int
	var palLen int
	// Loop invariant: seq[(i - palLen):i] is a palindrome.
	// Loop invariant: len(l) >= 2 * i - palLen. The code path that
	// increments palLen skips the l-filling inner-loop.
	// Loop invariant: len(l) < 2 * i + 1. Any code path that
	// increments i past seqLen - 1 exits the loop early and so skips
	// the l-filling inner loop.
	for i := 0; i < n; {
		// First, see if we can extend the current palindrome.  Note
		// that the center of the palindrome remains fixed.
		if i > palLen && seq[i-palLen-1] == seq[i] {
			palLen += 2
			i++
			continue
		}

		l = append(l, palLen)
		// Now to make further progress, we look for a smaller
		// palindrome sharing the right edge with the current
		// palindrome.  If we find one, we can try to expand it and see
		// where that takes us.  At the same time, we can fill the
		// values for l that we neglected during the loop above. We
		// make use of our knowledge of the length of the previous
		// palindrome (palLen) and the fact that the values of l for
		// positions on the right half of the palindrome are closely
		// related to the values of the corresponding positions on the
		// left half of the palindrome.

		// Traverse backwards starting from the second-to-last index up
		// to the edge of the last palindrome.
		s := len(l) - 2
		e := s - palLen
		found := false
		for j := s; j > e; j-- {
			// d is the value l[j] must have in order for the
			// palindrome centered there to share the left edge with
			// the last palindrome.  (Drawing it out is helpful to
			// understanding why the - 1 is there.)
			d := j - e - 1
			if l[j] == d {
				found = true
				palLen = d
				// We actually want to go to the beginning of the outer
				// loop,
				break
			}
			// Otherwise, we just copy the value over to the right
			// side.  We have to bound l[i] because palindromes on the
			// left side could extend past the left edge of the last
			// palindrome, whereas their counterparts won't extend past
			// the right edge.
			l = append(l, min(d, l[j]))
		}
		if !found {
			palLen = 1
			i++
		}
	}
	// We know from the loop invariant that len(l) < 2 * seqLen + 1, so
	// we must fill in the remaining values of l.

	// Obviously, the last palindrome we're looking at can't grow any
	// more.
	l = append(l, palLen)
	s := len(l) - 2
	e := s - (2*n + 1 - len(l))
	for i := s; i > e; i-- {
		// The d here uses the same formula as the d in the inner loop
		// above.  (Computes distance to left edge of the last
		// palindrome.)
		d := i - e - 1
		l = append(l, min(d, l[i]))
	}
	return l
}
