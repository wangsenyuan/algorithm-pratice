Timofey came to a famous summer school and found a tree on 𝑛
 vertices. A tree is a connected undirected graph without cycles.

Every vertex of this tree, except 𝑐0
, is colored white. The vertex 𝑐0
 is colored black.

Timofey wants to color all the vertices of this tree in black. To do this, he performs 𝑛−1
 operations. During the 𝑖
-th operation, he selects the vertex 𝑐𝑖
, which is currently white, and paints it black.

Let's call the positivity of tree the minimum distance between all pairs of different black vertices in it. The distance between the vertices 𝑣
 and 𝑢
 is the number of edges on the path from 𝑣
 to 𝑢
.

### ideas
1. 如果用c[i]和前面的所有的c[j]计算距离，那么复杂度是n^2；TLE。
2. 那么怎么加速这个过程呢？如果以c[i]为起点，是否可以很快的判断出，在多少距离内，有black的节点？
3. 不对。。。。
4. 还是没有找到题眼
5. 使用eular tour， 考虑两个子树，如果一个black在子树a中，另外一个在子树b中，那么它们的相差的距离是算不清楚的
6. 这是因为，a和b中间还有其他的子树。这里感觉似乎是高度是有关的
7. 假设这里有个主轴，任何一个黑色的点，都更新到主轴的距离，加上主轴上的下标
8. 那么对于任何一次查询，就可以找这个主轴上的最小值，然后新黑色节点到主轴的距离，并减去它的下标
9. 这样可以用heavy-light分解来处理。但感觉有点太复杂了
10. 假设从c0开始，然后一直往外找，找到第一个在c中的元素，那么c[i]对应的答案就是确定的（c[i]到c[0]的距离
11. 按找顺序处理，c[1] 和 c[0]找到它们的lcp节点，加入到这个集合中
12. 且更新lcp的距离为到(c[0], c[1])的最小值
13. 处理c[2], 如果它在外面，（产生了一个新的lcp节点）那么这个比较好处理
14. 如果它在lcp的内部，不好处理了。因为它有可能离c[0]近，也可能离c[1]近
15. 如果以c[0]为root，那么不考虑其他节点的影响的话， c[i]的答案就是它的深度
16. 如果以c[i]为root， 那么就是那些离c[i]最近的，且在它前面的节点
17. 假设c[i]的parent处理过来，然后移动到c[i]来处理，如果那些点在c[i]的parent那里是已经on的，它们的距离 + 1
18. 如果是在c[i]的子树范围内的，距离-1，然后找[0...i-1]范围内的最小值即可
19. 所以是范围查询+范围更新。 但还有问题，咋范围更新呢？
20. 