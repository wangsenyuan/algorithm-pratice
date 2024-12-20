Vasya has an array 𝑎1,𝑎2,…,𝑎𝑛
.

You don't know this array, but he told you 𝑚
 facts about this array. The 𝑖
-th fact is a triple of numbers 𝑡𝑖
, 𝑙𝑖
 and 𝑟𝑖
 (0≤𝑡𝑖≤1,1≤𝑙𝑖<𝑟𝑖≤𝑛
) and it means:

if 𝑡𝑖=1
 then subbarray 𝑎𝑙𝑖,𝑎𝑙𝑖+1,…,𝑎𝑟𝑖
 is sorted in non-decreasing order;
if 𝑡𝑖=0
 then subbarray 𝑎𝑙𝑖,𝑎𝑙𝑖+1,…,𝑎𝑟𝑖
 is not sorted in non-decreasing order. A subarray is not sorted if there is at least one pair of consecutive elements in this subarray such that the former is greater than the latter.
For example if 𝑎=[2,1,1,3,2]
 then he could give you three facts: 𝑡1=1,𝑙1=2,𝑟1=4
 (the subarray [𝑎2,𝑎3,𝑎4]=[1,1,3]
 is sorted), 𝑡2=0,𝑙2=4,𝑟2=5
 (the subarray [𝑎4,𝑎5]=[3,2]
 is not sorted), and 𝑡3=0,𝑙3=3,𝑟3=5
 (the subarray [𝑎3,𝑎5]=[1,3,2]
 is not sorted).

You don't know the array 𝑎
. Find any array which satisfies all the given facts.