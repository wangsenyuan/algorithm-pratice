Vasya will fancy any number as long as it is an integer power of two. Petya, on the other hand, is very conservative and
only likes a single integer 𝑝
(which may be positive, negative, or zero). To combine their tastes, they invented 𝑝
-binary numbers of the form 2𝑥+𝑝
, where 𝑥
is a non-negative integer.

For example, some −9
-binary ("minus nine" binary) numbers are: −8
(minus eight), 7
and 1015
(−8=20−9
, 7=24−9
, 1015=210−9
).

The boys now use 𝑝
-binary numbers to represent everything. They now face a problem: given a positive integer 𝑛
, what's the smallest number of 𝑝
-binary numbers (not necessarily distinct) they need to represent 𝑛
as their sum? It may be possible that representation is impossible altogether. Help them solve this problem.

For example, if 𝑝=0
we can represent 7
as 20+21+22
.

And if 𝑝=−9
we can represent 7
as one number (24−9)
.

Note that negative 𝑝
-binary numbers are allowed to be in the sum (see the Notes section for an example).

### ideas

1. n = x + m * p
2. n - m * p = x 可以被表示成m个2进制数的和
3. brute force，然后检查是否ok即可