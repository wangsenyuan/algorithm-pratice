Alice has got addicted to a game called Sirtet recently.

In Sirtet, player is given an 𝑛×𝑚
 grid. Initially 𝑎𝑖,𝑗
 cubes are stacked up in the cell (𝑖,𝑗)
. Two cells are called adjacent if they share a side. Player can perform the following operations:

stack up one cube in two adjacent cells;
stack up two cubes in one cell.
Cubes mentioned above are identical in height.

Here is an illustration of the game. States on the right are obtained by performing one of the above operations on the state on the left, and grey cubes are added due to the operation.

Player's goal is to make the height of all cells the same (i.e. so that each cell has the same number of cubes in it) using above operations.

Alice, however, has found out that on some starting grids she may never reach the goal no matter what strategy she uses. Thus, she is wondering the number of initial grids such that

𝐿≤𝑎𝑖,𝑗≤𝑅
 for all 1≤𝑖≤𝑛
, 1≤𝑗≤𝑚
;
player can reach the goal using above operations.
Please help Alice with it. Notice that the answer might be large, please output the desired value modulo 998,244,353
.

### ideas
1. 操作2，不会改变a[i][j]的高度的parity （这个从parity的角度看，没啥改变，好像可以忽略掉）
2. 操作1，同时改变了两个cell的高度的pairty
3. 考虑m=1时的情况，此时是一条线，如果此时，所有高度的parity的xor等于1的话，那么永远也做不到
4. 如果等于0，那么通过有限步，应该是可以做到的 （但是这里有个问题，就是如果那些已经等于R的，它们似乎需要特别处理)
5. 必须分段（在a[i][j] = R的地方进行分段）在每一段中， 高度的pairty的xor = 0
6. 那么在一个二维的平面上，就是看，被R高度，分割的区域里面的高度的xor是否是0了
7. 继续回到1维的情况考虑
8. 没想法～
9. 换个角度考虑，有多少种状态，可以从dp[h]得到
10. 所有的高度都是h，通过操作1、2能够到？
11. dp[l] = 1, (所有的都为l)
12. dp[l+1] = dp[l] (如果 n * m % 2 = 0, (通过操作， 到达dp[l]))
13.   最终的高度是可以超过r的
14. 可以理解solution，可以搞搞   

### solution
Observation 1. The actual values in the cells don't matter, only parity matters.

Proof. Using the second operation one can make all the values of same parity equal by applying it to the lowest value until done.

That observation helps us to get rid of the second operation, let us only have the first one.

Observation 2. Player is able to change parity of any pair of cells at the same time.

Proof. For any given cell 𝑢
 and cell 𝑣
, there exists a path from 𝑢
 to 𝑣
, namely 𝑝=𝑤0𝑤1…𝑤𝑘
, such that 𝑤0=𝑢
 and 𝑤𝑛=𝑣
. Perform operation for adjacent cells 𝑤𝑖−1
 and 𝑤𝑖
 for all 1≤𝑖≤𝑛
Observation 3. If 𝑛×𝑚
 is odd, player can always reach the goal no matter what the initial state is.

Proof. There are two cases: there is an even number of even cells or there is an even number of odd cells. Whichever of these holds, we can change all cells of that parity to the opposite one with the aforementioned operation, making all cells the same parity.

Observation 4. If 𝑛×𝑚
 is even, and ∑𝑛𝑖=1∑𝑚𝑗=1𝑎𝑖,𝑗
 is even, player can reach the goal.

Proof. Similar to the proof of Observation 3. Proof is omitted.

Observation 5. If 𝑛×𝑚
 is even and ∑𝑛𝑖=1∑𝑚𝑗=1𝑎𝑖,𝑗
 is odd, player can never reach the goal no matter what strategies player takes.

Proof. Note that applying the operation never changes the parity of the number of cells of each parity (i.e. if we start with an odd number of odd cells and an odd number of even cells, they will both be odd until the end). Thus, there is no way to make zero cells of some parity.

How does that help us to calculate the answer?

The first case (𝑛×𝑚
 is odd) is trivial, the answer is all grids. Let's declare this as 𝑡𝑜𝑡𝑎𝑙
 value.

The second case (𝑛×𝑚
 is even) is harder. Me and pikmike have different formulas to obtain it but the answer is the same.

WLOG, let's move the range of values from [𝐿;𝑅]
 to [0;𝑅−𝐿]
, let 𝑘=𝑅−𝐿
.

Easy combinatorics solution (me):

Let's find out the number of ways to choose the grid such that the number of even cells is even and 0≤𝑎𝑖≤𝑘
.

Suppose that there are 𝐸
 even numbers in [0,𝑘]
, 𝑂
 odds. Therefore, for any given 0≤𝑖≤𝑛𝑚
, the number of ways to have exactly 𝑖
 even numbers should be 𝐸𝑖𝑂𝑛𝑚−𝑖×(𝑛𝑚𝑖)
. Thus, the total answer should be ∑𝑖=0𝑛𝑚/2𝐸2𝑖𝑂𝑛𝑚−2𝑖(𝑛𝑚2𝑖)
, which should remind you of the Newton expansion.

Note that
(𝐸+𝑂)𝑛𝑚=∑𝑖=0𝑛𝑚/2𝐸2𝑖𝑂𝑛𝑚−2𝑖(𝑛𝑚2𝑖)+∑𝑖=1𝑛𝑚/2𝐸2𝑖−1𝑂𝑛𝑚−(2𝑖−1)(𝑛𝑚2𝑖−1)
and

(𝐸−𝑂)𝑛𝑚=∑𝑖=0𝑛𝑚/2𝐸2𝑖𝑂𝑛𝑚−2𝑖(𝑛𝑚2𝑖)−∑𝑖=1𝑛𝑚/2𝐸2𝑖−1𝑂𝑛𝑚−(2𝑖−1)(𝑛𝑚2𝑖−1)
.

Adding those two formulas will get us exactly the formula we were looking for but doubled. Thus, the answer is that divided by 2
.

Easy intuition solution (pikmike):

There is a general solution to this kind of problems. Let's try to pair up each valid grid with exactly one invalid grid. Valid in our problem is such a grid that the number of even cells is even. If such a matching exists then the answer is exactly half of all grids (𝑡𝑜𝑡𝑎𝑙2)
.

Let's come up with some way to pair them up. For example, this works but leaves us with two cases to handle.

Let 𝑘
 be odd. For each grid let's replace 𝑎0,0
 with 𝑎0,0 𝑥𝑜𝑟 1
. You can see that the parity changed, thus the number of even cells also changed its parity. So if the grid was invalid it became valid and vice versa.

For an even 𝑘
 it's slightly trickier but actually one can show that almost all grids pair up with each other and only a single grid remains unpaired. So we can add a fake grid and claim that the answer is 𝑡𝑜𝑡𝑎𝑙+12
.

The algorithm is the following. Find the first position such that the value 𝑎𝑖,𝑗
 on it is not equal to 𝑘
. Replace it with 𝑎𝑖,𝑗 𝑥𝑜𝑟 1
. You can easily see that only grid that consists of all numbers 𝑘
 has no pair.