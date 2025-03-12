You are given a strip, consisting of 𝑛
 cells, all cells are initially colored red.

In one operation, you can choose a segment of consecutive cells and paint them blue. Before painting, the chosen cells can be either red or blue. Note that it is not possible to paint them red. You are allowed to perform at most 𝑘
 operations (possibly zero).

For each cell, the desired color after all operations is specified: red or blue.

It is clear that it is not always possible to satisfy all requirements within 𝑘
 operations. Therefore, for each cell, a penalty is also specified, which is applied if the cell ends up the wrong color after all operations. For the 𝑖
-th cell, the penalty is equal to 𝑎𝑖
.

The penalty of the final painting is calculated as the maximum penalty among all cells that are painted the wrong color. If there are no such cells, the painting penalty is equal to 0
.

What is the minimum penalty of the final painting that can be achieved?

### ideas
1. 按照a[i]降序排列，那么假设要达到最小的x = a[i], 那么i（包括）前面的都必须得到正确的颜色
2. 它这个是可以选择一段连续的（长度可以自己定）
3. 还是二分。但是检查的时候，忽略比x小的（不管它们是蓝色/红色）只考虑那些>= x的部分