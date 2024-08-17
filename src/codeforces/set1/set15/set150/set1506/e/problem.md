A permutation is a sequence of 𝑛
 integers from 1
 to 𝑛
, in which all numbers occur exactly once. For example, [1]
, [3,5,2,1,4]
, [1,3,2]
 are permutations, and [2,3,2]
, [4,3,1]
, [0]
 are not.

Polycarp was presented with a permutation 𝑝
 of numbers from 1
 to 𝑛
. However, when Polycarp came home, he noticed that in his pocket, the permutation 𝑝
 had turned into an array 𝑞
 according to the following rule:

𝑞𝑖=max(𝑝1,𝑝2,…,𝑝𝑖)
.
Now Polycarp wondered what lexicographically minimal and lexicographically maximal permutations could be presented to him.

An array 𝑎
 of length 𝑛
 is lexicographically smaller than an array 𝑏
 of length 𝑛
 if there is an index 𝑖
 (1≤𝑖≤𝑛
) such that the first 𝑖−1
 elements of arrays 𝑎
 and 𝑏
 are the same, and the 𝑖
-th element of the array 𝑎
 is less than the 𝑖
-th element of the array 𝑏
. For example, the array 𝑎=[1,3,2,3]
 is lexicographically smaller than the array 𝑏=[1,3,4,2]
.

For example, if 𝑛=7
 and 𝑝=[3,2,4,1,7,5,6]
, then 𝑞=[3,3,4,4,7,7,7]
 and the following permutations could have been as 𝑝
 initially:

[3,1,4,2,7,5,6]
 (lexicographically minimal permutation);
[3,1,4,2,7,6,5]
;
[3,2,4,1,7,5,6]
;
[3,2,4,1,7,6,5]
 (lexicographically maximum permutation).
For a given array 𝑞
, find the lexicographically minimal and lexicographically maximal permutations that could have been originally presented to Polycarp.

Input
The first line contains one integer 𝑡
 (1≤𝑡≤104
). Then 𝑡
 test cases follow.

The first line of each test case contains one integer 𝑛
 (1≤𝑛≤2⋅105
).

The second line of each test case contains 𝑛
 integers 𝑞1,𝑞2,…,𝑞𝑛
 (1≤𝑞𝑖≤𝑛
).

It is guaranteed that the array 𝑞
 was obtained by applying the rule from the statement to some permutation 𝑝
.

It is guaranteed that the sum of 𝑛
 over all test cases does not exceed 2⋅105
.

Output
For each test case, output two lines:

on the first line output 𝑛
 integers — lexicographically minimal permutation that could have been originally presented to Polycarp;
on the second line print 𝑛
 integers — lexicographically maximal permutation that could have been originally presented to Polycarp;

 ### ideas
 1. q[i] = n， i第一个n出现的位置, 那么 p[i] = n 一定成立
 2. 那么此时i后面的就是那些没有出现的元素
 3. 如果考虑最大的情况，那么就吧n-1, n-2... k 这些没有出现的倒序排列在n的后面
 4. 考虑min，正好反过来，从1开始