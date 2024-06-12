You are given a tree of 𝑛
 vertices numbered from 1
 to 𝑛
. Initially, all vertices are colored white or black.

You are asked to perform 𝑞
 queries:

"u" — toggle the color of vertex 𝑢
 (if it was white, change it to black and vice versa).
After each query, you should answer whether all the black vertices form a chain. That is, there exist two black vertices such that the simple path between them passes through all the black vertices and only the black vertices. Specifically, if there is only one black vertex, they form a chain. If there are no black vertices, they do not form a chain.

### ideas
1. 考虑假设存在这样一个chain，其中的一个端点的depth肯定是最大的
2. 另外一个端点，就不好说了
3. 然后，那个depth最小的点，肯定也在这个path上
4. 考虑那个 heavy-light 分解，所形成的线段树，如果存在这样一个path，那么从那个最低点u开始，但那个最高点p，它们必然是连续的几段
5. 然后，在确认的过程中，先标记成white，然后再找到最低的点，继续处理一次
6. 最后再恢复
7. range update + range query
8. 