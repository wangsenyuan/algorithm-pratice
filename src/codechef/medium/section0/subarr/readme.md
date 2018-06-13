1. Description
  You are given an array A, which consists of N integers. Also, you have an integer K.

  Let's call a subarray A[L..R] perfect, if the average of the numbers AL, AL + 1, ..., AR - 1, AR is greater than or equal to K.

  I.e. the average of the numbers {2, 5, 9, -3} equals to 3.25.

  So, your task is quite simple: you are to count the number of perfect subarrays of A.

2. Solution:
  2.1 先考虑当K=0
    那么该问题就变化为求有多少个区间[L, R], SUM[L, R] >= 0
    2.1.1 计算prefix sum array, S[i] = SUM(A[0], A[1], .... A[i])
    2.1.2 i in [0..n-1]
          找到有多少个j (< i), 并且S[j] <= S[i], 假设为c (S[-1] = 0)
          ans += c
    2.1.3 使用Segment Tree可以快速的实现
    2.1.4 Segment Tree用来保存比S[i]小的数的count
  2.2. 当K != 0时，可以将所有的数减去K，使得变化为上面的方式；
    证明 SUM[L, R] = SUM(A[L], A[L+1], ... A[R]) = SUM(A[L] - K, ..., A[R] - K) + (R - L + 1) * K
     SUM(A[L], .. A[R]) >= (R - L + 1) * K
     => SUM(A[L] - K, ...A[R] - K) >= 0