Let's define a forest as a non-directed acyclic graph (also without loops and parallel edges). One day Misha played with the forest consisting of n vertices. For each vertex v from 0 to n - 1 he wrote down two integers, degreev and sv, were the first integer is the number of vertices adjacent to vertex v, and the second integer is the XOR sum of the numbers of vertices adjacent to v (if there were no adjacent vertices, he wrote down 0).

Next day Misha couldn't remember what graph he initially had. Misha has values degreev and sv left, though. Help him find the number of edges and the edges of the initial graph. It is guaranteed that there exists a forest that corresponds to the numbers written by Misha.

#### ideas
1. deg(u) = 1 这个是叶子节点，然后s(u) = parent （因为只有一个节点）
2. 这样就可以了