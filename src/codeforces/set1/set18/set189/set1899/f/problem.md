Alex's birthday is coming up, and Timofey would like to gift him a tree of 𝑛
vertices. However, Alex is a very moody boy. Every day for 𝑞
days, he will choose an integer, denoted by the integer chosen on the 𝑖
-th day by 𝑑𝑖
. If on the 𝑖
-th day there are not two leaves in the tree at a distance exactly 𝑑𝑖
, Alex will be disappointed.

Timofey decides to gift Alex a designer so that he can change his tree as he wants. Timofey knows that Alex is also
lazy (a disaster, not a human being), so at the beginning of every day, he can perform no more than one operation of the
following kind:

Choose vertices 𝑢
, 𝑣1
, and 𝑣2
such that there is an edge between 𝑢
and 𝑣1
and no edge between 𝑢
and 𝑣2
. Then remove the edge between 𝑢
and 𝑣1
and add an edge between 𝑢
and 𝑣2
. This operation cannot be performed if the graph is no longer a tree after it.
Somehow Timofey managed to find out all the 𝑑𝑖
. After that, he had another brilliant idea — just in case, make an instruction manual for the set, one that Alex
wouldn't be disappointed.

Timofey is not as lazy as Alex, but when he saw the integer 𝑛
, he quickly lost the desire to develop the instruction and the original tree, so he assigned this task to you. It can
be shown that a tree and a sequence of operations satisfying the described conditions always exist.

### thoughts

1. 考虑固定直径中的一个点作为其中的一个端点，假设直径连接的路径是有 v1, v2, v3... vk组成
2. 考虑其他叶子节点到这个路径节点上的距离位x1, x2, x3.... xm， 共有(m + 2个叶子节点)
3. 组成一个十字型，当需要一个di的时候，从如果存在一条比较长的臂 > di,然后保留di长度的边，并将多余的接到另外一个臂
4. 这里唯一有问题的在于，如果当前是一条直线，下一个需要的是
5. 所以应该是个人字形