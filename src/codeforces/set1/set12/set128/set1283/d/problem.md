There are 𝑛
Christmas trees on an infinite number line. The 𝑖
-th tree grows at the position 𝑥𝑖
. All 𝑥𝑖
are guaranteed to be distinct.

Each integer point can be either occupied by the Christmas tree, by the human or not occupied at all. Non-integer points
cannot be occupied by anything.

There are 𝑚
people who want to celebrate Christmas. Let 𝑦1,𝑦2,…,𝑦𝑚
be the positions of people (note that all values 𝑥1,𝑥2,…,𝑥𝑛,𝑦1,𝑦2,…,𝑦𝑚
should be distinct and all 𝑦𝑗
should be integer). You want to find such an arrangement of people that the value ∑𝑗=1𝑚min𝑖=1𝑛|𝑥𝑖−𝑦𝑗|
is the minimum possible (in other words, the sum of distances to the nearest Christmas tree for all people should be
minimized).

In other words, let 𝑑𝑗
be the distance from the 𝑗
-th human to the nearest Christmas tree (𝑑𝑗=min𝑖=1𝑛|𝑦𝑗−𝑥𝑖|
). Then you need to choose such positions 𝑦1,𝑦2,…,𝑦𝑚
that ∑𝑗=1𝑚𝑑𝑗
is the minimum possible.

### ideas

1. sort x first
2. 可以二分吗？
3. 一共n个tree，人排列肯定是在树的两边
4. 怎么感觉是bfs？最多访问m个点