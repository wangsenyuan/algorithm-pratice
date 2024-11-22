You are fed up with your messy room, so you decided to clean it up.

Your room is a bracket sequence 𝑠=𝑠1𝑠2…𝑠𝑛
 of length 𝑛
. Each character of this string is either an opening bracket '(' or a closing bracket ')'.

In one operation you can choose any consecutive substring of 𝑠
 and reverse it. In other words, you can choose any substring 𝑠[𝑙…𝑟]=𝑠𝑙,𝑠𝑙+1,…,𝑠𝑟
 and change the order of elements in it into 𝑠𝑟,𝑠𝑟−1,…,𝑠𝑙
.

For example, if you will decide to reverse substring 𝑠[2…4]
 of string 𝑠=
"((()))" it will be equal to 𝑠=
"()(())".

A regular (aka balanced) bracket sequence is a bracket sequence that can be transformed into a correct arithmetic expression by inserting characters '1' and '+' between the original characters of the sequence. For example, bracket sequences "()()", "(())" are regular (the resulting expressions are: "(1)+(1)", "((1+1)+1)"), and ")(" and "(" are not.

A prefix of a string 𝑠
 is a substring that starts at position 1
. For example, for 𝑠=
"(())()" there are 6
 prefixes: "(", "((", "(()", "(())", "(())(" and "(())()".

In your opinion, a neat and clean room 𝑠
 is a bracket sequence that:

the whole string 𝑠
 is a regular bracket sequence;
and there are exactly 𝑘
 prefixes of this sequence which are regular (including whole 𝑠
 itself).
For example, if 𝑘=2
, then "(())()" is a neat and clean room.

You want to use at most 𝑛
 operations to make your room neat and clean. Operations are applied one after another sequentially.

It is guaranteed that the answer exists. Note that you do not need to minimize the number of operations: find any way to achieve the desired configuration in 𝑛
 or less operations.

 ## ideas
 1. (...)()()()
 2. 后面k-1对（）