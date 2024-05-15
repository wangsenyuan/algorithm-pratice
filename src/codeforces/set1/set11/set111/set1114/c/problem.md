score in a game of tennis.
Aki is fond of numbers, especially those with trailing zeros. For example, the number 9200
 has two trailing zeros. Aki thinks the more trailing zero digits a number has, the prettier it is.

However, Aki believes, that the number of trailing zeros of a number is not static, but depends on the base (radix) it is represented in. Thus, he considers a few scenarios with some numbers and bases. And now, since the numbers he used become quite bizarre, he asks you to help him to calculate the beauty of these numbers.

Given two integers 𝑛
 and 𝑏
 (in decimal notation), your task is to calculate the number of trailing zero digits in the 𝑏
-ary (in the base/radix of 𝑏
) representation of 𝑛!
 (factorial of 𝑛
).

### ideas
1. 当 b = 10的时候，有个标准的做法是，计算它里面包含多少个5，减去25的倍数 + 125的倍数。。。
2. 当 b 为其他数时，应该怎么计算呢？b可以因数分解，比如 9可以分为3 * 3, 8 = 2 * 4
3. 然后计算其中多的那个数字的倍数
4. 当b = 2的时候，要特殊处理