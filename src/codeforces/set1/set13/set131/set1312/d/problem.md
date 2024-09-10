Your task is to calculate the number of arrays such that:

each array contains 𝑛
 elements;
each element is an integer from 1
 to 𝑚
;
for each array, there is exactly one pair of equal elements;
for each array 𝑎
, there exists an index 𝑖
 such that the array is strictly ascending before the 𝑖
-th element and strictly descending after it (formally, it means that 𝑎𝑗<𝑎𝑗+1
, if 𝑗<𝑖
, and 𝑎𝑗>𝑎𝑗+1
, if 𝑗≥𝑖
).


### ideas
1. 