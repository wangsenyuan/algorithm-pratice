You are given two positive integer numbers x and y. An array F is called an y-factorization of x iff the following
conditions are met:

There are y elements in F, and all of them are integer numbers;
.
You have to count the number of pairwise distinct arrays that are y-factorizations of x. Two arrays A and B are
considered different iff there exists at least one index i (1 ≤ i ≤ y) such that Ai ≠ Bi. Since the answer can be very
large, print it modulo 109 + 7.

### ideas

1. prime factor x
2. 假设 x 的 pf的个数分别是 f[1], .... f[n]
3. 那么就是将f[1] 分到y组里面去，f[2] 分到y组里面去， 可以有空的组（表示0)
4. nCr(f[i] + y, y) ?
5. nCr(f[i] + y - 1, f[i])

### solution

Fill the array with ones. Now we should take every prime divisor i of x and distribute cnti (maximum power of this prime
to appear in x) of it into some cells of the array. It is pretty well-known problem, it's equal to . Take product of
this values for every prime i. This will be the answer if there were no negative numbers. But we should also multiply it
by number of ways to select even number of position to put unary minuses — 2y - 1 (like you can fill in y - 1 position
anyhow and the final one will be determined by parity of current count).

To process many queries you should factorize numbers in  (by precalcing the smallest prime divisor of every number up to
106 with sieve of Eratosthenes), get in O(1) (by precalcing factorials and inverse factorials) and get 2k in  (binary
exponentiation).

