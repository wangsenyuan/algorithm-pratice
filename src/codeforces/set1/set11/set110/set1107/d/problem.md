You are given a binary matrix 𝐴
 of size 𝑛×𝑛
. Let's denote an 𝑥
-compression of the given matrix as a matrix 𝐵
 of size 𝑛𝑥×𝑛𝑥
 such that for every 𝑖∈[1,𝑛],𝑗∈[1,𝑛]
 the condition 𝐴[𝑖][𝑗]=𝐵[⌈𝑖𝑥⌉][⌈𝑗𝑥⌉]
 is met.

Obviously, 𝑥
-compression is possible only if 𝑥
 divides 𝑛
, but this condition is not enough. For example, the following matrix of size 2×2
 does not have any 2
-compression:

01
10
For the given matrix 𝐴
, find maximum 𝑥
 such that an 𝑥
-compression of this matrix is possible.

Note that the input is given in compressed form. But even though it is compressed, you'd better use fast input.


### 
1. 显然n % x = 0
2. 二分似乎也有点问题
3. 假设 x 是其中一个答案, 那么 x的因子是否也是一个答案呢？
4. 比如 x = 6, 可以，那么 x = 3是否也ok？
5. n % 6 = 0, 那么 n % 3 = 0成立
6. 假设 A[i][j] = B[i/6][j/6]
7. 考虑 i,j 都能整除6的情况
8. 比如 i = 12, j = 6, 有 A[12][6] = B[2][1]
9. 如果 x = 3, A[12][6] = B[4][2]
10. 似乎不一定成立，好像反过来是成立的
11. 上面的做法不大对
12. A[i][j] = B[1][1] 如果 i <=x && j <= x
13. A[i][j] = B[2][1] 如果 i > x && i <= 2 * x && j <= x
14. 换句话说，如果是x-compression，那么整个A可以分成 n / x * n / x 块
15. 每一块里面的值都是相同的
16. 为了让x最大，就是上 n / x 最小
17.  且x = 6， 满足条件， 那么2，3肯定也满足条件， 反过来却不一定
18.  但是反过来，如果x = 2, 不满足条件，那么 x = 6是不是肯定不满足？
19.  似乎是的，不然就形成了悖论
20.  