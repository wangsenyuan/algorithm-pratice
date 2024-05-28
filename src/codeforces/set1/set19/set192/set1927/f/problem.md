Given an undirected weighted graph with 𝑛
 vertices and 𝑚
 edges. There is at most one edge between each pair of vertices in the graph, and the graph does not contain loops (edges from a vertex to itself). The graph is not necessarily connected.

A cycle in the graph is called simple if it doesn't pass through the same vertex twice and doesn't contain the same edge twice.

Find any simple cycle in this graph in which the weight of the lightest edge is minimal.

Input
The first line of the input contains a single integer 𝑡
 (1≤𝑡≤104
) — the number of test cases. Then follow the descriptions of the test cases.

The first line of each test case contains two integers 𝑛
 and 𝑚
 (3≤𝑛≤𝑚≤min(𝑛⋅(𝑛−1)2,2⋅105)
) — the size of the graph and the number of edges.

The next 𝑚
 lines of the test case contain three integers 𝑢
, 𝑣
, and 𝑤
 (1≤𝑢,𝑣≤𝑛
, 𝑢≠𝑣
, 1≤𝑤≤106
) — vertices 𝑢
 and 𝑣
 are connected by an edge of weight 𝑤
.

It is guaranteed that there is at most one edge between each pair of vertices. Note that under the given constraints, there is always at least one simple cycle in the graph.

It is guaranteed that the sum of the values of 𝑚
 for all test cases does not exceed 2⋅105
.

Output
For each test case, output a pair of numbers 𝑏
 and 𝑘
, where:

𝑏
 — the minimum weight of the edge in the found cycle,
𝑘
 — the number of vertices in the found cycle.
On the next line, output 𝑘
 numbers from 1
 to 𝑛
  — the vertices of the cycle in traversal order.

Note that the answer always exists, as under the given constraints, there is always at least one simple cycle in the graph.

### ideas
1. b很容易找出来，用kruskal算法，从大往小的添加edge，如果出现了一个loop时，这个edge就是一个答案
2. 但是在答案出现的时候，怎么找到k（这k个节点）？
3. 在出现环的时候，记录两个端点，然后用dfs，从端点u找到端点v（但是要排除那个最后的edge）