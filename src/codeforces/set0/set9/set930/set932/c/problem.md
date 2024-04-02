For a permutation P[1... N] of integers from 1 to N, function f is defined as follows:

Let g(i) be the minimum positive integer j such that f(i, j)= i. We can show such j always exists.

For given N, A, B, find a permutation P of integers from 1 to N such that for 1 ≤ i ≤ N, g(i) equals either A or B.

### ideas

1. 好复杂
2. f(i, j) = { p[i] when j = 1, else f(p[i], j - 1) }
3. 假设p中有一个cycle的长度是r，那么f(i, j) = p[k] i和k的距离在cycle中的距离 = j - 1
4. g(i) = i 其实就是i所在的cycle的长度
5. 也就是，这辆要求有两类cycle，一类是cycle的长度=A，另一类cycle的长度=B
6. 假设有x个第一种cycle， B个第二种cycle: A * x + B * y =n
7. 然后迭代x的数量，计算处y的数量