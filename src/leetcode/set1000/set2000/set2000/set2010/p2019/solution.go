package p2019

/**

You are given a string s that contains digits 0-9, addition symbols '+', and multiplication symbols '*' only, representing a valid math expression of single digit numbers (e.g., 3+5*2). This expression was given to n elementary school students. The students were instructed to get the answer of the expression by following this order of operations:

Compute multiplication, reading from left to right; Then,
Compute addition, reading from left to right.
You are given an integer array answers of length n, which are the submitted answers of the students in no particular order. You are asked to grade the answers, by following these rules:

If an answer equals the correct answer of the expression, this student will be rewarded 5 points;
Otherwise, if the answer could be interpreted as if the student used the incorrect order of operations, once or multiple times, this student will be rewarded 2 points;
Otherwise, this student will be rewarded 0 points.
Return the sum of the points of the students.



来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/the-score-of-students-solving-math-expression
*/

func scoreOfStudents(s string, answers []int) int {
	l := len(s)
	h := (l + 1) / 2
	// every num will company by a ( or ),
	// a+b*c => ((a+b)*c) or (a+(b*c)) or (a)+((b)*c)
	good := calc(s)

	dp := make([][]map[int]bool, h)

	for i := 0; i < h; i++ {
		dp[i] = make([]map[int]bool, h)
		for j := 0; j < h; j++ {
			dp[i][j] = make(map[int]bool)
		}
		dp[i][i][int(s[2*i]-'0')] = true
	}

	for l := 2; l <= h; l++ {
		for i := 0; i+l <= h; i++ {
			j := i + l - 1
			// dp[i][j]
			for k := i; k < j; k++ {
				// dp[i][k] * dp[k+1][j]
				op := s[2*k+1]
				for x := range dp[i][k] {
					for y := range dp[k+1][j] {
						z := x + y
						if op == '*' {
							z = x * y
						}
						if z <= 1000 {
							dp[i][j][z] = true
						}
					}
				}
			}
		}
	}

	var res int

	cnt := make([]int, 1001)

	for _, cur := range answers {
		cnt[cur]++
	}

	for k, v := range cnt {
		if k == good {
			res += 5 * v
		} else if dp[0][h-1][k] {
			res += 2 * v
		}
	}

	return res
}

func calc(s string) int {
	n := len(s)
	correct := 0
	for i := 0; i < n; {
		mul := int(s[i] & 15)
		for i += 2; i < n && s[i-1] == '*'; i += 2 {
			mul *= int(s[i] & 15)
		}
		correct += mul
	}
	return correct
}
