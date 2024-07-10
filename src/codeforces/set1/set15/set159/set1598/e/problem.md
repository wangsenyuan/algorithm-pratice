You are given a matrix, consisting of 𝑛
 rows and 𝑚
 columns. The rows are numbered top to bottom, the columns are numbered left to right.

Each cell of the matrix can be either free or locked.

Let's call a path in the matrix a staircase if it:

starts and ends in the free cell;
visits only free cells;
has one of the two following structures:
the second cell is 1
 to the right from the first one, the third cell is 1
 to the bottom from the second one, the fourth cell is 1
 to the right from the third one, and so on;
the second cell is 1
 to the bottom from the first one, the third cell is 1
 to the right from the second one, the fourth cell is 1
 to the bottom from the third one, and so on.
In particular, a path, consisting of a single cell, is considered to be a staircase.


### ideas
1. 这两个构型似乎是一致的，就是a是b的first前面加一个locked
2. b是a的first的上面加一个
3. 假设这样的一个staircase的长度是w，那么任意两个之间组成的都是staircase, w * (w + 1) / 2
4. 假设(x, y)所在的长度是w，那么把它给flip掉以后，是可以计算出新的w(需要知道第一个格子)
5. 然后就可以计算出新的
6. 所以，可以用dsu，用第一个格子作为root？
7. 不可以，因为断开后，需要重建这个dsu，这个cost太大。
8. 还有一个麻烦的地方在于，同一个格子，有可能属于最多两条路径
9. 所以一条路径，必须知道(x, y, d) d = 0, 表示水平展开， d = 1表示垂直展开
10. flip一个格子后，要看它属于哪条路径，进而进行相应的操作
11. 每条路径（考虑水平开始的）用(m - 1 - j + i)表示
12. 然后当改变一个格子时，(x, y) ，找到线(m - 1 - y + x)的最近的空点，和 (m - 1 - (y - 1) + x), (m - 1 - (y + 1) + x)的空点
13. 1的空点的下一个(r + 1, c + 1)是水平开始的位置, 也要考虑2的空点的下方（r+1, c）,
14. 思路是对的。但细节好麻烦啊。缓缓
15. 通过（x,y）的只有两条路径，一条水平开始的，一条垂直开始的，这两条是没有重叠部分的（除了x,y)
16. 它们是有重叠部分的。就是那些拐角处
17. 然后找出这两条路径的起始点，和终止点。然后如果是flip out，就减去原来的长度。
18. 我发现题目都理解错误了
19. 