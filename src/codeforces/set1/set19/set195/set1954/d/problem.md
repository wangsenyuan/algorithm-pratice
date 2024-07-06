There are balls of 𝑛
 different colors; the number of balls of the 𝑖
-th color is 𝑎𝑖
.

The balls can be combined into groups. Each group should contain at most 2
 balls, and no more than 1
 ball of each color.

Consider all 2𝑛
 sets of colors. For a set of colors, let's denote its value as the minimum number of groups the balls of those colors can be distributed into. For example, if there are three colors with 3
, 1
 and 7
 balls respectively, they can be combined into 7
 groups (and not less than 7
), so the value of that set of colors is 7
.

Your task is to calculate the sum of values over all 2𝑛
 possible sets of colors. Since the answer may be too large, print it modulo 998244353
.

### ideas
1. 对于一个set， 它的总数是sum，最大的是x, 那么需要的最小的group = max(x, (sum + 1) / 2)
2. 所以，这里需要按照a[i]进行升序排列
3. dp[i][sum] 表示以i为最大值，且总和为sum的贡献, fp[i][sum] 是计数
4. dp[i][sum] = dp[i-1][sum - a[i]] + y * fp[j][sum - a[i]] + dp[i-1][sum]
5. fp[i][sum] = fp[i-1][sum] + fp[i-1][sum - a[i]]