Let's call left cyclic shift of some string 𝑡1𝑡2𝑡3…𝑡𝑛−1𝑡𝑛
 as string 𝑡2𝑡3…𝑡𝑛−1𝑡𝑛𝑡1
.

Analogically, let's call right cyclic shift of string 𝑡
 as string 𝑡𝑛𝑡1𝑡2𝑡3…𝑡𝑛−1
.

Let's say string 𝑡
 is good if its left cyclic shift is equal to its right cyclic shift.

You are given string 𝑠
 which consists of digits 0–9.

What is the minimum number of characters you need to erase from 𝑠
 to make it good?


 ### ideas
 1. 显然任何两个字符长的，都是good的
 2. 任何相同字符组成的串，也是good的
 3. ababab => 也是ok的
 4. ababa => 不ok
 5. 