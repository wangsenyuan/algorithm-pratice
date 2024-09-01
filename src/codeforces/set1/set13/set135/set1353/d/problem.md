You are given an array 𝑎
 of length 𝑛
 consisting of zeros. You perform 𝑛
 actions with this array: during the 𝑖
-th action, the following sequence of operations appears:

Choose the maximum by length subarray (continuous subsegment) consisting only of zeros, among all such segments choose the leftmost one;
Let this segment be [𝑙;𝑟]
. If 𝑟−𝑙+1
 is odd (not divisible by 2
) then assign (set) 𝑎[𝑙+𝑟2]:=𝑖
 (where 𝑖
 is the number of the current action), otherwise (if 𝑟−𝑙+1
 is even) assign (set) 𝑎[𝑙+𝑟−12]:=𝑖
.
Consider the array 𝑎
 of length 5
 (initially 𝑎=[0,0,0,0,0]
). Then it changes as follows:

Firstly, we choose the segment [1;5]
 and assign 𝑎[3]:=1
, so 𝑎
 becomes [0,0,1,0,0]
;
then we choose the segment [1;2]
 and assign 𝑎[1]:=2
, so 𝑎
 becomes [2,0,1,0,0]
;
then we choose the segment [4;5]
 and assign 𝑎[4]:=3
, so 𝑎
 becomes [2,0,1,3,0]
;
then we choose the segment [2;2]
 and assign 𝑎[2]:=4
, so 𝑎
 becomes [2,4,1,3,0]
;
and at last we choose the segment [5;5]
 and assign 𝑎[5]:=5
, so 𝑎
 becomes [2,4,1,3,5]
.
Your task is to find the array 𝑎
 of length 𝑛
 after performing all 𝑛
 actions. Note that the answer exists and unique.

You have to answer 𝑡
 independent test cases.


 ### ideas
 1. just a bfs procedure