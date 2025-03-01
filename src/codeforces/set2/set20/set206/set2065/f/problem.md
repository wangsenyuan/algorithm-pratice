Let's define the majority of a sequence of 𝑘
 elements as the unique value that appears strictly more than ⌊𝑘2⌋
 times. If such a value does not exist, then the sequence does not have a majority. For example, the sequence [1,3,2,3,3]
 has a majority 3
 because it appears 3>⌊52⌋=2
 times, but [1,2,3,4,5]
 and [1,3,2,3,4]
 do not have a majority.

Skibidus found a tree∗
 of 𝑛
 vertices and an array 𝑎
 of length 𝑛
. Vertex 𝑖
 has the value 𝑎𝑖
 written on it, where 𝑎𝑖
 is an integer in the range [1,𝑛]
.

For each 𝑖
 from 1
 to 𝑛
, please determine if there exists a non-trivial simple path†
 such that 𝑖
 is the majority of the sequence of integers written on the vertices that form the path.

∗
A tree is a connected graph without cycles.

†
A sequence of vertices 𝑣1,𝑣2,...,𝑣𝑚
 (𝑚≥2
) forms a non-trivial simple path if 𝑣𝑖
 and 𝑣𝑖+1
 are connected by an edge for all 1≤𝑖≤𝑚−1
 and all 𝑣𝑖
 are pairwise distinct. Note that the path must consist of at least 2
 vertices.

Input
Each test contains multiple test cases. The first line contains the number of test cases 𝑡
 (1≤𝑡≤104
). The description of the test cases follows.

The first line of each test case contains a single integer 𝑛
 (2≤𝑛≤5⋅105
)  — the number of vertices.

The second line of each test case contains 𝑎1,𝑎2,…,𝑎𝑛
 (1≤𝑎𝑖≤𝑛
)  — the integers written on the vertices.

Each of the next 𝑛−1
 lines contains two integers 𝑢𝑖
 and 𝑣𝑖
, denoting the two vertices connected by an edge (1≤𝑢𝑖,𝑣𝑖≤𝑛
, 𝑢𝑖≠𝑣𝑖
).

It is guaranteed that the given edges form a tree.

It is guaranteed that the sum of 𝑛
 over all test cases does not exceed 5⋅105
.