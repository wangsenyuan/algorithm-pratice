Denis, after buying flowers and sweets (you will learn about this story in the next task), went to a date with Nastya to ask her to become a couple. Now, they are sitting in the cafe and finally... Denis asks her to be together, but ... Nastya doesn't give any answer.

The poor boy was very upset because of that. He was so sad that he punched some kind of scoreboard with numbers. The numbers are displayed in the same way as on an electronic clock: each digit position consists of 7
 segments, which can be turned on or off to display different numbers. The picture shows how all 10
 decimal digits are displayed:


After the punch, some segments stopped working, that is, some segments might stop glowing if they glowed earlier. But Denis remembered how many sticks were glowing and how many are glowing now. Denis broke exactly 𝑘
 segments and he knows which sticks are working now. Denis came up with the question: what is the maximum possible number that can appear on the board if you turn on exactly 𝑘
 sticks (which are off now)?

It is allowed that the number includes leading zeros.

### ideas
1. 就是再找k个不亮的字符，和现有的亮的字符，组成有效的数字
2. dp[i][j] = true表示，可以用j个坏的stick，组成一个有效数字
3. 那么在第i位时，在保证dp[i+1][?]为true的时候，让当前位最大