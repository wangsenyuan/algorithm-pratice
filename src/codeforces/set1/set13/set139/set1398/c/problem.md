You are given an array 𝑎1,𝑎2,…,𝑎𝑛
 consisting of integers from 0
 to 9
. A subarray 𝑎𝑙,𝑎𝑙+1,𝑎𝑙+2,…,𝑎𝑟−1,𝑎𝑟
 is good if the sum of elements of this subarray is equal to the length of this subarray (∑𝑖=𝑙𝑟𝑎𝑖=𝑟−𝑙+1
).

For example, if 𝑎=[1,2,0]
, then there are 3
 good subarrays: 𝑎1…1=[1],𝑎2…3=[2,0]
 and 𝑎1…3=[1,2,0]
.

Calculate the number of good subarrays of the array 𝑎
.

### ideas
1. sum(r) - sum(l) = r - l => sum(r) - r = sum(l) - l