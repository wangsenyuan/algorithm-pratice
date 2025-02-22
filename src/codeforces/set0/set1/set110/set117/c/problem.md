A tournament is a directed graph without self-loops in which every pair of vertexes is connected by exactly one directed edge. That is, for any two vertexes u and v (u ≠ v) exists either an edge going from u to v, or an edge from v to u.

You are given a tournament consisting of n vertexes. Your task is to find there a cycle of length three.

Input
The first line contains an integer n (1 ≤ n ≤ 5000). Next n lines contain the adjacency matrix A of the graph (without spaces). Ai, j = 1 if the graph has an edge going from vertex i to vertex j, otherwise Ai, j = 0. Ai, j stands for the j-th character in the i-th line.

It is guaranteed that the given graph is a tournament, that is, Ai, i = 0, Ai, j ≠ Aj, i (1 ≤ i, j ≤ n, i ≠ j).

Output
Print three distinct vertexes of the graph a1, a2, a3 (1 ≤ ai ≤ n), such that Aa1, a2 = Aa2, a3 = Aa3, a1 = 1, or "-1", if a cycle whose length equals three does not exist.

If there are several solutions, print any of them.

### ideas
1. 如果从j出发，知道其他所有节点到j的距离（不超过2的节点）
2. 那么如果i -> j 是一条边，且j存在一个节点k, j -> k -> i，那么就找到了

### solution

Approval. If the tournament has at least one cycle, then there exist a cycle of length three.

A constructive proof. Find any cycle in the tournament by using any standard algorithm, such as depth-first search. If there is no cycle, then output -1, else choose any three consecutive vertices of the cycle v1 v2 v3 (Av1, v2 = Av2, v3 = 1). Since given graph is a tournament, then there is an edge (v3, v1), or there is an edge (v1, v3). The first of these two cases, we find immediately the cycle of length three of the vertices v1 v2 v3,  the second, we can reduce the length of the loop (erase vertex v2). 

We can reduce the length of the cycle until we find a cycle of length three.
