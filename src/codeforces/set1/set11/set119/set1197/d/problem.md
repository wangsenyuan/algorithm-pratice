You are given an array 𝑎1,𝑎2,…,𝑎𝑛
 and two integers 𝑚
 and 𝑘
.

You can choose some subarray 𝑎𝑙,𝑎𝑙+1,…,𝑎𝑟−1,𝑎𝑟
.

The cost of subarray 𝑎𝑙,𝑎𝑙+1,…,𝑎𝑟−1,𝑎𝑟
 is equal to ∑𝑖=𝑙𝑟𝑎𝑖−𝑘⌈𝑟−𝑙+1𝑚⌉
, where ⌈𝑥⌉
 is the least integer greater than or equal to 𝑥
.

The cost of empty subarray is equal to zero.

For example, if 𝑚=3
, 𝑘=10
 and 𝑎=[2,−4,15,−3,4,8,3]
, then the cost of some subarrays are:

𝑎3…𝑎3:15−𝑘⌈13⌉=15−10=5
;
𝑎3…𝑎4:(15−3)−𝑘⌈23⌉=12−10=2
;
𝑎3…𝑎5:(15−3+4)−𝑘⌈33⌉=16−10=6
;
𝑎3…𝑎6:(15−3+4+8)−𝑘⌈43⌉=24−20=4
;
𝑎3…𝑎7:(15−3+4+8+3)−𝑘⌈53⌉=27−20=7
.
Your task is to find the maximum cost of some subarray (possibly empty) of array 𝑎
.

### ideas
1. m比较小，只有10。这个可能是个关键点
2. cost -= 这个区间的长度 / m的上界 * k
3. 一个数字，要么开始一个新的区间,要么加入前面的区间
4. dp[i][j] 表示数字i加入前面的区间，且它的index % m = j 时的最优解
5. 如果 j = 0, j > 0
6. dp[i][j] = dp[i-1][j-1] + a[i] (如果 j > 0)
7.           max(dp[i-1][m-1], 0) + a[i] - k (0表示它开启一个新的区间，m-1表示和前面的区间连接起来)