Let's denote a m-free matrix as a binary (that is, consisting of only 1's and 0's) matrix such that every square
submatrix of size m × m of this matrix contains at least one zero.

Consider the following problem:

You are given two integers n and m. You have to construct an m-free square matrix of size n × n such that the number of
1's in this matrix is maximum possible. Print the maximum possible number of 1's in such matrix.

You don't have to solve this problem. Instead, you have to construct a few tests for it.

You will be given t numbers x1, x2, ..., xt. For every , find two integers ni and mi (ni ≥ mi) such that the answer for
the aforementioned problem is exactly xi if we set n = ni and m = mi.

### ideas

1. 给定的一个matrix (n * n), 最多的1的数量，是可以计算的
2. 只在必须出现0的时候，放一个0即可（在右下角）
3. 如果 n % m = 0, 0的数量 = n / m * n / m
4. 如果 n % m != 0, let nx = n / m * m 也是 n/m * n / m
5. 但是问题是两个变量呐，
6. x(1的数量) = n * n - l * l
7. x = 两个平方数的差
8. x = (n - l) * (n + l)
9. 所以，要对x进行因式分解
10. x > 0, n != l