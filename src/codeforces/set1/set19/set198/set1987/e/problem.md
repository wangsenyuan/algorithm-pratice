A Random Pebble
You are given a tree with 𝑛
 vertices, rooted at vertex 1
. The 𝑖
-th vertex has an integer 𝑎𝑖
 written on it.

Let 𝐿
 be the set of all direct children∗
 of 𝑣
. A tree is called wonderful, if for all vertices 𝑣
 where 𝐿
 is not empty,
𝑎𝑣≤∑𝑢∈𝐿𝑎𝑢.
In one operation, you choose any vertex 𝑣
 and increase 𝑎𝑣
 by 1
.

Find the minimum number of operations needed to make the given tree wonderful!

∗
 Vertex 𝑢
 is called a direct child of vertex 𝑣
 if:

𝑢
 and 𝑣
 are connected by an edge, and
𝑣
 is on the (unique) path from 𝑢
 to the root of the tree.
Input
Each test contains multiple test cases. The first line of input contains a single integer 𝑡
 (1≤𝑡≤1000
) — the number of test cases. The description of the test cases follows.

The first line of each test case contains a single integer 𝑛
 (2≤𝑛≤5000
) — the number of vertices in the tree.

The second line of each test case contains 𝑛
 integers 𝑎1,𝑎2,…,𝑎𝑛
 (0≤𝑎𝑖≤109
) — the values initially written on the vertices.

The third line of each test case contains 𝑛−1
 integers 𝑝2,𝑝3,…,𝑝𝑛
 (1≤𝑝𝑖<𝑖
), indicating that there is an edge from vertex 𝑝𝑖
 to vertex 𝑖
. It is guaranteed that the given edges form a tree.

It is guaranteed that the sum of 𝑛
 over all test cases does not exceed 5000
.

Output
For each test case, output a single integer — the minimum number of operations needed to make the tree wonderful.

### ideas
1. 显然叶子节点时不会增加的？好像也可以增加，否则增加它们的parent，有可能造成更上层的节点也要增加
2. 但是root似乎时不能增加的，它增加了，会增加direct child的的下限
3. 此外对于每个节点来说，似乎时能不增加就不增加
4. 假设sum(children of root) < a[root]
5. 此时必须增加某(个/些)child的值
6. 但是这样会造成连锁反应
7. 假设 dp[u]是在只考虑子树u时的最优解
8. 怎么处理v(u的parent呢)
9. 好像和节点所在的层有关系
10. 比如a[0] <= sum(l 1) <= sum(l, 2) <=  ... sum(l, i)
11. 对于每个子树，这个条件都要成立
12. 那这个条件时充分条件吗？似乎也是的。因为 sum(l, 1) >= a[u] 就是题目要求的部分
13. 假设root的影响，一直到了u，那么这个过程中的cost是多少？
14. 考虑一条线 10 -> 5 -> 4 -> 1
15. 先加5， 10 -> 10 -> 4 -> 1
16. 加6, 10 -> 10 -> 10 -> 1
17. +9, 10 -> 10 -> 10 -> 10
18. 换个例子  10 -> 3 -> 3 -> 3
19. +7, 10 -> 10 -> 10 -> 10
20. +7 + 7, 21 
21. 也就是说a[1]要一直找到一个>= a[1]的节点，在此之前，假设通过了x个节点，那么就是 x * a[1] - sum(...)
22. 但是变成树以后，还成立吗？感觉是成立的
23. 如果这样是这样的，要怎么编码？