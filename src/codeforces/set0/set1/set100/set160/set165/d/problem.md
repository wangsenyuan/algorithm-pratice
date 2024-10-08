Let's define a non-oriented connected graph of n vertices and n - 1 edges as a beard, if all of its vertices except, perhaps, one, have the degree of 2 or 1 (that is, there exists no more than one vertex, whose degree is more than two). Let us remind you that the degree of a vertex is the number of edges that connect to it.

Let each edge be either black or white. Initially all edges are black.

You are given the description of the beard graph. Your task is to analyze requests of the following types:

paint the edge number i black. The edge number i is the edge that has this number in the description. It is guaranteed that by the moment of this request the i-th edge is white
paint the edge number i white. It is guaranteed that by the moment of this request the i-th edge is black
find the length of the shortest path going only along the black edges between vertices a and b or indicate that no such path exists between them (a path's length is the number of edges in it)
The vertices are numbered with integers from 1 to n, and the edges are numbered with integers from 1 to n - 1.

Input
The first line of the input contains an integer n (2 ≤ n ≤ 105) — the number of vertices in the graph. Next n - 1 lines contain edges described as the numbers of vertices vi, ui (1 ≤ vi, ui ≤ n, vi ≠ ui) connected by this edge. It is guaranteed that the given graph is connected and forms a beard graph, and has no self-loops or multiple edges.

The next line contains an integer m (1 ≤ m ≤ 3·105) — the number of requests. Next m lines contain requests in the following form: first a line contains an integer type, which takes values ​​from 1 to 3, and represents the request type.

If type = 1, then the current request is a request to paint the edge black. In this case, in addition to number type the line should contain integer id (1 ≤ id ≤ n - 1), which represents the number of the edge to paint.

If type = 2, then the current request is a request to paint the edge white, its form is similar to the previous request.

If type = 3, then the current request is a request to find the distance. In this case, in addition to type, the line should contain two integers a, b (1 ≤ a, b ≤ n, a can be equal to b) — the numbers of vertices, the distance between which must be found.

The numbers in all lines are separated by exactly one space. The edges are numbered in the order in which they are given in the input.

Output
For each request to "find the distance between vertices a and b" print the result. If there is no path going only along the black edges between vertices a and b, then print "-1" (without the quotes). Print the results in the order of receiving the requests, separate the numbers with spaces or line breaks.

### ideas
1. tree?
2. 涂成白色，相当于把一棵树分成两棵树
3. 涂成黑色，相当于连起来
4. 或者简单的来看，在u, v中间时候有白色的边
5. beard graph，有一个中心，其他的都是一条线