You are given an array of integers 𝑎
of size 𝑛
. You can choose a set of non-overlapping subarrays of the given array (note that some elements may be not included in
any subarray, this is allowed). For each selected subarray, calculate the MEX of its elements, and then calculate the
bitwise XOR of all the obtained MEX values. What is the maximum bitwise XOR that can be obtained?

The MEX (minimum excluded) of an array is the smallest non-negative integer that does not belong to the array. For
instance:

The MEX of [2,2,1]
is 0
, because 0
does not belong to the array.
The MEX of [3,1,0,1]
is 2
, because 0
and 1
belong to the array, but 2
does not.
The MEX of [0,3,1,2]
is 4
, because 0
, 1
, 2
and 3
belong to the array, but 4
does not.

### thoughts

1. result >= mex(arr)
2. a = mex(arr), 那么a的最高位，肯定在最后的最高位中有体现

### solution

Let's solve the problem using dynamic programming, let's store 𝑑𝑝[𝑖][𝑗]
such that 𝑑𝑝[𝑖][𝑗]=1
if it is possible to obtain an 𝑋𝑂𝑅
of 𝑀𝐸𝑋
values from the prefix up to 𝑖
(excluding 𝑖
) equal to 𝑗
, and 0
otherwise. Notice that the answer cannot be greater than 𝑛
. Therefore, the size of this two-dimensional array is 𝑂(𝑛2)
.

Let's learn how to solve this in 𝑂(𝑛3)
:

We iterate over 𝑙
from 1
to 𝑛
, and inside that, we iterate over 𝑟
from 𝑙
to 𝑛
, maintaining the value 𝑥=𝑀𝐸𝑋([𝑎𝑙,𝑎𝑙+1,…,𝑎𝑟])
. Then, for all 𝑗
from 0
to 𝑛
, we update 𝑑𝑝[𝑟+1][𝑗⊕𝑥]=1
if 𝑑𝑝[𝑙][𝑗]=1
. This way, we update based on the case when the resulting set includes the set [𝑎𝑙,𝑎𝑙+1,…,𝑎𝑟]
.

We also need to update if 𝑑𝑝[𝑙][𝑥]=1
, we assign 𝑑𝑝[𝑙+1][𝑥]=1
. This accounts for the case when we do not include 𝑎𝑙
in the answer.

Let's define 𝑀𝐸𝑋(𝑙,𝑟)=𝑀𝐸𝑋([𝑎𝑙,𝑎𝑙+1,…,𝑎𝑟])
, this will make the following text clearer.

Let's consider the segment 𝑙,𝑟
. Notice that if there exist 𝑙2
and 𝑟2
(𝑙≤𝑙2≤𝑟2≤𝑟
) such that 𝑙2≠𝑙
or 𝑟2≠𝑟
and 𝑀𝐸𝑋(𝑙,𝑟)=𝑀𝐸𝑋(𝑙2,𝑟2)
, then we can take the segment 𝑙2,𝑟2
instead of the segment 𝑙,𝑟
in the set of 𝑀𝐸𝑋
values, and the answer will remain the same. If, however, there is no such segment 𝑙2,𝑟2
for the segment 𝑙,𝑟
, then we call this segment 𝑙,𝑟
irreplaceable.

Now let's prove that there are no more than 𝑛⋅2
irreplaceable segments. For each irreplaceable segment, let's look at the larger element of the pair 𝑎𝑙,𝑎𝑟
, let's assume 𝑎𝑙
is larger (the other case is symmetric). Now, let's prove that there is at most one segment where 𝑎𝑙
is the left element and 𝑎𝑙≥𝑎𝑟
, by contradiction:

Suppose there are at least 2 such segments, let's call their right boundaries 𝑟1,𝑟2
(𝑟1<𝑟2
). Notice that 𝑀𝐸𝑋(𝑙,𝑟1)>𝑎𝑙
, otherwise the segment 𝑙,𝑟1
would not be irreplaceable (we could remove 𝑎𝑙
). Since 𝑎𝑙≥𝑎𝑟2
, then 𝑀𝐸𝑋(𝑙,𝑟1)>𝑎𝑟2
. It is obvious that 𝑎𝑟2
appears among the elements [𝑎𝑙,…,𝑎𝑟1]
, and therefore 𝑀𝐸𝑋(𝑙,𝑟2−1)=𝑀𝐸𝑋(𝑙,𝑟2)
, which means that the segment 𝑙,𝑟2
is not irreplaceable, contradiction.

For each 𝑎𝑖
, there is at most one irreplaceable segment where it is the smaller of the two extremes, and at most one where it is
the larger. Therefore, the total number of irreplaceable segments is no more than 2⋅𝑛
.

Let's update the DP only through the irreplaceable segments, then the solution works in 𝑂(𝑛2+𝑛⋅𝐶)
time, where 𝐶
is the number of irreplaceable segments. However, we have already proven that 𝐶≤2𝑛
, so the overall time complexity is 𝑂(𝑛2)
.