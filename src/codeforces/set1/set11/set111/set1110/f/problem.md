Let's define the Eulerian traversal of a tree (a connected undirected graph without cycles) as follows: consider a
depth-first search algorithm which traverses vertices of the tree and enumerates them in the order of visiting (only the
first visit of each vertex counts). This function starts from the vertex number 1
and then recursively runs from all vertices which are connected with an edge with the current vertex and are not yet
visited in increasing numbers order. Formally, you can describe this function using the following pseudocode:

next_id = 1
id = array of length n filled with -1
visited = array of length n filled with false

function dfs(v):
visited[v] = true
id[v] = next_id
next_id += 1
for to in neighbors of v in increasing order:
if not visited[to]:
dfs(to)
You are given a weighted tree, the vertices of which were enumerated with integers from 1
to 𝑛
using the algorithm described above.

A leaf is a vertex of the tree which is connected with only one other vertex. In the tree given to you, the vertex 1
is not a leaf. The distance between two vertices in the tree is the sum of weights of the edges on the simple path
between them.

You have to answer 𝑞
queries of the following type: given integers 𝑣
, 𝑙
and 𝑟
, find the shortest distance from vertex 𝑣
to one of the leaves with indices from 𝑙
to 𝑟
inclusive.

### ideas

1. 假设u是包含l...r的子树的根节点, dp[u] 等于u到其中某个叶子节点的最短距离
2. 如果v在u的外面，那么ans = dist(u, v) + dp[u]
3. 如果v在u的内部， 那么ans = min(dp[v], dist(v, p) + dp[p]) (p是v的父节点)
4. u = v, dp[u]
5. 不大对，
6. 在一个欧拉tour中，一颗子树内部的节点的编号是连续的
7. l...i 是在一颗子树中，i+1...j在一颗子树内。。。j...r是一颗子树

### solution

Let's answer the queries offline: for each vertex we'll remember all queries for it.

Let's make vertex 1
root and find distances from it to all leaves using DFS. Now for answering queries for vertex 1
we should simply answer some range minimum queries, so we'll build a segment tree for it.

For answering queries for other vertices we will make segment trees with distances to leaves (like we did for handling
queries for vertex 1
) for all vertices. We will traverse vertices of a tree using DFS and maintain actual distances to leaves in segment
tree. Suppose, DFS went through edge (𝑢,𝑣)
with weight 𝑤
where 𝑢
is an ancestor of 𝑣
. Distances to leaves in subtree of 𝑣
will decrease by 𝑤
and distances to other vertices will increase by 𝑤
. Since set of leaves in subtree of some vertex is a subsegment of vertices written in euler traversal order, we should
simply make range additions/subtractions on segment tree to maintain distances to leaves when passing through an edge.

This way we got a solution with 𝑂(𝑛log𝑛+𝑞log𝑛)
complexity.