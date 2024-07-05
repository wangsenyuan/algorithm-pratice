You are given an array 𝑎
 consisting of 𝑛
 integers. You should divide 𝑎
 into continuous non-empty subarrays (there are 2𝑛−1
 ways to do that).

Let 𝑠=𝑎𝑙+𝑎𝑙+1+…+𝑎𝑟
. The value of a subarray 𝑎𝑙,𝑎𝑙+1,…,𝑎𝑟
 is:

(𝑟−𝑙+1)
 if 𝑠>0
,
0
 if 𝑠=0
,
−(𝑟−𝑙+1)
 if 𝑠<0
.
What is the maximum sum of values you can get with a partition?

### ideas
1. 考虑一个TLE的dp， dp[i] = max(dp[j] + score[j...i])
2. score[j...i] = (i - j) * sign(sum(j...i))
3. 考虑这样的j，sum(j...i)是正值 > 0
4. dp[i] = max(dp[j] + (i - j)) = i + max(dp[j] - j)
5. 如果 sum[j...i] < 0
6. dp[i] = max(dp[j] - (i - j)) = -i + max(dp[j] + j)
7. 这里考虑两种情况
8. 但是对于i如何找到那些 sum(j...i)是正值的，且从中找到最大值呢？
9. sum(j..i) = sum(i) - sum(j)
10. 也就是说对于i来说，找到那些小于 sum(i) 或者 大于sum(i) 的部分，其中的最大值
11. 所以要用sum(i) 做key，在其中寻找最大值
12. sum(i) 只有n个，所以可以采用数轴压缩的方式