Consider a number axis. The capital of your empire is initially at 0
. There are 𝑛
unconquered kingdoms at positions 0<𝑥1<𝑥2<…<𝑥𝑛
. You want to conquer all other kingdoms.

There are two actions available to you:

You can change the location of your capital (let its current position be 𝑐1
) to any other conquered kingdom (let its position be 𝑐2
) at a cost of 𝑎⋅|𝑐1−𝑐2|
.
From the current capital (let its current position be 𝑐1
) you can conquer an unconquered kingdom (let its position be 𝑐2
) at a cost of 𝑏⋅|𝑐1−𝑐2|
. You cannot conquer a kingdom if there is an unconquered kingdom between the target and your capital.
Note that you cannot place the capital at a point without a kingdom. In other words, at any point, your capital can only
be at 0
or one of 𝑥1,𝑥2,…,𝑥𝑛
. Also note that conquering a kingdom does not change the position of your capital.

Find the minimum total cost to conquer all kingdoms. Your capital can be anywhere at the end.

### ideas

1. 如果a <= b, 貌似随着征服新的王国，就把首都搬过去，是个更优的选择？
2. 假设当前首都位置在x1, 且最右边征服的王国是x2, 现在要征服它右边的王国x3
3. 如果不移动首都征服x3的cost = b * (x3 - x1)
4. 如果移动移动首都到x2, a * (x2 - x1), 征服x3 + b * (x3 - x2)
5. 因为 a <= b => a * (x2 - x1) + b * (x3 -x2) <= b * (x2 - x1) + b * (x3 - x2) = b * (x3 - x1)
6. 如果 a > b似乎情况要复杂一些
7. 所以要换个似乎，如果最后的落脚点在xi，那么右边的cost是确定的, = b * sum(xj - xi) = b * (sum(xj) - (n - i) * xi)
8. 左边的就是不断移动到xi，征服一个区域就移动过去