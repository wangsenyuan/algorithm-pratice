Given an array of integers 𝑠1,𝑠2,…,𝑠𝑙
, every second, cosmic rays will cause all 𝑠𝑖
 such that 𝑖=1
 or 𝑠𝑖≠𝑠𝑖−1
 to be deleted simultaneously, and the remaining parts will be concatenated together in order to form the new array 𝑠1,𝑠2,…,𝑠𝑙′
.

Define the strength of an array as the number of seconds it takes to become empty.

You are given an array of integers compressed in the form of 𝑛
 pairs that describe the array left to right. Each pair (𝑎𝑖,𝑏𝑖)
 represents 𝑎𝑖
 copies of 𝑏𝑖
, i.e. 𝑏𝑖,𝑏𝑖,⋯,𝑏𝑖𝑎𝑖 times
.

For each 𝑖=1,2,…,𝑛
, please find the strength of the sequence described by the first 𝑖
 pairs.



### ideas
1. 考虑，对于一个序列，如何最快的消灭完？
2. f(arr) <= len(arr) 每次都消灭最前面的，那么正好需要这样的长度
3. 或者，如果存在一些 i, arr[i] != arr[i-1]， 那么使用这个策略更优
4. 1, 1, 2, 2, 3, 4, 4, 4, 5, 5
5. 1， 2， 4， 4， 5
6. 1， 4
7. 1
8. ’‘
9. 似乎和最长的子序列的长度有关
10. 1，2，2， 1，2，2
11. 1, 2, 2
12. 1, 2
13. 1
14. ''
15. 如果 b[i] = b[j], 它们之间的距离（还没有想清楚）如果太短了，那么它们是有机会接起来
16. 否则的话，它的贡献 = a[i], dp[i] = max(dp[i-1], a[i])
17. 然后考虑这个距离是什么
18. 在上面的例子里， 第一段2和第二段2的距离 = 1, 最长的序列的长度
19. dp[i] = max(dp[i-1], a[i], dp[j-1] + a[j] + a[i] - dist)
20. 