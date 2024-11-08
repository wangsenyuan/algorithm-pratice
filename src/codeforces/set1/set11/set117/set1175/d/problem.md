You are given an array 𝑎1,𝑎2,…,𝑎𝑛
 and an integer 𝑘
.

You are asked to divide this array into 𝑘
 non-empty consecutive subarrays. Every element in the array should be included in exactly one subarray. Let 𝑓(𝑖)
 be the index of subarray the 𝑖
-th element belongs to. Subarrays are numbered from left to right and from 1
 to 𝑘
.

Let the cost of division be equal to ∑𝑖=1𝑛(𝑎𝑖⋅𝑓(𝑖))
. For example, if 𝑎=[1,−2,−3,4,−5,6,−7]
 and we divide it into 3
 subbarays in the following way: [1,−2,−3],[4,−5],[6,−7]
, then the cost of division is equal to 1⋅1−2⋅1−3⋅1+4⋅2−5⋅2+6⋅3−7⋅3=−9
.

Calculate the maximum cost you can obtain by dividing the array 𝑎
 into 𝑘
 non-empty consecutive subarrays.


 