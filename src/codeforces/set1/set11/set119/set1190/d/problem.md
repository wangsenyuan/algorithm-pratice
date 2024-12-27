There are 𝑛
 points on the plane, the 𝑖
-th of which is at (𝑥𝑖,𝑦𝑖)
. Tokitsukaze wants to draw a strange rectangular area and pick all the points in the area.

The strange area is enclosed by three lines, 𝑥=𝑙
, 𝑦=𝑎
 and 𝑥=𝑟
, as its left side, its bottom side and its right side respectively, where 𝑙
, 𝑟
 and 𝑎
 can be any real numbers satisfying that 𝑙<𝑟
. The upper side of the area is boundless, which you can regard as a line parallel to the 𝑥
-axis at infinity. The following figure shows a strange rectangular area.


A point (𝑥𝑖,𝑦𝑖)
 is in the strange rectangular area if and only if 𝑙<𝑥𝑖<𝑟
 and 𝑦𝑖>𝑎
. For example, in the above figure, 𝑝1
 is in the area while 𝑝2
 is not.

Tokitsukaze wants to know how many different non-empty sets she can obtain by picking all the points in a strange rectangular area, where we think two sets are different if there exists at least one point in one set of them but not in the other.

Input
The first line contains a single integer 𝑛
 (1≤𝑛≤2×105
) — the number of points on the plane.

The 𝑖
-th of the next 𝑛
 lines contains two integers 𝑥𝑖
, 𝑦𝑖
 (1≤𝑥𝑖,𝑦𝑖≤109
) — the coordinates of the 𝑖
-th point.

All points are distinct.

### ideas
1. 考虑3个点,a, b,c, 其中a, c 在左右两边，b在中间
2. 假设他们的y相同，那么有 (a), (b), (c), (a, b), (b, c), (a, b, c) (但是没有 a, c)
3. 如果b在a，c的下方一点，就存在(a, c)
4. 一个感觉，可以从高往低的进行处理
5. 假设在y = i处，只考虑这个高度以上的边，得到的答案
6. 现在考虑 y = i - 1 处，假设在这条线上有w个点，那么就是 w * (1 + w) / 2 ? (至少有这么多个)
7. 感觉要把上面的点映射下来（计算w）但是呢，必须有当前行的至少一个节点（否则就和上面的重复了）
8. 假设在上一个这条线上的点，到目前这个点之间上面映射下来的点有v个，那么当前点可以恭喜 v + 1 个
9. 知道了，前缀和