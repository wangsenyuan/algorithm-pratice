package p1895

func minOperationsToFlip(expression string) int {
	n := len(expression)
	// n <= 1e5
	stack := make([]int, n)
	var p int
	pair := make([]int, n)

	for i := 0; i < n; i++ {
		pair[i] = i
		if expression[i] == '(' {
			stack[p] = i
			p++
		} else if expression[i] == ')' {
			pair[stack[p-1]] = i
			p--
		}
	}

	var dfs func(i, j int) []int
	dfs = func(i, j int) []int {
		ans := make([]int, 3)
		if i == j {
			v := int(expression[i] - '0')
			ans[v] = 0
			ans[1-v] = 1
			ans[2] = v
		} else if pair[i] == j {
			ans = dfs(i+1, j-1)
		} else {
			var arr [][]int
			arr = append(arr, dfs(i, pair[i]))

			for k := pair[i] + 1; k <= j; {
				// expression[k] has to be & or |
				arr = append(arr, dfs(k+1, pair[k+1]))
				k = pair[k+1] + 1
			}
			if len(arr) == 1 {
				ans = arr[0]
			} else {
				v := arr[0][2]
				id := 1
				for k := pair[i] + 1; k <= j; {
					if expression[k] == '&' {
						if v == 1 {
							v = arr[id][2]
						}
					} else {
						if v == 0 {
							v = arr[id][2]
						}
					}
					id++
					k = pair[k+1] + 1
				}
				ans[2] = v
				// we need to know ans[1-v]
				id = 1
				dp := make([]int, 2)
				dp[0] = arr[0][0]
				dp[1] = arr[0][1]
				for k := pair[i] + 1; k <= j; {
					a, b := dp[0], dp[1]
					if expression[k] == '&' {
						// leave it unchange
						dp[0] = min(a, b) + arr[id][0]
						dp[0] = min(dp[0], a+min(arr[id][0], arr[id][1]))
						dp[1] = b + arr[id][1]
						// change & to |
						dp[0] = min(dp[0], a+1+arr[id][0])
						dp[1] = min(dp[1], min(a, b)+1+arr[id][1])
						dp[1] = min(dp[1], b+min(arr[id][0], arr[id][1])+1)
					} else {
						// leave it unchange
						dp[0] = a + arr[id][0]
						dp[1] = b + min(arr[id][0], arr[id][1])
						dp[1] = min(dp[1], min(a, b)+arr[id][1])
						// change it to &
						dp[0] = min(dp[0], min(a, b)+1+arr[id][0])
						dp[0] = min(dp[0], a+1+min(arr[id][0], arr[id][1]))
						dp[1] = min(dp[1], b+1+arr[id][1])
					}
					id++
					k = pair[k+1] + 1
				}
				ans[v] = dp[v]
				ans[1-v] = dp[1-v]
			}
		}
		return ans
	}
	res := dfs(0, n-1)
	return res[1-res[2]]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
