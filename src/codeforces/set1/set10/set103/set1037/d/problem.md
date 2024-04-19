The BFS algorithm is defined as follows.

Consider an undirected graph with vertices numbered from 1
to 𝑛
. Initialize 𝑞
as a new queue containing only vertex 1
, mark the vertex 1
as used.
Extract a vertex 𝑣
from the head of the queue 𝑞
.
Print the index of vertex 𝑣
.
Iterate in arbitrary order through all such vertices 𝑢
that 𝑢
is a neighbor of 𝑣
and is not marked yet as used. Mark the vertex 𝑢
as used and insert it into the tail of the queue 𝑞
.
If the queue is not empty, continue from step 2.
Otherwise finish.
Since the order of choosing neighbors of each vertex can vary, it turns out that there may be multiple sequences which
BFS can print.

In this problem you need to check whether a given sequence corresponds to some valid BFS traversal of the given tree
starting from vertex 1
. The tree is an undirected graph, such that there is exactly one simple path between any two vertices.

### ideas

1. 同一层的顺序不定，但不同层的不能交换
2. 还有个特点是，节点u的子节点，顺序不定，但肯定是连续的
3. 如果节点u的两个子节点v1, v2， 如果v1在v2的前面，那么v1的子节点同一层的肯定也在v2的前面