Kavi has 2𝑛
 points lying on the 𝑂𝑋
 axis, 𝑖
-th of which is located at 𝑥=𝑖
.

Kavi considers all ways to split these 2𝑛
 points into 𝑛
 pairs. Among those, he is interested in good pairings, which are defined as follows:

Consider 𝑛
 segments with ends at the points in correspondent pairs. The pairing is called good, if for every 2
 different segments 𝐴
 and 𝐵
 among those, at least one of the following holds:

One of the segments 𝐴
 and 𝐵
 lies completely inside the other.
𝐴
 and 𝐵
 have the same length.
Consider the following example:


### ideas
1. 例子3把各种情况都列举了
2. 如果存在k中不同的长度
3. 那么除了最后一种，其他的k-1种，每个必须有一个，且嵌套在一起，且它们位于中心
4. 不对
5. 要分3类，完全在内层的，完全在外层的，（它们只有一个且相互嵌套），在中间层的，被所有外层包括，且包括了所有内层，且长度相同
6. 中间层的，因为相等，会有很多种选择
7. 假设没有内层是，中间层的长度为 2 * d + 1时
8. 这里一共有 2 * d + 2 个节点
9. 其中和最左端节点匹配的节点，决定了其他所有节点的配对
10. 所以，共有 2 * d 中情况, 比如当d = 1是，共有3种情况
11. 不包括最后一种情况 (n-d...n+d)连接在一起
12. f(d) = f(d-1) + 2 * d
13. res = f(n-1) + 1 (最后一种情况是把两个端点连起来)
14. f(0) = 0
15. f(1) = 0 + 2 * 1 
16. f(2) = 2 + 4 = 6 (但是还是要 + 1， 变成7了)
17. f(d) 表示当长度是 2 * d + 1 时，符合条件2的计数
18. 如果d是奇数， f(d) = f(d-1) + f(d/2) + d - 1
19. 如果d是偶数  f(d) =  f(d-1) + d - 1
20. 