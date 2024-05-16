You are in a nuclear laboratory that is about to explode and destroy the Earth. You must save the Earth before the final countdown reaches zero.

The countdown consists of 𝑛
 (1≤𝑛≤4⋅105
) mechanical indicators, each showing one decimal digit. You noticed that when the countdown changes its state from 𝑥
 to 𝑥−1
, it doesn't happen in one move. Instead, each change of a single digit takes one second.

So, for example, if the countdown shows 42, then it will change to 41 in one second, because only one digit is changed, but if the countdown shows 2300, then it will change to 2299 in three seconds, because the three last digits are changed.

Find out how much time is left before the countdown reaches zero.

Input
The first line of input contains a single integer 𝑡
 (1≤𝑡≤104
) — the number of test cases. Then the descriptions of the test cases follow.

The first line of each test case contains a single integer 𝑛
 (1≤𝑛≤4⋅105
).

The second line contains a string of 𝑛
 digits, the current state of the countdown. It is guaranteed that at least one digit is not zero.

The sum of 𝑛
 for all tests does not exceed 4⋅105
.

Output
For each test case, print a single integer without leading zeroes, the number of seconds left before the countdown reaches zero. Note that this number may be huge.

### ideas
1. 很有意思
2. 个要减去1，只需要1秒钟
3. 如果最后一位为0，那么只能更改十位， 比如10， 那么它需要两秒变成 09
4. 然后再花费9秒，把9减到0
5. 10位数字多花了一秒
6. 百位应该是多花2秒
7. 