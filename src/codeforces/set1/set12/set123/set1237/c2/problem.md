This is a harder version of the problem. In this version, 𝑛≤50000
.

There are 𝑛
 distinct points in three-dimensional space numbered from 1
 to 𝑛
. The 𝑖
-th point has coordinates (𝑥𝑖,𝑦𝑖,𝑧𝑖)
. The number of points 𝑛
 is even.

You'd like to remove all 𝑛
 points using a sequence of 𝑛2
 snaps. In one snap, you can remove any two points 𝑎
 and 𝑏
 that have not been removed yet and form a perfectly balanced pair. A pair of points 𝑎
 and 𝑏
 is perfectly balanced if no other point 𝑐
 (that has not been removed yet) lies within the axis-aligned minimum bounding box of points 𝑎
 and 𝑏
.

Formally, point 𝑐
 lies within the axis-aligned minimum bounding box of points 𝑎
 and 𝑏
 if and only if min(𝑥𝑎,𝑥𝑏)≤𝑥𝑐≤max(𝑥𝑎,𝑥𝑏)
, min(𝑦𝑎,𝑦𝑏)≤𝑦𝑐≤max(𝑦𝑎,𝑦𝑏)
, and min(𝑧𝑎,𝑧𝑏)≤𝑧𝑐≤max(𝑧𝑎,𝑧𝑏)
. Note that the bounding box might be degenerate.

Find a way to remove all points in 𝑛2
 snaps.

 ### ideas
 1. 考虑将points按照x升序排
 2. 那么考虑左边的一个点a, 如果在x方向上有两个点b, c,其中b在x方向离a更近，但是在比如y方向，离得比较远
 3. 这种情况下，可以选择b和a配对吗？
 4. 在2维的情况下考虑，似乎是可以的。如果在(a, b)中间有一个点，无论这个点的y坐标如何变化，它肯定离a在x方向上更近
 5. 但是这样子可以在3维成立吗？
 6. 似乎也是成立的
 7. 所以直接排序就好了。在x相等的情况下，寻找y方向的，在y相等的情况下，再按照z方向排