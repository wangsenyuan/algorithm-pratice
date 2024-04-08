Vasya has a graph containing both directed (oriented) and undirected (non-oriented) edges. There can be multiple edges
between a pair of vertices.

Vasya has picked a vertex s from the graph. Now Vasya wants to create two separate plans:

to orient each undirected edge in one of two possible directions to maximize number of vertices reachable from vertex s;
to orient each undirected edge in one of two possible directions to minimize number of vertices reachable from vertex s.
In each of two plans each undirected edge must become directed. For an edge chosen directions can differ in two plans.

Help Vasya find the plans.

Input
The first line contains three integers n, m and s (2 ≤ n ≤ 3·105, 1 ≤ m ≤ 3·105, 1 ≤ s ≤ n) — number of vertices and
edges in the graph, and the vertex Vasya has picked.

The following m lines contain information about the graph edges. Each line contains three integers ti, ui and vi (1 ≤
ti ≤ 2, 1 ≤ ui, vi ≤ n, ui ≠ vi) — edge type and vertices connected by the edge. If ti = 1 then the edge is directed and
goes from the vertex ui to the vertex vi. If ti = 2 then the edge is undirected and it connects the vertices ui and vi.

It is guaranteed that there is at least one undirected edge in the graph.

Output
The first two lines should describe the plan which maximizes the number of reachable vertices. The lines three and four
should describe the plan which minimizes the number of reachable vertices.

A description of each plan should start with a line containing the number of reachable vertices. The second line of a
plan should consist of f symbols '+' and '-', where f is the number of undirected edges in the initial graph. Print '+'
as the j-th symbol of the string if the j-th undirected edge (u, v) from the input should be oriented from u to v.
Print '-' to signify the opposite direction (from v to u). Consider undirected edges to be numbered in the same order
they are given in the input.

If there are multiple solutions, print any of them.

### ideas

1. greedy?
2. 从s出发，一直沿着directed edge；然后如果要最大化，那么就将所有当前双向的边，变成单向，然后继续访问
3. 感觉好像很简单么～
4. 