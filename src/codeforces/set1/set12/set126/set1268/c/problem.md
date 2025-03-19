You are given a permutation 𝑝1,𝑝2,…,𝑝𝑛
.

In one move you can swap two adjacent values.

You want to perform a minimum number of moves, such that in the end there will exist a subsegment 1,2,…,𝑘
, in other words in the end there should be an integer 𝑖
, 1≤𝑖≤𝑛−𝑘+1
 such that 𝑝𝑖=1,𝑝𝑖+1=2,…,𝑝𝑖+𝑘−1=𝑘
.

Let 𝑓(𝑘)
 be the minimum number of moves that you need to make a subsegment with values 1,2,…,𝑘
 appear in the permutation.

You need to find 𝑓(1),𝑓(2),…,𝑓(𝑛)
.

### ideas
1. range update with lazy segment
2. 假设0的开始位置是i，现在要把n放进去， 它的cost是多少呢？
3. 如果n在0...n-1的外部，这个cost应该 = abs(pos[n] - (i + n)) ?
4. 如果pos[n]在0...n-1的内部， 是不是就是0？
5. 因为再移动0.。。n-1的时候，自然的就可以把n移动到需要的位置去？
6. 不大对。 考虑序列 3, 2, 1, 0
7. dp[0] = 3, dp[1] = 2, dp[2] = 1, dp[3] = 0, f(0) = 0
8. dp[0] = 4, dp[1] = 2, dp[2] = 1, dp[3] = inf, f(1) = 1
9. dp[0] = 4, dp[1] = 3, dp[2] = inf, dp[3] = inf, f(2) = 3
10. dp[0] = 6, dp[1] = inf, ... f(3) = 6
11. 第一个点就是，把新的n加入到里面后， 怎么计算新的cost？
12. 还有就是， 新的dp，貌似不能从原有的dp直接得到
13. 特别的在处理3的时候，并不是把dp[0] + 3从位置0移动到位置3的cost
14. 好像和inversion count 有关系
15. 只考虑0...n，至少要使用inversion count 的操作次数，来使的它们排好序
16. 先把它们聚合在一起，占据位置i...i+n-1, 然后排序
17. 比如 1,3,2,0,4 要计算f(2), 先得到120，假设起点在1， 那么移动距离 = 1 + 0 + 0 = 1
18. 这个貌似不能dp，应该是贪心。就是把i移动到靠近中间的位置
19. 假设这个区间是 l...r, 那么移动的最短距离 = ？
20. dp[i] = 表示最左边的位置在i的时候，把前n-1移动到一起时的最短移动距离（不考虑inversion count)
21. 比如把1，2，0移动到位置0开始，需要花费2
22. 现在考虑3，花费居然是0.。。
23. i...i+n-1, 这个区间的cost = 在这个区间外边的, 但是value 却在这个区间里边的移动的距离之和
24. 假设l...r的第mid个点，那么只要把0...n的mid和它对齐就可以了
25. 假设起点是i, i - pos[0] + i + 1 - pos[1] + ... i + mid - pos[mid] + pos[mid+1] - (i+mid+1) + ... + pos[n] - (i + n)
26. l - pos[0] + pos[n] - r = pos[n] - pos[0] - (r - l) 
27. l + 1 - pos[1] + pos[n-1] - (r - 1) = pos[n-1] - pos[1] - (r - l) + 2
28. ...
29. l + mid - pos[mid] + pos[n - mid] - (r + mid ) = pos[n-mid] - pos[mid] - (r - l) + 2 * mid 
30. 所以，其实可以和i无关的
31. sum of second half - sum of first half - n / 2 * n + 2 * (0 + ... + mid)
32. 前n个的移动距离 + inversion count 就是答案
33. 移动距离 = 前一半的位置和的负数 + 后一半的位置和 
34. inversion count 也很好处理
35. 居然搞对了～～