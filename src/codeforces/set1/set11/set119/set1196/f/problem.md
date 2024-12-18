You are given a connected undirected weighted graph consisting of 𝑛
 vertices and 𝑚
 edges.

You need to print the 𝑘
-th smallest shortest path in this graph (paths from the vertex to itself are not counted, paths from 𝑖
 to 𝑗
 and from 𝑗
 to 𝑖
 are counted as one).

More formally, if 𝑑
 is the matrix of shortest paths, where 𝑑𝑖,𝑗
 is the length of the shortest path between vertices 𝑖
 and 𝑗
 (1≤𝑖<𝑗≤𝑛
), then you need to print the 𝑘
-th element in the sorted array consisting of all 𝑑𝑖,𝑗
, where 1≤𝑖<𝑗≤𝑛
.

Input
The first line of the input contains three integers 𝑛,𝑚
 and 𝑘
 (2≤𝑛≤2⋅105
, 𝑛−1≤𝑚≤min(𝑛(𝑛−1)2,2⋅105)
, 1≤𝑘≤min(𝑛(𝑛−1)2,400)
 — the number of vertices in the graph, the number of edges in the graph and the value of 𝑘
, correspondingly.

Then 𝑚
 lines follow, each containing three integers 𝑥
, 𝑦
 and 𝑤
 (1≤𝑥,𝑦≤𝑛
, 1≤𝑤≤109
, 𝑥≠𝑦
) denoting an edge between vertices 𝑥
 and 𝑦
 of weight 𝑤
.

It is guaranteed that the given graph is connected (there is a path between any pair of vertices), there are no self-loops (edges connecting the vertex with itself) and multiple edges (for each pair of vertices 𝑥
 and 𝑦
, there is at most one edge between this pair of vertices in the graph).

Output
Print one integer — the length of the 𝑘
-th smallest shortest path in the given graph (paths from the vertex to itself are not counted, paths from 𝑖
 to 𝑗
 and from 𝑗
 to 𝑖
 are counted as one).

 ### ideas
 1. 全部去计算肯定不行的
 2. 但是一个个的去计算，似乎也不知道怎么搞？
 3. bfs ?
 4. dp[i][j] 表示以i为终点的一条路径的第j短的path的长度 pair{length, from}
 5. 似乎是可行的。再想想