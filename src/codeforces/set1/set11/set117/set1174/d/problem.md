Given two integers 𝑛
 and 𝑥
, construct an array that satisfies the following conditions:

for any element 𝑎𝑖
 in the array, 1≤𝑎𝑖<2𝑛
;
there is no non-empty subsegment with bitwise XOR equal to 0
 or 𝑥
,
its length 𝑙
 should be maximized.
A sequence 𝑏
 is a subsegment of a sequence 𝑎
 if 𝑏
 can be obtained from 𝑎
 by deletion of several (possibly, zero or all) elements from the beginning and several (possibly, zero or all) elements from the end.

 ### ideas
 1. 不能出现连续的一段，xor = 0, 或者 xor = x
 2. 单独考虑0的问题，
 3. 1 ？ 1 ？ 1 ？。。。。