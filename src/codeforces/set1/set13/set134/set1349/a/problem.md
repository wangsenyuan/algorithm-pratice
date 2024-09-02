For the multiset of positive integers 𝑠={𝑠1,𝑠2,…,𝑠𝑘}
, define the Greatest Common Divisor (GCD) and Least Common Multiple (LCM) of 𝑠
 as follow:

gcd(𝑠)
 is the maximum positive integer 𝑥
, such that all integers in 𝑠
 are divisible on 𝑥
.
lcm(𝑠)
 is the minimum positive integer 𝑥
, that divisible on all integers from 𝑠
.
For example, gcd({8,12})=4,gcd({12,18,6})=6
 and lcm({4,6})=12
. Note that for any positive integer 𝑥
, gcd({𝑥})=lcm({𝑥})=𝑥
.

Orac has a sequence 𝑎
 with length 𝑛
. He come up with the multiset 𝑡={lcm({𝑎𝑖,𝑎𝑗}) | 𝑖<𝑗}
, and asked you to find the value of gcd(𝑡)
 for him. In other words, you need to calculate the GCD of LCMs of all pairs of elements in the given sequence.


 ### ideas
 1. 如果只有两个数，那么 f(arr) = lcm(a[0], a[1])
 2. 把它们拆成质因数来看，对于a[i], a[j]来说，保留的是该质因数的最大指数
 3. 但是从整体的gcd来说，又取的是最小的那个质因数的指数
 4. 所以，是取所有数中第二大的质因数的指数