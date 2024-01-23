Alex got a new game called "GCD permutations" as a birthday present. Each round of this game proceeds as follows:

First, Alex chooses a permutation†
𝑎1,𝑎2,…,𝑎𝑛
of integers from 1
to 𝑛
.
Then, for each 𝑖
from 1
to 𝑛
, an integer 𝑑𝑖=gcd(𝑎𝑖,𝑎(𝑖mod𝑛)+1)
is calculated.
The score of the round is the number of distinct numbers among 𝑑1,𝑑2,…,𝑑𝑛
.
Alex has already played several rounds so he decided to find a permutation 𝑎1,𝑎2,…,𝑎𝑛
such that its score is as large as possible.

Recall that gcd(𝑥,𝑦)
denotes the greatest common divisor (GCD) of numbers 𝑥
and 𝑦
, and 𝑥mod𝑦
denotes the remainder of dividing 𝑥
by 𝑦
.

†
A permutation of length 𝑛
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

1. 显然最后gcd都会变成1，
2. 在此前，最好将gcd(a[i], a[i+1]) = a[i]的安排在一起
3. 那么这样子，每次只有最后一个 a[j] a[j+1] = 1
4. 1,2,4,8,16.3.6.9...
5. 怎么证明？