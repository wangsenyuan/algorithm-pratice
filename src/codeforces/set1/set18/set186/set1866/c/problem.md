Pak Chanek has a directed acyclic graph (a directed graph that does not have any cycles) containing 𝑁
vertices. Vertex 𝑖
has 𝑆𝑖
edges directed away from that vertex. The 𝑗
-th edge of vertex 𝑖
that is directed away from it, is directed towards vertex 𝐿𝑖,𝑗
and has an integer 𝑊𝑖,𝑗
(0≤𝑊𝑖,𝑗≤1
). Another information about the graph is that the graph is shaped in such a way such that each vertex can be reached
from vertex 1
via zero or more directed edges.

Pak Chanek has an array 𝑍
that is initially empty.

Pak Chanek defines the function dfs as follows:

```

// dfs from vertex i
void dfs(int i) {
    // iterate each edge of vertex i that is directed away from it
    for(int j = 1; j <= S[i]; j++) {
        Z.push_back(W[i][j]); // add the integer in the edge to the end of Z
        dfs(L[i][j]); // recurse to the next vertex
    }
}

```

Note that the function does not keep track of which vertices have been visited, so each vertex can be processed more
than once.

Let's say Pak Chanek does dfs(1) once. After that, Pak Chanek will get an array 𝑍
containing some elements 0
or 1
. Define an inversion in array 𝑍
as a pair of indices (𝑥,𝑦)
(𝑥<𝑦
) such that 𝑍𝑥>𝑍𝑦
. How many different inversions in 𝑍
are there if Pak Chanek does dfs(1) once? Since the answer can be very big, output the answer modulo 998244353
.

### thoughts

1. 在访问到1的时候，需要直到后面有多少个0
2. 完全没有想法呐～
3. 如果在访问到一条边的时候，它前面的1的个数，如果能够确定的话
4. 那么这条边的贡献也是确定的
5. 但是问题在于，节点v访问完以后，可能会被第二次、第三次访问
6. 在后续的访问中，节点v中的1的个数，也会影响到后续的结果
7. 但是如果只考虑连接v的那条边，假设它为0，在v的前面部分有1的个数是p1
8. 且假设它被重复访问了x次
8. 那么 p1 + (p1 + v1) + (p1 + 2 * v1) + ... + (p1 + (x - 1) * v1)
9. 错了～

### solution

Note that during the entire process, each time we do dfs(x) for the same value of 𝑥
, the sequence of values appended to the end of 𝑍
is always the same. So, for each vertex 𝑥
, we want to obtain some properties about the sequence of values that will be appended if we do dfs(x). Since we want to
calculate the number of inversions in a sequence of 0
and 1
elements, for each 𝑥
, we want to calculate:

𝑓0[𝑥]
: the number of 0
elements appended.
𝑓1[𝑥]
: the number of 1
elements appended.
inv[𝑥]
: the number of inversions completely within the sequence appended.
Then, we can solve the problem using dynamic programming on directed acyclic graph. For a vertex 𝑥
, we iterate each vertex 𝑦
that it is directed towards (following the order given in the input). If the integer in the current edge is 𝑤
, then we temporarily add 𝑓𝑤[𝑦]
by 1
and if 𝑤=1
, then we temporarily add inv[𝑦]
by 𝑓0[𝑦]
. Then, we do the following transitions:

Add inv[𝑥]
by inv[𝑦]+𝑓1(𝑥)×𝑓0(𝑦)
.
Add 𝑓0[𝑥]
by 𝑓0[𝑦]
.
Add 𝑓1[𝑥]
by 𝑓1[𝑦]
.
The answer is inv[1]
.

Time complexity: 𝑂(𝑁+∑𝑆)