You are given a positive integer 𝑛
.

In this problem, the MEX
of a collection of integers 𝑐1,𝑐2,…,𝑐𝑘
is defined as the smallest positive integer 𝑥
which does not occur in the collection 𝑐
.

The primality of an array 𝑎1,…,𝑎𝑛
is defined as the number of pairs (𝑙,𝑟)
such that 1≤𝑙≤𝑟≤𝑛
and MEX(𝑎𝑙,…,𝑎𝑟)
is a prime number.

Find any permutation of 1,2,…,𝑛
with the maximum possible primality among all permutations of 1,2,…,𝑛
.

Note:

A prime number is a number greater than or equal to 2
that is not divisible by any positive integer except 1
and itself. For example, 2,5,13
are prime numbers, but 1
and 6
are not prime numbers.
A permutation of 1,2,…,𝑛
is an array consisting of 𝑛
distinct integers from 1
to 𝑛
in arbitrary order. For example, [2,3,1,5,4]
is a permutation, but [1,2,2]
is not a permutation (2
appears twice in the array), and [1,3,4]
is also not a permutation (𝑛=3
but there is 4
in the array).

### thoughts

1. 观察一下，如果1和2是分开的，那些只包含1的区间，它们的mex = 2, 那些只包括2的区间，它们的mex = 1
2. 那些包含区间1，2的，它们的mex >= 3
3. 如果它1放在最中间，且把2放在一端，那么可以制作出最多mex=2的区间，如果存在3，把它放在另外一端即可