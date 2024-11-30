We had a really tough time generating tests for problem D. In order to prepare strong tests, we had to solve the following problem.

Given an undirected labeled tree consisting of 𝑛
 vertices, find a set of segments such that:

both endpoints of each segment are integers from 1
 to 2𝑛
, and each integer from 1
 to 2𝑛
 should appear as an endpoint of exactly one segment;
all segments are non-degenerate;
for each pair (𝑖,𝑗)
 such that 𝑖≠𝑗
, 𝑖∈[1,𝑛]
 and 𝑗∈[1,𝑛]
, the vertices 𝑖
 and 𝑗
 are connected with an edge if and only if the segments 𝑖
 and 𝑗
 intersect, but neither segment 𝑖
 is fully contained in segment 𝑗
, nor segment 𝑗
 is fully contained in segment 𝑖
.
Can you solve this problem too?


### ideas
1. u -> v 有一条边，说明它们要么在左端重叠，要么在右端重叠
2. 假设u是v的父节点，那么v的所有的子节点都在u的内部
3. 但是v在u的一端
4. 