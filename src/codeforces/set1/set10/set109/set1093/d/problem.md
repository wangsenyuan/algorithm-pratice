You are given an undirected unweighted graph consisting of 𝑛
 vertices and 𝑚
 edges.

You have to write a number on each vertex of the graph. Each number should be 1
, 2
 or 3
. The graph becomes beautiful if for each edge the sum of numbers on vertices connected by this edge is odd.

Calculate the number of possible ways to write numbers 1
, 2
 and 3
 on vertices so the graph becomes beautiful. Since this number may be large, print it modulo 998244353
.

Note that you have to write exactly one number on each vertex.

The graph does not have any self-loops or multiple edges.

### ideas
1. 如果两个点相连，他们的奇偶性必须相反，所以不能出现奇数的环
2. 然后，不同component的需求不一样
3. 假设在一个component中，分别有两个颜色的cnt0, cnt1, 如果其中一种选择偶数，两个一种就是奇数2 ** cnt0 + 2 ** cnt1