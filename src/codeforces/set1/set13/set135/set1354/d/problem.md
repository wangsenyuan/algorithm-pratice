You are given a multiset consisting of 𝑛
 integers. You have to process queries of two types:

add integer 𝑘
 into the multiset;
find the 𝑘
-th order statistics in the multiset and remove it.
𝑘
-th order statistics in the multiset is the 𝑘
-th element in the sorted list of all elements of the multiset. For example, if the multiset contains elements 1
, 4
, 2
, 1
, 4
, 5
, 7
, and 𝑘=3
, then you have to find the 3
-rd element in [1,1,2,4,4,5,7]
, which is 2
. If you try to delete an element which occurs multiple times in the multiset, only one occurence is removed.

After processing all queries, print any number belonging to the multiset, or say that it is empty.

