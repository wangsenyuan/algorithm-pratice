The world is a grid with 𝑛
 rows and 𝑚
 columns. The rows are numbered 0,1,…,𝑛−1
, while the columns are numbered 0,1,…,𝑚−1
. In this world, the columns are cyclic (i.e. the top and the bottom cells in each column are adjacent). The cell on the 𝑖
-th row and the 𝑗
-th column (0≤𝑖<𝑛,0≤𝑗<𝑚
) is denoted as (𝑖,𝑗)
.

At time 0
, the cell (𝑖,𝑗)
 (where 0≤𝑖<𝑛,0≤𝑗<𝑚
) contains either a rock or nothing. The state of cell (𝑖,𝑗)
 can be described using the integer 𝑎𝑖,𝑗
:

If 𝑎𝑖,𝑗=1
, there is a rock at (𝑖,𝑗)
.
If 𝑎𝑖,𝑗=0
, there is nothing at (𝑖,𝑗)
.
As a result of aftershocks from the earthquake, the columns follow tectonic plate movements: each column moves cyclically upwards at a velocity of 1
 cell per unit of time. Formally, for some 0≤𝑖<𝑛,0≤𝑗<𝑚
, if (𝑖,𝑗)
 contains a rock at the moment, it will move from (𝑖,𝑗)
 to (𝑖−1,𝑗)
 (or to (𝑛−1,𝑗)
 if 𝑖=0
).

The robot called RT is initially positioned at (0,0)
. It has to go to (𝑛−1,𝑚−1)
 to carry out an earthquake rescue operation (to the bottom rightmost cell). The earthquake doesn't change the position of the robot, they only change the position of rocks in the world.

Let RT's current position be (𝑥,𝑦)
 (0≤𝑥<𝑛,0≤𝑦<𝑚
), it can perform the following operations:

Go one cell cyclically upwards, i.e. from (𝑥,𝑦)
 to ((𝑥+𝑛−1)mod𝑛,𝑦)
 using 1
 unit of time.
Go one cell cyclically downwards, i.e. (𝑥,𝑦)
 to ((𝑥+1)mod𝑛,𝑦)
 using 1
 unit of time.
Go one cell to the right, i.e. (𝑥,𝑦)
 to (𝑥,𝑦+1)
 using 1
 unit of time. (RT may perform this operation only if 𝑦<𝑚−1
.)
Note that RT cannot go left using the operations nor can he stay at a position.

Unfortunately, RT will explode upon colliding with a rock. As such, when RT is at (𝑥,𝑦)
 and there is a rock at ((𝑥+1)mod𝑛,𝑦)
 or ((𝑥+2)mod𝑛,𝑦)
, RT cannot move down or it will be hit by the rock.


### ideas
1. 这里机器人可以往上走，所以有点复杂性
2. dp[i,j,k]表示到达位置(i, j)时，且在这个过程中向上移动了k次后，是否安全
3. dp[i-1, j, k] = dp[i, j, k - 1] 这个一直成立（相当于随着column j移动）
4. dp[i+1, j, k] = dp[i+1, j, k] 向下移动，可以判断一下，进过了时刻(i+1 + j + k)后，在位置(i+1, j)是否安全
5. dp[i, j+1, k] = dp[i, j, k]向右移动，可以判断一下，进过了时间(i+j+1+k)后，在位置(i, j+1)是否安全
6. 这个复杂性时 n * m * k (向上移动的次数)显然是不行的
7. 且k看上去似乎可以一直增长
8. 按列考虑，在每一列，对于每一列考虑什么时候能够安全的到达，然后向右边移动，因为column是shift的，所以，相当于可以移动到(i+1)的位置，然后在同一列里面上下移动，直到找到一个安全的位置，移动到下一列
9. got
10. 按列考虑dp[i]表示这一列中，最早到到i行的时刻
11. 这个地方出现了问题，看例子1，当机器人向上的时候，其实相当于它的位置没有变化

### solution

By viewing the robot's movement relative to the rocks, Robot RT's three moves become as follows:

Up: Stationary
Down, (𝑥,𝑦)
 to ((𝑥+2)mod𝑛,𝑦)
Right, (𝑥,𝑦)
 to ((𝑥+1)mod𝑛,𝑦+1)
As staying stationary is not necessary now when we are finding the minimum time, we can run a bfs/dp from (0,0)
 to find the minimum time required to reach every grid in the second last column (𝑥𝑚𝑜𝑑𝑛,𝑚−2)
. Finally, choose the best among all n tiles after waiting for the endpoint to cycle back.