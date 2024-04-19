You are given 𝑛
segments on a number line; each endpoint of every segment has integer coordinates. Some segments can degenerate to
points. Segments can intersect with each other, be nested in each other or even coincide.

The intersection of a sequence of segments is such a maximal set of points (not necesserily having integer coordinates)
that each point lies within every segment from the sequence. If the resulting set isn't empty, then it always forms some
continuous segment. The length of the intersection is the length of the resulting segment or 0
in case the intersection is an empty set.

For example, the intersection of segments [1;5]
and [3;10]
is [3;5]
(length 2
), the intersection of segments [1;5]
and [5;7]
is [5;5]
(length 0
) and the intersection of segments [1;5]
and [6;6]
is an empty set (length 0
).

Your task is to remove exactly one segment from the given sequence in such a way that the intersection of the
remaining (𝑛−1)
segments has the maximal possible length.

### ideas

1. 考虑整个区间，重叠区域的l肯定是某个区域的左端点，右端点，肯定是每个区域的右端点
2. 对于任何一个这样的点（假设是左端点），假设它的active = n - 1
3. 那么如果存在某个区域还没有添加进来（在后面）或者它已经在前面终结了（这样的区域只能有1个）
4. 那么目前最小的右端点到当前节点的距离就是一个答案
5. 有没有其他可能性呢？
6. 不会，另外一个就是当 active = n的时候，但是因为r只会变的更小，且l变的更大，所以区间更小了
7. 第一种case就是去掉最近结束的点，或者最远开始的点