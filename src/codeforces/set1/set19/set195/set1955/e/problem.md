A binary string 𝑠
 of length 𝑛
 is given. A binary string is a string consisting only of the characters '1' and '0'.

You can choose an integer 𝑘
 (1≤𝑘≤𝑛
) and then apply the following operation any number of times: choose 𝑘
 consecutive characters of the string and invert them, i.e., replace all '0' with '1' and vice versa.

Using these operations, you need to make all the characters in the string equal to '1'.

For example, if 𝑛=5
, 𝑠=00100
, you can choose 𝑘=3
 and proceed as follows:

choose the substring from the 1
-st to the 3
-rd character and obtain 𝑠=11000
;
choose the substring from the 3
-rd to the 5
-th character and obtain 𝑠=11111
;
Find the maximum value of 𝑘
 for which it is possible to make all the characters in the string equal to '1' using the described operations. Note that the number of operations required to achieve this is not important.

### ideas
1. 显然k = 1 => ok 的 
2. 如果 所有的相同，那么 k = n也是ok的
3. 对于给定的k，如何check它是否ok呢？
4. 考虑最后一次，它必须是连续的k个相同的字符（0/1）
5. 每次变化里面，不变的是什么？
6. 10001 =》 01110 
7. 就brute froce from left