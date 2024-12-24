Consider a permutation∗
 𝑝1,𝑝2,…,𝑝𝑛
 of integers from 1
 to 𝑛
. We can introduce the following sum for it†
:

𝑆(𝑝)=∑1≤𝑙≤𝑟≤𝑛min(𝑝𝑙,𝑝𝑙+1,…,𝑝𝑟)

Let us consider all permutations of length 𝑛
 with the maximum possible value of 𝑆(𝑝)
. Output the 𝑘
-th of them in lexicographical‡
order, or report that there are less than 𝑘
 of them.

∗
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

†
For example:

For the permutation [1,2,3]
 the value of 𝑆(𝑝)
 is equal to min(1)+min(1,2)+min(1,2,3)+min(2)+min(2,3)+min(3)=
 1+1+1+2+2+3=10
For the permutation [2,4,1,3]
 the value of 𝑆(𝑝)
 is equal to min(2)+min(2,4)+min(2,4,1)+min(2,4,1,3) +
 min(4)+min(4,1)+min(4,1,3) +
 min(1)+min(1,3) +
 min(3)=
 2+2+1+1+4+1+1+1+1+3=17
.
‡
An array 𝑎
 is lexicographically smaller than an array 𝑏
 if and only if one of the following holds:

𝑎
 is a prefix of 𝑏
, but 𝑎≠𝑏
; or
in the first position where 𝑎
 and 𝑏
 differ, the array 𝑎
 has a smaller element than the corresponding element in 𝑏
.