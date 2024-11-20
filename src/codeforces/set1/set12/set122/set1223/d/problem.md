You are given a sequence 𝑎1,𝑎2,…,𝑎𝑛
, consisting of integers.

You can apply the following operation to this sequence: choose some integer 𝑥
 and move all elements equal to 𝑥
 either to the beginning, or to the end of 𝑎
. Note that you have to move all these elements in one direction in one operation.

For example, if 𝑎=[2,1,3,1,1,3,2]
, you can get the following sequences in one operation (for convenience, denote elements equal to 𝑥
 as 𝑥
-elements):

[1,1,1,2,3,3,2]
 if you move all 1
-elements to the beginning;
[2,3,3,2,1,1,1]
 if you move all 1
-elements to the end;
[2,2,1,3,1,1,3]
 if you move all 2
-elements to the beginning;
[1,3,1,1,3,2,2]
 if you move all 2
-elements to the end;
[3,3,2,1,1,1,2]
 if you move all 3
-elements to the beginning;
[2,1,1,1,2,3,3]
 if you move all 3
-elements to the end;
You have to determine the minimum number of such operations so that the sequence 𝑎
 becomes sorted in non-descending order. Non-descending order means that for all 𝑖
 from 2
 to 𝑛
, the condition 𝑎𝑖−1≤𝑎𝑖
 is satisfied.

Note that you have to answer 𝑞
 independent queries.

### ideas
1. 如果 1没有在最左端，那么它必须被移动到最左端
2. 如果n没有在最右端，那么它必须被移动到最右端
3. 从1开始，如果它最右边的下标是r,如果，在r的前面存在任何比1大的数，答案+1
4. 完蛋了，越来越混乱了