378QAQ has a tree with 𝑛
 vertices. Initially, all vertices are white.

There are two chess pieces called 𝑃𝐴
 and 𝑃𝐵
 on the tree. 𝑃𝐴
 and 𝑃𝐵
 are initially located on vertices 𝑎
 and 𝑏
 respectively. In one step, 378QAQ will do the following in order:

Move 𝑃𝐴
 to a neighboring vertex. If the target vertex is white, this vertex will be painted red.
Move 𝑃𝐵
 to a neighboring vertex. If the target vertex is colored in red, this vertex will be painted blue.
Initially, the vertex 𝑎
 is painted red. If 𝑎=𝑏
, the vertex 𝑎
 is painted blue instead. Note that both the chess pieces must be moved in each step. Two pieces can be on the same vertex at any given time.

378QAQ wants to know the minimum number of steps to paint all vertices blue.

### ideas
1. bfs on vertics with status
2. 不大对
3. 对a的移动，简单来说，是不受b的限制的（真的不受影响吗？）
4. 如果存在一个状态，b跟在a的后面，那么这时，没有多余的步骤
5. 在此之前，b走的步骤，有可能是浪费的
6. 所以，在(a, b)相遇以前，b的步骤都是浪费的，（a的也可以认为是浪费的）
7. 应该考虑它们相遇后的情况
8. 如果它们之间的距离（边的个数）是偶数，那么它们可以到到一个中点（u),从那里开始，blue的颜色开始考虑
9. 如果是奇数，假设中间是（u, v), 那么b到到u时，a到到v（u时blue，但是v时red）
10. 就是从v开始算 + 1