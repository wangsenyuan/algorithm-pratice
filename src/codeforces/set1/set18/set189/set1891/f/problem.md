You are given a rooted tree with the root at vertex 1
, initially consisting of a single vertex. Each vertex has a numerical value, initially set to 0
. There are also 𝑞
queries of two types:

The first type: add a child vertex with the number 𝑠𝑧+1
to vertex 𝑣
, where 𝑠𝑧
is the current size of the tree. The numerical value of the new vertex will be 0
.
The second type: add 𝑥
to the numerical values of all vertices in the subtree of vertex 𝑣
.
After all queries, output the numerical value of all of the vertices in the final tree.

### thoughts

1. 先按照操作1构造树
2. 然后构造eular序列
3. 然后从后往前处理操作，同时在in[u] + x， 在out[u]处进行 -x
4. 然后求in[u]的前缀和即可