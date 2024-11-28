You are given a connected undirected graph consisting of 𝑛
 vertices and 𝑚
 edges. There are no self-loops or multiple edges in the given graph.

You have to direct its edges in such a way that the obtained directed graph does not contain any paths of length two or greater (where the length of path is denoted as the number of traversed edges).

Input
The first line contains two integer numbers 𝑛
 and 𝑚
 (2≤𝑛≤2⋅105
, 𝑛−1≤𝑚≤2⋅105
) — the number of vertices and edges, respectively.

The following 𝑚
 lines contain edges: edge 𝑖
 is given as a pair of vertices 𝑢𝑖
, 𝑣𝑖
 (1≤𝑢𝑖,𝑣𝑖≤𝑛
, 𝑢𝑖≠𝑣𝑖
). There are no multiple edges in the given graph, i. e. for each pair (𝑢𝑖,𝑣𝑖
) there are no other pairs (𝑢𝑖,𝑣𝑖
) and (𝑣𝑖,𝑢𝑖
) in the list of edges. It is also guaranteed that the given graph is connected (there is a path between any pair of vertex in the given graph).

Output
If it is impossible to direct edges of the given graph in such a way that the obtained directed graph does not contain paths of length at least two, print "NO" in the first line.

Otherwise print "YES" in the first line, and then print any suitable orientation of edges: a binary string (the string consisting only of '0' and '1') of length 𝑚
. The 𝑖
-th element of this string should be '0' if the 𝑖
-th edge of the graph should be directed from 𝑢𝑖
 to 𝑣𝑖
, and '1' otherwise. Edges are numbered in the order they are given in the input.

### ideas
1. 对于节点u的两条边来说，要们他们都指向u，要么都远离u
2. 否则的话，就会出现长度为2的path
3. 且两个相邻的节点，u, v如果u是边远离的那些，那么v就是边进入的节点
4. 所以可以分成二部图