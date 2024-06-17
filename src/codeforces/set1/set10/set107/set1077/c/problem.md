Let's call an array good if there is an element in the array that equals to the sum of all other elements. For example, the array 𝑎=[1,3,3,7]
 is good because there is the element 𝑎4=7
 which equals to the sum 1+3+3
.

You are given an array 𝑎
 consisting of 𝑛
 integers. Your task is to print all indices 𝑗
 of this array such that after removing the 𝑗
-th element from the array it will be good (let's call such indices nice).

For example, if 𝑎=[8,3,5,2]
, the nice indices are 1
 and 4
:

if you remove 𝑎1
, the array will look like [3,5,2]
 and it is good;
if you remove 𝑎4
, the array will look like [8,3,5]
 and it is good.
You have to consider all removals independently, i. e. remove the element, check if the resulting array is good, and return the element into the array.

