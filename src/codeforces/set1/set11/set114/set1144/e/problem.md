You are given two strings 𝑠
 and 𝑡
, both consisting of exactly 𝑘
 lowercase Latin letters, 𝑠
 is lexicographically less than 𝑡
.

Let's consider list of all strings consisting of exactly 𝑘
 lowercase Latin letters, lexicographically not less than 𝑠
 and not greater than 𝑡
 (including 𝑠
 and 𝑡
) in lexicographical order. For example, for 𝑘=2
, 𝑠=
"az" and 𝑡=
"bf" the list will be ["az", "ba", "bb", "bc", "bd", "be", "bf"].

Your task is to print the median (the middle element) of this list. For the example above this will be "bc".

It is guaranteed that there is an odd number of strings lexicographically not less than 𝑠
 and not greater than 𝑡
.

### ideas
1. 给定一个字符串，可以计算它的位置
2. 这里可以先根据s/t计算出它们的位置，然后找median位置的字符串
3. 但是k长度的（1e5)，且26进制（即使是2进制）这个数量也太大了。
4. 考虑10进制小，1,2,3的中位数是2, 1,2,3,4,5的中位数是3
5. 如果最后一个数是奇数 -1， /2。 如果是偶数/2
6. 然后考虑   3,4,5的中位数是4
7. 13，14，15的中位数是14
8. 所以是最后两位相加/2