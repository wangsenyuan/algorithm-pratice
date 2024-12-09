You are given a sequence of 𝑛
 digits 𝑑1𝑑2…𝑑𝑛
. You need to paint all the digits in two colors so that:

each digit is painted either in the color 1
 or in the color 2
;
if you write in a row from left to right all the digits painted in the color 1
, and then after them all the digits painted in the color 2
, then the resulting sequence of 𝑛
 digits will be non-decreasing (that is, each next digit will be greater than or equal to the previous digit).
For example, for the sequence 𝑑=914
 the only valid coloring is 211
 (paint in the color 1
 two last digits, paint in the color 2
 the first digit). But 122
 is not a valid coloring (9
 concatenated with 14
 is not a non-decreasing sequence).

It is allowed that either of the two colors is not used at all. Digits painted in the same color are not required to have consecutive positions.

Find any of the valid ways to paint the given sequence of digits or determine that it is impossible to do.

### ideas
1. 如果将当前字符paint为1，那么前面所有颜色为1的，都必须比<= s[i], 且前面所有2的都必须 >= s[i]
2. 似乎可以把某个 <= x都变成1