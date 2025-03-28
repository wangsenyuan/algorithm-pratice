package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s := readString(reader)
	res := solve(s)
	fmt.Println(len(res))
	for _, cur := range res {
		fmt.Println(cur[0], cur[1])
	}
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func solve(s string) [][]int {
	l := fastLongestPalindromes(s)

	n := len(s)
	var ans [][]int
	for i := 0; i < n; i++ {
		if len(ans) == 0 || ans[0][1] < l[2*i+1] {
			ans = [][]int{{i - l[2*i+1]/2 + 1, l[2*i+1]}}
		}
	}

	if ans[0][1] == n {
		return ans
	}
	rs := reverse(s)
	lps := kmp(rs)

	fp := make([]int, n+1)
	for i := range n + 1 {
		fp[i] = n
	}
	fp[0] = -1

	for i, j := 0, 0; i < n; i++ {
		for j > 0 && s[i] != rs[j] {
			j = lps[j-1]
		}
		if s[i] == rs[j] {
			j++
		}
		if i < n-j {
			fp[j] = min(fp[j], i)
		}
	}

	var best int
	var ans2 [][]int

	update := func(p1 int, p2 int, p3 int) {
		tmp := 2*(n-p3) + l[p2*2+1]
		if p3 == n || tmp <= best {
			return
		}
		best = tmp

		ans2 = ans2[:0]

		ans2 = append(ans2, []int{p1 + 1, n - p3})
		ans2 = append(ans2, []int{p2 - l[2*p2+1]/2 + 1, l[2*p2+1]})
		ans2 = append(ans2, []int{p3 + 1, n - p3})
	}

	for i := range n {
		half := (l[2*i+1] + 1) / 2
		r := i + half
		l := i - half
		// 考虑后缀 0,1,2....j
		j := n - r
		k := sort.Search(j+1, func(k int) bool {
			return fp[k] > l
		})
		k--
		// fp[k] <= l
		if k >= 0 {
			update(fp[k]-k+1, i, n-k)
		}
	}

	if best > ans[0][1] {
		return ans2
	}
	return ans
}

func kmp(s string) []int {
	n := len(s)
	p := make([]int, n)
	for i := 1; i < n; i++ {
		j := p[i-1]
		for j > 0 && s[j] != s[i] {
			j = p[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		p[i] = j
	}
	return p
}

func reverse(s string) string {
	buf := []byte(s)
	for i, j := 0, len(buf)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}
	return string(buf)
}

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

func bruteForce(s string) [][]int {
	l := fastLongestPalindromes(s)

	n := len(s)
	var ans [][]int
	for i := 0; i < n; i++ {
		if len(ans) == 0 || ans[0][1] < l[2*i+1] {
			ans = [][]int{{i - l[2*i+1]/2 + 1, l[2*i+1]}}
		}
	}

	if ans[0][1] == n {
		return ans
	}
	rs := reverse(s)

	lps := kmp(rs)

	var best int

	var ans2 [][]int

	for i, j := 0, 0; i < n; i++ {
		for j > 0 && s[i] != rs[j] {
			j = lps[j-1]
		}
		if s[i] == rs[j] {
			j++
		}
		if i+j >= n {
			break
		}
		if j == 0 {
			continue
		}
		x := s[i+1 : n-j]
		tmp := fastLongestPalindromes(x)
		for k := 0; k < len(x); k++ {
			if tmp[2*k+1]+2*j > best {
				best = tmp[2*k+1] + 2*j
				ans2 = ans2[:0]
				ans2 = append(ans2, []int{i - j + 2, j})
				ans2 = append(ans2, []int{i + k - tmp[2*k+1]/2 + 2, tmp[2*k+1]})
				ans2 = append(ans2, []int{n - j + 1, j})
			}
		}
	}

	if best > ans[0][1] {
		return ans2
	}
	return ans
}
