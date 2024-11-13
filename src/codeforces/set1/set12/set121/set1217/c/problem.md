You are given a binary string 𝑠
 (recall that a string is binary if each character is either 0
 or 1
).

Let 𝑓(𝑡)
 be the decimal representation of integer 𝑡
 written in binary form (possibly with leading zeroes). For example 𝑓(011)=3,𝑓(00101)=5,𝑓(00001)=1,𝑓(10)=2,𝑓(000)=0
 and 𝑓(000100)=4
.

The substring 𝑠𝑙,𝑠𝑙+1,…,𝑠𝑟
 is good if 𝑟−𝑙+1=𝑓(𝑠𝑙…𝑠𝑟)
.

For example string 𝑠=1011
 has 5
 good substrings: 𝑠1…𝑠1=1
, 𝑠3…𝑠3=1
, 𝑠4…𝑠4=1
, 𝑠1…𝑠2=10
 and 𝑠2…𝑠4=011
.

Your task is to calculate the number of good substrings of string 𝑠
.

You have to answer 𝑡
 independent queries.

 ### ideas
 1. good的string的长度假设很长（超过1000）那么它的前缀肯定是0，最后的几位就是s
 2. 且这个长度是固定的，比如 0000111 (长度必须是7，且最后3位是7)
 3. 1, 10, 101, 1000 , 01001, 001010, 
 4. 2e5最多有20位（可能还不用）检查r结尾的20位长，看看是否能组成good字符串，然后加上去即可