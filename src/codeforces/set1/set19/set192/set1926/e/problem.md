Vladislav has 𝑛
 cards numbered 1,2,…,𝑛
. He wants to lay them down in a row as follows:

First, he lays down all the odd-numbered cards from smallest to largest.
Next, he lays down all cards that are twice an odd number from smallest to largest (i.e. 2
 multiplied by an odd number).
Next, he lays down all cards that are 3
 times an odd number from smallest to largest (i.e. 3
 multiplied by an odd number).
Next, he lays down all cards that are 4
 times an odd number from smallest to largest (i.e. 4
 multiplied by an odd number).
And so on, until all cards are laid down.
What is the 𝑘
-th card he lays down in this process? Once Vladislav puts a card down, he cannot use that card again.


### ideas
1. 模拟一下
2. n = 10
3. 1, 3, 5, 7, 9, 2, 6, 10, 3, 9, 4, 5, 6, .... 10
4. 变成矩阵 
5. 1，3, 5, 7, 9
6. 2, 6, 10
7. 3, 9
8. 4,
9. ...
10. 二分查找对于num，比它小的数有多少个？
11. 不行， 问题是这个num不是递增的
12. 先找到k所在的行？
13. 可以在sqrt(k)次找到答案
14. 这里隐藏着一个重要的信息是，所有的card只能被使用一次
15. 1， 3， 5， 7， 9, 11, 13
16. 2, 6, 10, 14 18, 22, 26
17.       15  21  27, 33, 39
18. 4  12, 20, 28, 36, .. 52
19.       25
20. 先不考虑同一个数字被使用了多次的情况
21. 1, 3, 5, 7, 9, 11, 13, 15  (n = 16)
22. 2, 6, 10, 14
23. 3, 9, 15
24. 4, 12 
25. 5, 15
26. 6 ...
27. 然后对于第i行，如何计算它们出现了多少个数字呢?(出现多次算一个)
28. 比如上面的第4行，共出现了 8 + 4 + 3 + 2 - 1(3) - 1 (9) - 1（15） = 14个数 （除了8, 和 16没有出现）
29. 考虑一个数x，它能出现几次，假设它的因数是 x1, x2, x3, .... x1 = 1
30. 如果它是奇数，那么它出现在第一行
31. 如果它 = 2 * 某个奇数 ， 那么它出现在第二行
32. 如果它 = 3 * 某个奇数 (这个不可能) 这样子，它必然已经是一个奇数了 
33. 如果它 = 4 * 某个奇数 (那么它肯定出现在第4行)
34. 如果它 = 5 * 奇数 （也不可能）
35. 6 * 某个奇数  
36. cnt[i] 表示第i层出现的个数，如果不考虑之前已经出现的情况 
37. 假设 i * x <= n => x <= n / i, (x + 1) / 2  cnt[i] = (n / i + 1) / 2
38. - 前面出现过的📖， 假设这一行的某个数 w = i * x, 如果 w = j * y 且 j < i, 且y是个奇数，
39. i * x = j * y => i 是j的倍数 =》y / x = i / j
40. cnt[i] -= cnt[j] 如果 i % j = 0?
41. 肯定不对的, 因为cnt[j] 肯定 > cnt[i]
42. 1, 3, 5, 7, 9, ...
43. 2, 6, 10, 14, ...
44. 4, 12, 20, 28, 
45. 8, 24, ...
46. 
