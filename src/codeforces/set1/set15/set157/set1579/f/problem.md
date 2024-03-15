You are given an array 𝑎[0…𝑛−1]=[𝑎0,𝑎1,…,𝑎𝑛−1]
of zeroes and ones only. Note that in this problem, unlike the others, the array indexes are numbered from zero, not
from one.

In one step, the array 𝑎
is replaced by another array of length 𝑛
according to the following rules:

First, a new array 𝑎→𝑑
is defined as a cyclic shift of the array 𝑎
to the right by 𝑑
cells. The elements of this array can be defined as 𝑎→𝑑𝑖=𝑎(𝑖+𝑛−𝑑)mod𝑛
, where (𝑖+𝑛−𝑑)mod𝑛
is the remainder of integer division of 𝑖+𝑛−𝑑
by 𝑛
.
It means that the whole array 𝑎→𝑑
can be represented as a sequence
𝑎→𝑑=[𝑎𝑛−𝑑,𝑎𝑛−𝑑+1,…,𝑎𝑛−1,𝑎0,𝑎1,…,𝑎𝑛−𝑑−1]
Then each element of the array 𝑎𝑖
is replaced by 𝑎𝑖&𝑎→𝑑𝑖
, where &
is a logical "AND" operator.
For example, if 𝑎=[0,0,1,1]
and 𝑑=1
, then 𝑎→𝑑=[1,0,0,1]
and the value of 𝑎
after the first step will be [0&1,0&0,1&0,1&1]
, that is [0,0,0,1]
.

The process ends when the array stops changing. For a given array 𝑎
, determine whether it will consist of only zeros at the end of the process. If yes, also find the number of steps the
process will take before it finishes.