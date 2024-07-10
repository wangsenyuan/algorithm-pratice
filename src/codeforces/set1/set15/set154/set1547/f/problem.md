You are given an array of positive integers 𝑎=[𝑎0,𝑎1,…,𝑎𝑛−1]
 (𝑛≥2
).

In one step, the array 𝑎
 is replaced with another array of length 𝑛
, in which each element is the greatest common divisor (GCD) of two neighboring elements (the element itself and its right neighbor; consider that the right neighbor of the (𝑛−1)
-th element is the 0
-th element).

Formally speaking, a new array 𝑏=[𝑏0,𝑏1,…,𝑏𝑛−1]
 is being built from array 𝑎=[𝑎0,𝑎1,…,𝑎𝑛−1]
 such that 𝑏𝑖
 =gcd(𝑎𝑖,𝑎(𝑖+1)mod𝑛)
, where gcd(𝑥,𝑦)
 is the greatest common divisor of 𝑥
 and 𝑦
, and 𝑥mod𝑦
 is the remainder of 𝑥
 dividing by 𝑦
. In one step the array 𝑏
 is built and then the array 𝑎
 is replaced with 𝑏
 (that is, the assignment 𝑎
 := 𝑏
 is taking place).

For example, if 𝑎=[16,24,10,5]
 then 𝑏=[gcd(16,24)
, gcd(24,10)
, gcd(10,5)
, gcd(5,16)]
 =[8,2,5,1]
. Thus, after one step the array 𝑎=[16,24,10,5]
 will be equal to [8,2,5,1]
.

For a given array 𝑎
, find the minimum number of steps after which all values 𝑎𝑖
 become equal (that is, 𝑎0=𝑎1=⋯=𝑎𝑛−1
). If the original array 𝑎
 consists of identical elements then consider the number of steps is equal to 0
.