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
 1. 3维很难思考，考虑2维的情况
 2. 在2维的情况下，就是从最左、最下的点开始，找到离它最近的点，然后移除它们
 3. 重复就可以了
 4. 在3维的情况下，也是ok的。就是找到最外面的一个点，然后移除离她最近的点
 5. 在n <= 2000 的情况下，确实是可以work的
 6. 先做完这个case再考虑