This is an easier version of the next problem. In this version, 𝑞=0
.

A sequence of integers is called nice if its elements are arranged in blocks like in [3,3,3,4,1,1]
. Formally, if two elements are equal, everything in between must also be equal.

Let's define difficulty of a sequence as a minimum possible number of elements to change to get a nice sequence. However, if you change at least one element of value 𝑥
 to value 𝑦
, you must also change all other elements of value 𝑥
 into 𝑦
 as well. For example, for [3,3,1,3,2,1,2]
 it isn't allowed to change first 1
 to 3
 and second 1
 to 2
. You need to leave 1
's untouched or change them to the same value.

You are given a sequence of integers 𝑎1,𝑎2,…,𝑎𝑛
 and 𝑞
 updates.

Each update is of form "𝑖
 𝑥
" — change 𝑎𝑖
 to 𝑥
. Updates are not independent (the change stays for the future).

Print the difficulty of the initial sequence and of the sequence after every update.


### ideas
1. 对于例子 1 2 1 2 3 1 1 1 50 1
2. 可以把2，3，50都转换为1， 答案为4
3. 把数字都变成区间,有重叠的都算在一个区间内，然后看这个区间内的，最多的那个，把余下的都变成最多的这个