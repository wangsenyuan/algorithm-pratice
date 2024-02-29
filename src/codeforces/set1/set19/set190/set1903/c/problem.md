Theofanis easily gets obsessed with problems before going to sleep and often has nightmares about them. To deal with his
obsession he visited his doctor, Dr. Emix.

In his latest nightmare, he has an array 𝑎
of size 𝑛
and wants to divide it into non-empty subarrays†
such that every element is in exactly one of the subarrays.

For example, the array [1,−3,7,−6,2,5]
can be divided to [1][−3,7][−6,2][5]
.

The Cypriot value of such division is equal to Σ𝑘𝑖=1𝑖⋅sum𝑖
where 𝑘
is the number of subarrays that we divided the array into and sum𝑖
is the sum of the 𝑖
-th subarray.

The Cypriot value of this division of the array [1][−3,7][−6,2][5]=1⋅1+2⋅(−3+7)+3⋅(−6+2)+4⋅5=17
.

Theofanis is wondering what is the maximum Cypriot value of any division of the array.

†
An array 𝑏
is a subarray of an array 𝑎
if 𝑏
can be obtained from 𝑎
by deletion of several (possibly, zero or all) elements from the beginning and several (possibly, zero or all) elements
from the end. In particular, an array is a subarray of itself.

### thoughts

1. 考虑最优的分配中，靠近的两段，在边界的位置
2. [a...b] [c...d]
3. 如果 b > 0, 那么把它移动到后面是更好，因为后面的序号更大，乘正数会得到更好的结果
4. 如果 c < 0, 那么把它移动到前面更好
5. 所以，可以确定的是，边界必然发生在一个负数和一个正数之间
6. 倒过来处理，尽量保证每段的sum > 0