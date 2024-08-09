The only difference between this problem and D2 is that you don't have to provide the way to construct the answer in this problem, but you have to do it in D2.

There's a table of 𝑛×𝑚
 cells (𝑛
 rows and 𝑚
 columns). The value of 𝑛⋅𝑚
 is even.

A domino is a figure that consists of two cells having a common side. It may be horizontal (one of the cells is to the right of the other) or vertical (one of the cells is above the other).

You need to find out whether it is possible to place 𝑛𝑚2
 dominoes on the table so that exactly 𝑘
 of them are horizontal and all the other dominoes are vertical. The dominoes cannot overlap and must fill the whole table.

### ideas
1. n * m is even, 所以肯定可以选择 n * m / 2 个domino，把这个grid铺满
2. 但要求其中必须有k个是horitontal的
3. 如果 n % 2 == 1, 或者 m % 2 == 1 (只能有一个为1)
4. 那么这个边, n - 1, 且这些只能用 m/2 水平的来铺
5. 如果 k < m/2 => false
6. else剩下的必须 k - m / 2 必须是偶数 （2个水平的换两个垂直的）