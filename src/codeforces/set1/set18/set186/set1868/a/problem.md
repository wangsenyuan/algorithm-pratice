There is an empty matrix 𝑀
of size 𝑛×𝑚
.

Zhongkao examination is over, and Daniel would like to do some puzzle games. He is going to fill in the matrix 𝑀
using permutations of length 𝑚
. That is, each row of 𝑀
must be a permutation of length 𝑚†
.

Define the value of the 𝑖
-th column in 𝑀
as 𝑣𝑖=MEX(𝑀1,𝑖,𝑀2,𝑖,…,𝑀𝑛,𝑖)‡
. Since Daniel likes diversity, the beauty of 𝑀
is 𝑠=MEX(𝑣1,𝑣2,⋯,𝑣𝑚)
.

You have to help Daniel fill in the matrix 𝑀
and maximize its beauty.

†
A permutation of length 𝑚
is an array consisting of 𝑚
distinct integers from 0
to 𝑚−1
in arbitrary order. For example, [1,2,0,4,3]
is a permutation, but [0,1,1]
is not a permutation (1
appears twice in the array), and [0,1,3]
is also not a permutation (𝑚−1=2
but there is 3
in the array).

‡
The MEX
of an array is the smallest non-negative integer that does not belong to the array. For example, MEX(2,2,1)=0
because 0
does not belong to the array, and MEX(0,3,1,2)=4
because 0
, 1
, 2
and 3
appear in the array, but 4
does not.

### thoughts

1. 考虑最后的结果，假设是x, 那么 0,1,2...x - 1 必须在某一列中存在
2. 如何得到x - 1 呢