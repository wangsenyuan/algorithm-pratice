Bessie has received a birthday cake from her best friend Elsie, and it came in the form of a regular polygon with 𝑛
 sides. The vertices of the cake are numbered from 1
 to 𝑛
 clockwise. You and Bessie are going to choose some of those vertices to cut non-intersecting diagonals into the cake. In other words, the endpoints of the diagonals must be part of the chosen vertices.

Bessie would only like to give out pieces of cake which result in a triangle to keep consistency. The size of the pieces doesn't matter, and the whole cake does not have to be separated into all triangles (other shapes are allowed in the cake, but those will not be counted).

Bessie has already chosen 𝑥
 of those vertices that can be used to form diagonals. She wants you to choose no more than 𝑦
 other vertices such that the number of triangular pieces of cake she can give out is maximized.

What is the maximum number of triangular pieces of cake Bessie can give out?

### ideas
1. 一段连续选中的节点，假设有x个，能形成多少个三角形？
2. 假设有两段，它们是最靠近的（在环上）长度分别是a，b，那么一共可以形成 (a - 1) + (b - 1) 个三角形，且如果它们中间只隔了1个未选中的点，还可以额外 + 1, 如果在另外一段，也正好距离1，再额外+1
3. 所以从上面的公式可以看出来，分割出距离为1的段数越多越好。
4. 选中间隔1，且被选中的点，切下这个角后，n--, result++
5. 先把所有绿色的点，这样处理后，得到一个结果result
6. 那么剩余的，就是一些连续的绿色节点，且距离超过1的段
7. 然后对于每一段，在它距离2的地方，添加黄色的点，从而不断增加新的result，并合并点
8. 这样子知道新的黄色的点把两段绿色的连起来了，再进行切，剩下它们的两个端点变成相邻的
9. 这个过程好复杂。还涉及到删除点后，相邻节点的更新
10. 而且有点不清楚。
11. 连续绿色的节点，貌似只能切出来m-2个三角形，
12. 然后剩下的都是1个绿色节点，2个绿色节点，这样的分段，或者没有绿色节点
13. 