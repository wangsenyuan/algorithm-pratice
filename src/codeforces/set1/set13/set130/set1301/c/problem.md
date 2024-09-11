Ayoub thinks that he is a very smart person, so he created a function 𝑓(𝑠)
, where 𝑠
 is a binary string (a string which contains only symbols "0" and "1"). The function 𝑓(𝑠)
 is equal to the number of substrings in the string 𝑠
 that contains at least one symbol, that is equal to "1".

More formally, 𝑓(𝑠)
 is equal to the number of pairs of integers (𝑙,𝑟)
, such that 1≤𝑙≤𝑟≤|𝑠|
 (where |𝑠|
 is equal to the length of string 𝑠
), such that at least one of the symbols 𝑠𝑙,𝑠𝑙+1,…,𝑠𝑟
 is equal to "1".

For example, if 𝑠=
"01010" then 𝑓(𝑠)=12
, because there are 12
 such pairs (𝑙,𝑟)
: (1,2),(1,3),(1,4),(1,5),(2,2),(2,3),(2,4),(2,5),(3,4),(3,5),(4,4),(4,5)
.

Ayoub also thinks that he is smarter than Mahmoud so he gave him two integers 𝑛
 and 𝑚
 and asked him this problem. For all binary strings 𝑠
 of length 𝑛
 which contains exactly 𝑚
 symbols equal to "1", find the maximum value of 𝑓(𝑠)
.

Mahmoud couldn't solve the problem so he asked you for help. Can you help him?