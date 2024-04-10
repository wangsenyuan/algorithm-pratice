Slava plays his favorite game "Peace Lightning". Now he is flying a bomber on a very specific map.

Formally, map is a checkered field of size 1 × n, the cells of which are numbered from 1 to n, in each cell there can be
one or several tanks. Slava doesn't know the number of tanks and their positions, because he flies very high, but he can
drop a bomb in any cell. All tanks in this cell will be damaged.

If a tank takes damage for the first time, it instantly moves to one of the neighboring cells (a tank in the cell n can
only move to the cell n - 1, a tank in the cell 1 can only move to the cell 2). If a tank takes damage for the second
time, it's counted as destroyed and never moves again. The tanks move only when they are damaged for the first time,
they do not move by themselves.

Help Slava to destroy all tanks using as few bombs as possible.

### ideas

1. 很有趣的一个题目
2. 如果n是奇数2,4,6... 1,3,5... 2,4,6