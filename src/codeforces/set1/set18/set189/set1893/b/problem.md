You are given an array 𝑎
consisting of 𝑛
integers, as well as an array 𝑏
consisting of 𝑚
integers.

Let LIS(𝑐)
denote the length of the longest increasing subsequence of array 𝑐
. For example, LIS([2,1⎯⎯,1,3⎯⎯])
= 2
, LIS([1⎯⎯,7⎯⎯,9⎯⎯])
= 3
, LIS([3,1⎯⎯,2⎯⎯,4⎯⎯])
= 3
.

You need to insert the numbers 𝑏1,𝑏2,…,𝑏𝑚
into the array 𝑎
, at any positions, in any order. Let the resulting array be 𝑐1,𝑐2,…,𝑐𝑛+𝑚
. You need to choose the positions for insertion in order to minimize LIS(𝑐)
.

Formally, you need to find an array 𝑐1,𝑐2,…,𝑐𝑛+𝑚
that simultaneously satisfies the following conditions:

The array 𝑎1,𝑎2,…,𝑎𝑛
is a subsequence of the array 𝑐1,𝑐2,…,𝑐𝑛+𝑚
.
The array 𝑐1,𝑐2,…,𝑐𝑛+𝑚
consists of the numbers 𝑎1,𝑎2,…,𝑎𝑛,𝑏1,𝑏2,…,𝑏𝑚
, possibly rearranged.
The value of LIS(𝑐)
is the minimum possible among all suitable arrays 𝑐
.

### thoughts

1. 一个直观的想法是， b要倒序排，因为顺序的情况下，它（有可能）会贡献超过1
2. 在处理到(i, j)时，假设目前获得的lis是c，i是仍然正常处理，在lis中找到正确的位置放入
3. 这时候应该选择什么样的j呢？如果j的加入不影响结果，直接加入是ok的，而且应该尽早的加入
4. 