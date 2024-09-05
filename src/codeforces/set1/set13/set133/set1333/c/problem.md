Eugene likes working with arrays. And today he needs your help in solving one challenging task.

An array 𝑐
 is a subarray of an array 𝑏
 if 𝑐
 can be obtained from 𝑏
 by deletion of several (possibly, zero or all) elements from the beginning and several (possibly, zero or all) elements from the end.

Let's call a nonempty array good if for every nonempty subarray of this array, sum of the elements of this subarray is nonzero. For example, array [−1,2,−3]
 is good, as all arrays [−1]
, [−1,2]
, [−1,2,−3]
, [2]
, [2,−3]
, [−3]
 have nonzero sums of elements. However, array [−1,2,−1,−3]
 isn't good, as his subarray [−1,2,−1]
 has sum of elements equal to 0
.

Help Eugene to calculate the number of nonempty good subarrays of a given array 𝑎
.

### ideas
1. 如果 a[l...r] sum = 0， 所有包含它的都是bad
2. 所以可以反过来计算，那些bad的，然后从总数中减去，剩下的就是good的