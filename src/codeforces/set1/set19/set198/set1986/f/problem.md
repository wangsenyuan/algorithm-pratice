You are given a connected undirected graph, the vertices of which are numbered with integers from 1
 to 𝑛
. Your task is to minimize the number of pairs of vertices 1≤𝑢<𝑣≤𝑛
 between which there exists a path in this graph. To achieve this, you can remove exactly one edge from the graph.

Find the smallest number of pairs of vertices!

Input
Each test consists of several sets of input data. The first line contains a single integer 𝑡
 (1≤𝑡≤104
) — the number of sets of input data. Then follows their description.

The first line of each set of input data contains two integers 𝑛
 and 𝑚
 (2≤𝑛≤105
, 𝑛−1≤𝑚≤min(105,𝑛⋅(𝑛−1)2)
) — the number of vertices in the graph and the number of edges.

Each of the next 𝑚
 lines contains two integers 𝑢
 and 𝑣
 (1≤𝑢,𝑣≤𝑛,𝑢≠𝑣
), indicating that there is an undirected edge in the graph between vertices 𝑢
 and 𝑣
.

It is guaranteed that the given graph is connected and without multiple edges.

It is guaranteed that the sum of 𝑛
 and the sum of 𝑚
 over all sets of input data does not exceed 2⋅105
.


### ideas
1. 如果图是强连通的，那么删掉任何边，也不会有效果
2. 所以，只能删掉那些bridge（这个有标准算法）
3. 找到这些bridge就可以了。看看它的贡献是多少，即可
4. 这个bridge，貌似就是 low[v] > dis[u]的那些
5. 这里还有个问题，就是要算这个bridge分割后的component的size
6. 好像也没有问题