You have an integer array 𝑎
of length 𝑛
. There are two kinds of operations you can make.

Remove an integer from 𝑎
. This operation costs 𝑐
.
Insert an arbitrary positive integer 𝑥
to any position of 𝑎
(to the front, to the back, or between any two consecutive elements). This operation costs 𝑑
.
You want to make the final array a permutation of any positive length. Please output the minimum cost of doing that. Note that you can make the array empty during the operations, but the final array must contain at least one integer.

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