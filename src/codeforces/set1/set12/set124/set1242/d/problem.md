Ujan has a lot of useless stuff in his drawers, a considerable part of which are his math notebooks: it is time to sort them out. This time he found an old dusty graph theory notebook with a description of a graph.

It is an undirected weighted graph on 𝑛
 vertices. It is a complete graph: each pair of vertices is connected by an edge. The weight of each edge is either 0
 or 1
; exactly 𝑚
 edges have weight 1
, and all others have weight 0
.

Since Ujan doesn't really want to organize his notes, he decided to find the weight of the minimum spanning tree of the graph. (The weight of a spanning tree is the sum of all its edges.) Can you find the answer for Ujan so he stops procrastinating?

Input
The first line of the input contains two integers 𝑛
 and 𝑚
 (1≤𝑛≤105
, 0≤𝑚≤min(𝑛(𝑛−1)2,105)
), the number of vertices and the number of edges of weight 1
 in the graph.

The 𝑖
-th of the next 𝑚
 lines contains two integers 𝑎𝑖
 and 𝑏𝑖
 (1≤𝑎𝑖,𝑏𝑖≤𝑛
, 𝑎𝑖≠𝑏𝑖
), the endpoints of the 𝑖
-th edge of weight 1
.

It is guaranteed that no edge appears twice in the input.

### ideas
1. 很有意思的一个题目。
2. 显然所有的0边，可以被用来形成component（当然没法全部迭代）然后component 数量 - 1 就是答案
3. 但问题是如何形成component？
4. 考虑2个集合(1, 2, 3) (4, 5)
5. 那么必须存在1-4, 2-4,3-4, 1-5,... 共6条边存在
6. 假设一个component的数量是sz, 那么它必须贡献 sz * (n - sz)条边
7. 也就是说如果1的边的数量 < n - 1, 答案= 0
8. 现在找那个deg 最大的节点, 如果它的 deg = x (那么表示，它的component的sz = n - x)
9. 如果 (n - x) * x > 给定边的数量, 就没法再分割出一个component
10. (n - x) * x <= ..数量, 那么就可以把这个compoent给分离出来
11. 且在同一个component中的deg肯定是一样的
12. 但是deg相同的不一定在一个component中
13. 