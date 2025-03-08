Let A = {a1, a2, ..., an} be any permutation of the first n natural numbers {1, 2, ..., n}. You are given a positive integer k and another sequence B = {b1, b2, ..., bn}, where bi is the number of elements aj in A to the left of the element at = i such that aj ≥ (i + k).

For example, if n = 5, a possible A is {5, 1, 4, 2, 3}. For k = 2, B is given by {1, 2, 1, 0, 0}. But if k = 3, then B = {1, 1, 0, 0, 0}.

For two sequences X = {x1, x2, ..., xn} and Y = {y1, y2, ..., yn}, let i-th elements be the first elements such that xi ≠ yi. If xi < yi, then X is lexicographically smaller than Y, while if xi > yi, then X is lexicographically greater than Y.

Given n, k and B, you need to determine the lexicographically smallest A.

### ideas
1. a[j] = n >= i + k => i <= n - k
2. 如果b[i] = w, 如果 i <= n - k (这个是确定的)， 但是如果b[i] = 0
3. A = [5, 1, 4, 2, 3], B = [1, 2, 1, 0, 0], and k = 2
4. i <= 5 - 2 = 3, 也就是说这里A[1] = 5, 所以它在1.。。3这个区间贡献为1
5. 如果A[2] = 5, 那么就在区间[2.3]贡献= 1，
6. 如果A[3] = 5, 那么就在区间[3.3]贡献1
7. B[i] = 2 > 0, 那么表示在存在某个数 x >= i + k = i + 2