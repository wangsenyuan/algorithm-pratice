Sasha wants to take a walk with his girlfriend in the city. The city consists of 𝑛
 intersections, numbered from 1
 to 𝑛
. Some of them are connected by roads, and from any intersection, there is exactly one simple path†
 to any other intersection. In other words, the intersections and the roads between them form a tree.

Some of the intersections are considered dangerous. Since it is unsafe to walk alone in the city, Sasha does not want to visit three or more dangerous intersections during the walk.

Sasha calls a set of intersections good if the following condition is satisfied:

If in the city only the intersections contained in this set are dangerous, then any simple path in the city contains no more than two dangerous intersections.
However, Sasha does not know which intersections are dangerous, so he is interested in the number of different good sets of intersections in the city. Since this number can be very large, output it modulo 998244353
.

†
A simple path is a path that passes through each intersection at most once.

### ideas
1. 任何长度超过2的path，它们的(至少3个)节点不能同时在set中
2. 正难则反？
3. 如果选中u，那么任何它子树中的和它外边的链接，通过它的两个节点都是不行的
4. (pow(2, n - cnt[u])  - 1)* (pow(2, cnt[u]) - 1)
5. u中的也有同样的办法计算。
6. 但是会不会有重复的？
7. 会有的，考虑u, v， 会在计算v的时候，把u的贡献也统计进去， 在计算u的时候也会
8. 那就计算good的？
9. 以u为一个端点的，如果是u中的节点，那么只能选择一个，这样的有cnt[u] - 1 中
10. 每个可选/可不选，总数 = pow(2, cnt[u] - 1)中
11. 其中bad的 = pow(2, ....) * .... 这是是要排除掉的
12. got it
13. 