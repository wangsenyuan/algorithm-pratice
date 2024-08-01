You are given a string 𝑠
 of length 𝑛>1
, consisting of digits from 0
 to 9
. You must insert exactly 𝑛−2
 symbols +
 (addition) or ×
 (multiplication) into this string to form a valid arithmetic expression.

In this problem, the symbols cannot be placed before the first or after the last character of the string 𝑠
, and two symbols cannot be written consecutively. Also, note that the order of the digits in the string cannot be changed. Let's consider 𝑠=987009
:

From this string, the following arithmetic expressions can be obtained: 9×8+70×0+9=81
, 98×7×0+0×9=0
, 9+8+7+0+09=9+8+7+0+9=33
. Note that the number 09
 is considered valid and is simply transformed into 9
.
From this string, the following arithmetic expressions cannot be obtained: +9×8×70+09
 (symbols should only be placed between digits), 98×70+0+9
 (since there are 6
 digits, there must be exactly 4
 symbols).
The result of the arithmetic expression is calculated according to the rules of mathematics — first all multiplication operations are performed, then addition. You need to find the minimum result that can be obtained by inserting exactly 𝑛−2
 addition or multiplication symbols into the given string 𝑠
.

### ideas
1. 使用n-2个+/*, 也就是说中间的所有字符前，都必须有一个操作符，除了一个
2. 所以这里选中i，在i和i+1中间不适用操作符，其他的都需要时，在左右进行计算后
3. 能够获得的最小值