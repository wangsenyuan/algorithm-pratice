Parsa has a humongous tree on 𝑛
 vertices.

On each vertex 𝑣
 he has written two integers 𝑙𝑣
 and 𝑟𝑣
.

To make Parsa's tree look even more majestic, Nima wants to assign a number 𝑎𝑣
 (𝑙𝑣≤𝑎𝑣≤𝑟𝑣
) to each vertex 𝑣
 such that the beauty of Parsa's tree is maximized.

Nima's sense of the beauty is rather bizarre. He defines the beauty of the tree as the sum of |𝑎𝑢−𝑎𝑣|
 over all edges (𝑢,𝑣)
 of the tree.

Since Parsa's tree is too large, Nima can't maximize its beauty on his own. Your task is to find the maximum possible beauty for Parsa's tree.

Input
The first line contains an integer 𝑡
 (1≤𝑡≤250)
 — the number of test cases. The description of the test cases follows.

The first line of each test case contains a single integer 𝑛
 (2≤𝑛≤105)
 — the number of vertices in Parsa's tree.

The 𝑖
-th of the following 𝑛
 lines contains two integers 𝑙𝑖
 and 𝑟𝑖
 (1≤𝑙𝑖≤𝑟𝑖≤109)
.

Each of the next 𝑛−1
 lines contains two integers 𝑢
 and 𝑣
 (1≤𝑢,𝑣≤𝑛,𝑢≠𝑣)
 meaning that there is an edge between the vertices 𝑢
 and 𝑣
 in Parsa's tree.

It is guaranteed that the given graph is a tree.

It is guaranteed that the sum of 𝑛
 over all test cases doesn't exceed 2⋅105
.

### ideas
1. 显然没发尝试所有的 a[i] >= l[i] <= r[i]
2. 但是，考虑这样一种情况， 如果 a[u]已经确定了，现在考虑a[v]的赋值情况, 如果 a[v] >= a[u], 那么为了这段更大，那么最优的结果就是让a[v] = r[v]
3. 反之亦然，也就是可以让a[v]取 l[v]
4. 所以， dp[u][0] 表示 a[u] = l[u]时的最大值， dp[u][1] = a[u] = r[u]时的最大值