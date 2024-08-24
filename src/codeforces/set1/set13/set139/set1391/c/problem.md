A permutation of length 𝑛
 is an array consisting of 𝑛
 distinct integers from 1
 to 𝑛
 in arbitrary order. For example, [2,3,1,5,4]
 is a permutation, but [1,2,2]
 is not a permutation (2
 appears twice in the array) and [1,3,4]
 is also not a permutation (𝑛=3
 but there is 4
 in the array).

Consider a permutation 𝑝
 of length 𝑛
, we build a graph of size 𝑛
 using it as follows:

For every 1≤𝑖≤𝑛
, find the largest 𝑗
 such that 1≤𝑗<𝑖
 and 𝑝𝑗>𝑝𝑖
, and add an undirected edge between node 𝑖
 and node 𝑗
For every 1≤𝑖≤𝑛
, find the smallest 𝑗
 such that 𝑖<𝑗≤𝑛
 and 𝑝𝑗>𝑝𝑖
, and add an undirected edge between node 𝑖
 and node 𝑗
In cases where no such 𝑗
 exists, we make no edges. Also, note that we make edges between the corresponding indices, not the values at those indices.

For clarity, consider as an example 𝑛=4
, and 𝑝=[3,1,4,2]
; here, the edges of the graph are (1,3),(2,1),(2,3),(4,3)
.

A permutation 𝑝
 is cyclic if the graph built using 𝑝
 has at least one simple cycle.

Given 𝑛
, find the number of cyclic permutations of length 𝑛
. Since the number may be very large, output it modulo 109+7
.

Please refer to the Notes section for the formal definition of a simple cycle

There are 16
 cyclic permutations for 𝑛=4
. [4,2,1,3]
 is one such permutation, having a cycle of length four: 4→3→2→1→4
.

Nodes 𝑣1
, 𝑣2
, …
, 𝑣𝑘
 form a simple cycle if the following conditions hold:

𝑘≥3
.
𝑣𝑖≠𝑣𝑗
 for any pair of indices 𝑖
 and 𝑗
. (1≤𝑖<𝑗≤𝑘
)
𝑣𝑖
 and 𝑣𝑖+1
 share an edge for all 𝑖
 (1≤𝑖<𝑘
), and 𝑣1
 and 𝑣𝑘
 share an edge.


 ### ideas
 1. 规则1说的是，i和左边最近的j, p[j] > p[i]之间有边，反过来说，就是对于j来说，是和右边比p[j]小的i有边
 2. 规则2说的是，i和右边最近的j, p[j] > p[i]之间有边
 3. 两个规则综合起来就是对于l来说，假设右边比它大的最近的位置是r，那么l, r有一条边，且l到所有(l...r)之间的数又一条边
 4. 只要它们中间有一个数，就肯定能形成环，假设（l, r)中间最大的数的下标是u, 那么根据规则1， p[l] > p[u],它们之间有一条边
 5. 且p[u] < p[l] < p[r] 那么根据规则2， u和r中间存在一条边
 6. 所以只要序列不是递增、递减就肯定存在环
 7. P(n) - 2? 
 8. P(4) = 4 * 3 * 2 * 1 - 2 = 22?
 9. 不对，要考虑4的位置，如果4切好把序列分成两段，且这两段是递增、递降也不存在环
1.  P(4) - C(3, 0) - C(3, 1) - C(3, 2) - C(3, 3) = 24 - 1 - 3 - 3 - 1 = 16 
2.  P(n) - pow(2, n - 1)