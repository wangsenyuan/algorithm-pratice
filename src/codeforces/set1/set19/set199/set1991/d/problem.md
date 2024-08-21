You are given an undirected graph with n
 vertices, numbered from 1
 to n
. There is an edge between vertices u
 and v
 if and only if u⊕v
 is a prime number, where ⊕
 denotes the bitwise XOR operator.

Color all vertices of the graph using the minimum number of colors, such that no two vertices directly connected by an edge have the same color.

Input
Each test contains multiple test cases. The first line contains the number of test cases t
 (1≤t≤500
). The description of test cases follows.

The only line contains a single integer n
 (1≤n≤2⋅105
) — the number of vertices in the graph.

It is guaranteed that the sum of n
 over all test cases does not exceed 2⋅105
.

Output
For each test case, output two lines.

The first line should contain a single integer k
 (1≤k≤n
) — the minimum number of colors required.

The second line should contain n
 integers c1,c2,…,cn
 (1≤ci≤k
) — the color of each vertex.

If there are multiple solutions, output any of them.

### ideas

### solutions
For n≥6
, the minimum number of colors is always 4
.

Proof

First, we can show that the number of colors cannot be less than 4
. This is because vertices 1
, 3
, 4
, and 6
 form a clique, meaning they are all connected, so they must have different colors.

Next, we can provide a construction where the number of colors is 4
. For the i
-th vertex, assign the color imod4+1
. This ensures that any two vertices of the same color have a difference that is a multiple of 4
, so their XOR is a multiple of 4
, which is not a prime number.

For n<6
, the example provides the coloring for all cases:

n=1
: A valid coloring is [1]
.
n=2
: A valid coloring is [1,2]
.
n=3
: A valid coloring is [1,2,2]
.
n=4
: A valid coloring is [1,2,2,3]
.
n=5
: A valid coloring is [1,2,2,3,3]
.