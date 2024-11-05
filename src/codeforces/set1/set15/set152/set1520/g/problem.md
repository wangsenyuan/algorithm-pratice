Dima overslept the alarm clock, which was supposed to raise him to school.

Dima wonders if he will have time to come to the first lesson. To do this, he needs to know the minimum time it will take him to get from home to school.

The city where Dima lives is a rectangular field of 𝑛×𝑚
 size. Each cell (𝑖,𝑗)
 on this field is denoted by one number 𝑎𝑖𝑗
:

The number −1
 means that the passage through the cell is prohibited;
The number 0
 means that the cell is free and Dima can walk though it.
The number 𝑥
 (1≤𝑥≤109
) means that the cell contains a portal with a cost of 𝑥
. A cell with a portal is also considered free.
From any portal, Dima can go to any other portal, while the time of moving from the portal (𝑖,𝑗)
 to the portal (𝑥,𝑦)
 corresponds to the sum of their costs 𝑎𝑖𝑗+𝑎𝑥𝑦
.

In addition to moving between portals, Dima can also move between unoccupied cells adjacent to one side in time 𝑤
. In particular, he can enter a cell with a portal and not use it.

Initially, Dima is in the upper-left cell (1,1)
, and the school is in the lower right cell (𝑛,𝑚)
.

### ideas
1. 计算每个格子的最早到达时间
2. 主要的问题在于，如果处理那些从某个portal到达的时间
3. 假设进入(i, j)的时间是t[i][j], 那么所有其他有portal的位置，的时间 = t[i][j] + a[i][j] + a[x][y]
4. 显然不可能把所有的portal都算一遍
5. 这里考虑i, j是目前最早能够到达的位置，那么所有portal能够到达的位置，必然还没有到达
6. 也就是说，那些能够到达的位置的时间，它们同时改变了一个时间（a[i][j])
7. 也就是说，那些格子其实是 a[x][y]最小的那个
8. 也就是假设到到了 a[i][j]， 所以利用portal能够到达的时间，都是 a[i][j] + a[x][y]