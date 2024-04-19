You are given 𝑛
rectangles on a plane with coordinates of their bottom left and upper right points. Some (𝑛−1)
of the given 𝑛
rectangles have some common point. A point belongs to a rectangle if this point is strictly inside the rectangle or
belongs to its boundary.

Find any point with integer coordinates that belongs to at least (𝑛−1)
given rectangles.

Input
The first line contains a single integer 𝑛
(2≤𝑛≤132674
) — the number of given rectangles.

Each the next 𝑛
lines contains four integers 𝑥1
, 𝑦1
, 𝑥2
and 𝑦2
(−109≤𝑥1<𝑥2≤109
, −109≤𝑦1<𝑦2≤109
) — the coordinates of the bottom left and upper right corners of a rectangle.

Output
Print two integers 𝑥
and 𝑦
— the coordinates of any point that belongs to at least (𝑛−1)
given rectangles.

### ideas

1. 只需要关注4个端点
2. 假设存在这样的一个点，考虑它在x/y上的投影
3. 那么它的x必须被至少n-1个点区域覆盖
4. y也同样如此（但有可能不是同一批）
5. 但是否这样的点一定满足条件呢？
6. 不一定，考虑3个矩形，从x/y看都能找到2个区域的重叠，但整体上却没有重叠
7. 分情况考虑，如果这个点是被n个区域包围x，和n个区域包围y，那么 =》 ok
8. 或者 n个x/n个y也是 => ok
9. 如果是 n - 1 个区域，那么必须是这n-1个区域是重叠的
10. 找n-1比较难，那就找不包括的，如果那个区域不包含x，而不是2个 =》 OK
11. 似乎还是一样的，找出两个方向上最远的矩形，把它们给删除掉，判断剩下的区域是否有重叠即可