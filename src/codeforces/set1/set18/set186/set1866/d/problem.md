There are 𝑁
arrays, each array has 𝑀
positive integer elements The 𝑗
-th element of the 𝑖
-th array is 𝐴𝑖,𝑗
.

Initially, Chaneka's digital wallet contains 0
money. Given an integer 𝐾
. Chaneka will do 𝑀−𝐾+1
operations. In the 𝑝
-th operation, Chaneka does the following procedure:

Choose any array. Let's say Chaneka chooses the 𝑥
-th array.
Choose an index 𝑦
in that array such that 𝑝≤𝑦≤𝑝+𝐾−1
.
Add the value of 𝐴𝑥,𝑦
to the total money in the wallet.
Change the value of 𝐴𝑥,𝑦
into 0
.
Determine the maximum total money that can be earned!

### thoughts

1. 如果k=1, 那么就可以进行m次操作，且每次操作，只能选择第i列之后的列
2. 一种策略是选择每一列的最大值，但假如很后面某列有比较大的值，也可以替换掉
3. 但是第一列，只有一次几乎，第二列有2个机会，。。。。
4. 没想法～～～

### solution

First, notice that since the initial value of all elements are positive, it is always optimal to always choose an
element that has not been chosen before in each operation.

Let's look at the 𝑁
arrays as a grid of 𝑁
rows and 𝑀
columns. We can solve this problem by iterating each column from left to right and using dynamic programming. Define the
following:

dp[𝑖][𝑗]
: the maximum money Chaneka can earn by only taking the elements of the first 𝑖
columns in 𝑗
operations.
When we iterate a new column 𝑖
, we can choose to choose 𝑐
elements in that column for some value 𝑐
(0≤𝑐≤𝑁
). If we choose to choose 𝑐
elements in that column, it is always the most optimal to choose the 𝑐
biggest elements in that column.

Notice that if we only consider the first 𝑖
columns, in the optimal strategy, we must do a minimum of 𝑖−𝐾+1
operations and a maximum of 𝑖
operations. So there are only a maximum of 𝐾
different values of 𝑗
we need to consider when calculating all values of dp[𝑖][𝑗]
for each new column 𝑖
. For each dp[𝑖][𝑗]
, there are only 𝑁+1
possibilities for the number of elements we choose in column 𝑖
, and each of them can be handled in 𝑂(1)
.
