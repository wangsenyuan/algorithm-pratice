You are given a rooted tree consisting of 𝑛
 vertices numbered from 1
 to 𝑛
. Vertex 1
 is the root of the tree. Each vertex has an integer value. The value of 𝑖
-th vertex is 𝑎𝑖
. You can do the following operation at most 𝑘
 times.

Choose a vertex 𝑣
 that has not been chosen before and an integer 𝑥
 such that 𝑥
 is a common divisor of the values of all vertices of the subtree of 𝑣
. Multiply by 𝑥
 the value of each vertex in the subtree of 𝑣
.
What is the maximum possible value of the root node 1
 after at most 𝑘
 operations? Formally, you have to maximize the value of 𝑎1
.

A tree is a connected undirected graph without cycles. A rooted tree is a tree with a selected vertex, which is called the root. The subtree of a node 𝑢
 is the set of all nodes 𝑦
 such that the simple path from 𝑦
 to the root passes through 𝑢
. Note that 𝑢
 is in the subtree of 𝑢
.



#### ideas
1. 先思考一个问题，就是如果在最优的答案中，u，v都进行了操作， 且u是v的parent，那么是否存在先处理u，再处理v的情况？
2. 似乎是不好的。假设处理u，那么x divides gcd(a[u], a[v], ...)
3. 这个x的上限是被a[v]确定的。但是如果先处理v，那么这个上限是被a[v] * y决定的
4. 其次，如果要操作，似乎使用gcd(sub[v])是个更好的选择
5. 至少确定了两点，就是从下到上，且使用gcd
6. v被处理后，它就不能再被处理，所以可以确定一个点，v这个子树的贡献，最大到 a[v] * a[v] (前提a[v]成为了大家的最小公倍数)
7. 如果某个子树的gcd = 1, 会怎样？
8. 比如 4 -> 2 -> 3, => 不会变化
9. dp[u][x] = 将该子树的gcd提供pow x也就是是a[u] = pow(a[u], x)时所需的最小操作次数
10. dp[u][x] = sum(dp[v][y], 且 pow(a[v], y) % pow(a[u], x) = 0) 
11. x只能是1
12. 假设 u整棵树的gcd = g[u]
13. 如果变成质数2 * 3 * 5...
14. 对于每个质因子来说，相当于，使用这颗子树中的，最小的个数，提换root的个数
15. 比如，对于2来说，如果这个子树中，最小的是w，不管a[u]中有几个，操作1一次后， =》 2 * w
16. 还有一个点就是，如果在a[u]中，不存在的质因子，其实是不用考虑的
17. 进一步的，所有parent中，不存在的，都不需要考虑
18. 如果没有任何限制，也就是k = n, 那没有任何问题，直接所有的节点处理一遍
19. 现在主要是k有限制，就要有取舍
20. a[u]只能增加到 a[u] * a[u]
21. dp[a[u]][i] 表示将 a[u] * i的最小操作次数
22. 显然i越大，dp[u][i]越大
23. dp[u][i] = sum (dp[v][j] where a[v] * j % a[u] * i = 0, j越小越好)  