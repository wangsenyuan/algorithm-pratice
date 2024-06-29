One day Masha came home and noticed n mice in the corridor of her flat. Of course, she shouted loudly, so scared mice started to run to the holes in the corridor.

The corridor can be represeted as a numeric axis with n mice and m holes on it. ith mouse is at the coordinate xi, and jth hole — at coordinate pj. jth hole has enough room for cj mice, so not more than cj mice can enter this hole.

What is the minimum sum of distances that mice have to go through so that they all can hide in the holes? If ith mouse goes to the hole j, then its distance is |xi - pj|.

Print the minimum sum of distances.


### ideas

 dp[i][j] = min(dp[i-1][k] + last (j - k) mouse move to ith hole 的距离)
		 要迭代j和k => TLE
		 迭代k， 如果 dp[i-1][k]确定的情况下， i的容量是 c[i]
		 那么1....c[i]个老鼠可以跑进去
		 那么j = k+1....k+c[i]
		 假设这个老鼠在p[i]的后面，x[j] - p[i] + dp[i-1][k]
		 如果在p[i]的前面 p[i] - x[j] + dp[i-1][k]
		 在迭代k的情况下，对于k后面的c[i]个老鼠
		 需要把它们分成两组，一组是在p[i]的前面，一组在后面
		 假设有u个老鼠跑进去了, 假设这u个全在前面那组
		 dp[i][k+u] = dp[i-1][k] + u * p[i] - sum of those u (x[j])
		    如果u个前面，v个在后面
		 dp[i][k+u+v] = dp[i-1][k] + u * p[i] - sum u (x[j]) + sum v (x[j]) - v * p[i]
		 感觉是可以批量去更新的
		 考虑第一个case， 怎么去更新，p[i] - x[j] 的前缀和
		 dp[i-1][k] + sum of (p[i] - x[j])
		 分成两种情况去处理，第一种，处理所有p[i]前面的老鼠
		 这样子符号是确定的，对于这样的j来说，就是dp[i-1][k] + (sum[j] - sum[k])
		 dp[i-1][k] + sum[k]最小的那个
		 这个地方不大对，

		 x[r] > holes[i][0]
		 这个范围内的 j = dp[i-1][k] + sum[j] - sum[k]
		  需要使用最小的 dp[i-1][k] - sum[k] 来更新dp[j]