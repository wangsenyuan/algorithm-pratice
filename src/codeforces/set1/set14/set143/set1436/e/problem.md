In this problem MEX of a certain array is the smallest positive integer not contained in this array.

Everyone knows this definition, including Lesha. But Lesha loves MEX, so he comes up with a new problem involving MEX
every day, including today.

You are given an array 𝑎
of length 𝑛
. Lesha considers all the non-empty subarrays of the initial array and computes MEX for each of them. Then Lesha
computes MEX of the obtained numbers.

An array 𝑏
is a subarray of an array 𝑎
, if 𝑏
can be obtained from 𝑎
by deletion of several (possible none or all) elements from the beginning and several (possibly none or all) elements
from the end. In particular, an array is a subarray of itself.

Lesha understands that the problem is very interesting this time, but he doesn't know how to solve it. Help him and find
the MEX of MEXes of all the subarrays!

### ideas

1. 思路有点乱，不大对
2. 一个subarray的mex = 这个区间内第一个没有出现的数x
3. 假设从1开始，像正常的找mex一样，往两边扩张，直达无法扩张，此时得到的区间是[l...r, x]的一个组合
4. 然后再找下个 [l...r, x]的组合，这个时候它们两个区间会重叠吗？