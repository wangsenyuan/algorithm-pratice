Maxim has an array 𝑎
 of 𝑛
 integers and an array 𝑏
 of 𝑚
 integers (𝑚≤𝑛
).

Maxim considers an array 𝑐
 of length 𝑚
 to be good if the elements of array 𝑐
 can be rearranged in such a way that at least 𝑘
 of them match the elements of array 𝑏
.

For example, if 𝑏=[1,2,3,4]
 and 𝑘=3
, then the arrays [4,1,2,3]
 and [2,3,4,5]
 are good (they can be reordered as follows: [1,2,3,4]
 and [5,2,3,4]
), while the arrays [3,4,5,6]
 and [3,4,3,4]
 are not good.

Maxim wants to choose every subsegment of array 𝑎
 of length 𝑚
 as the elements of array 𝑐
. Help Maxim count how many selected arrays will be good.

In other words, find the number of positions 1≤𝑙≤𝑛−𝑚+1
 such that the elements 𝑎𝑙,𝑎𝑙+1,…,𝑎𝑙+𝑚−1
 form a good array.

