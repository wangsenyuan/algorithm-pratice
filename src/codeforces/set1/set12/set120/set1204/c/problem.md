You are given a directed unweighted graph without loops with 𝑛
 vertexes and a path in it (that path is not necessary simple) given by a sequence 𝑝1,𝑝2,…,𝑝𝑚
 of 𝑚
 vertexes; for each 1≤𝑖<𝑚
 there is an arc from 𝑝𝑖
 to 𝑝𝑖+1
.

Define the sequence 𝑣1,𝑣2,…,𝑣𝑘
 of 𝑘
 vertexes as good, if 𝑣
 is a subsequence of 𝑝
, 𝑣1=𝑝1
, 𝑣𝑘=𝑝𝑚
, and 𝑝
 is one of the shortest paths passing through the vertexes 𝑣1
, …
, 𝑣𝑘
 in that order.

A sequence 𝑎
 is a subsequence of a sequence 𝑏
 if 𝑎
 can be obtained from 𝑏
 by deletion of several (possibly, zero or all) elements. It is obvious that the sequence 𝑝
 is good but your task is to find the shortest good subsequence.

If there are multiple shortest good subsequences, output any of them.

### ideas
1. d1[i] 表示在序列i到p1的最短距离
2. p[i]如果被删除，那么, 前面的某个节点到后面的某个节点的距离，不能变短
3. 否则p[i]就不能删除
4. 假设dp[i][v] 表示前i个序列，最后落脚点在v时，是否ok？
5. p[i+1] = w
6. 如果 v, w 的最短距离是x, dp[i][w] = dp[i-x][v] + 1