Suppose you have a special 𝑥
-𝑦
-counter. This counter can store some value as a decimal number; at first, the counter has value 0
.

The counter performs the following algorithm: it prints its lowest digit and, after that, adds either 𝑥
 or 𝑦
 to its value. So all sequences this counter generates are starting from 0
. For example, a 4
-2
-counter can act as follows:

it prints 0
, and adds 4
 to its value, so the current value is 4
, and the output is 0
;
it prints 4
, and adds 4
 to its value, so the current value is 8
, and the output is 04
;
it prints 8
, and adds 4
 to its value, so the current value is 12
, and the output is 048
;
it prints 2
, and adds 2
 to its value, so the current value is 14
, and the output is 0482
;
it prints 4
, and adds 4
 to its value, so the current value is 18
, and the output is 04824
.
This is only one of the possible outputs; for example, the same counter could generate 0246802468024
 as the output, if we chose to add 2
 during each step.

You wrote down a printed sequence from one of such 𝑥
-𝑦
-counters. But the sequence was corrupted and several elements from the sequence could be erased.

Now you'd like to recover data you've lost, but you don't even know the type of the counter you used. You have a decimal string 𝑠
 — the remaining data of the sequence.

For all 0≤𝑥,𝑦<10
, calculate the minimum number of digits you have to insert in the string 𝑠
 to make it a possible output of the 𝑥
-𝑦
-counter. Note that you can't change the order of digits in string 𝑠
 or erase any of them; only insertions are allowed.



### ideas
1. 对于确定的(x, y)
2. dp[i][j] 表示到i为止，且最后一个print出来的是digit j时的最少删除字符
3. dp[i][j] = dp[i-1][k] 如果 (k + x) % 10 = j or (k + y) % 10 = j 且s[i] = j
4. 