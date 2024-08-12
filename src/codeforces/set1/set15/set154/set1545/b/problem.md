Cirno gave AquaMoon a chessboard of size 1×𝑛
. Its cells are numbered with integers from 1
 to 𝑛
 from left to right. In the beginning, some of the cells are occupied with at most one pawn, and other cells are unoccupied.

In each operation, AquaMoon can choose a cell 𝑖
 with a pawn, and do either of the following (if possible):

Move pawn from it to the (𝑖+2)
-th cell, if 𝑖+2≤𝑛
 and the (𝑖+1)
-th cell is occupied and the (𝑖+2)
-th cell is unoccupied.
Move pawn from it to the (𝑖−2)
-th cell, if 𝑖−2≥1
 and the (𝑖−1)
-th cell is occupied and the (𝑖−2)
-th cell is unoccupied.
You are given an initial state of the chessboard. AquaMoon wants to count the number of states reachable from the initial state with some sequence of operations. But she is not good at programming. Can you help her? As the answer can be large find it modulo 998244353
.

Input
The input consists of multiple test cases. The first line contains a single integer 𝑡
 (1≤𝑡≤10000
) — the number of test cases.

The first line contains a single integer 𝑛
 (1≤𝑛≤105
) — the size of the chessboard.

The second line contains a string of 𝑛
 characters, consists of characters "0" and "1". If the 𝑖
-th character is "1", the 𝑖
-th cell is initially occupied; otherwise, the 𝑖
-th cell is initially unoccupied.

It is guaranteed that the sum of 𝑛
 over all test cases does not exceed 105
.

Output
For each test case, print the number of states that reachable from the initial state with some sequence of operations modulo 998244353
.

### ideas
1. 如果只有一个pawn，它没办法移动
2. 如果有两个连续的pawn，可以随意左右一起移动（直到遇到下一个pawn）
3. 如果是3个连续的pawn呢？4个呢？
4. 如果是4个的话，这个和2个的有点像，区别在于，它们不需要一直连在一起。
5. dp[i][state] 表示到i为止，最后两位有状态state表示的情况
6. 如果当前位为1， dp[i][10] = dp[i-1][01] (把当前的1移动到 i-2 位置)
7.              dp[i][11]/dp[i][01] 表示当前的1不移动
8.  但是如何表示把i移动到后面去呢？
9.       如果当前位为0， dp[i][11] = dp[i-1][11] 表示把i-2的1移动到i位置
10. 如果只移动一次是可以的。但是这个可以连续移动，就不大对了
11. 因为移入一个新的1后，前面的状态就不能用最后两位来表示了
12. 感觉还是要两个一组
13. 多出来的那个去掉（因为它要么不能移动，要们造成其他的不能移动）
14. 