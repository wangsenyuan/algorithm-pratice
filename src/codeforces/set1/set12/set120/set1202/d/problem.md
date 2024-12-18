The subsequence is a sequence that can be derived from another sequence by deleting some elements without changing the order of the remaining elements.

You are given an integer 𝑛
.

You have to find a sequence 𝑠
 consisting of digits {1,3,7}
 such that it has exactly 𝑛
 subsequences equal to 1337
.

For example, sequence 337133377
 has 6
 subsequences equal to 1337
:

3371⎯⎯33⎯⎯3⎯⎯77⎯⎯
 (you can remove the second and fifth characters);
3371⎯⎯3⎯⎯33⎯⎯77⎯⎯
 (you can remove the third and fifth characters);
3371⎯⎯3⎯⎯3⎯⎯377⎯⎯
 (you can remove the fourth and fifth characters);
3371⎯⎯33⎯⎯3⎯⎯7⎯⎯7
 (you can remove the second and sixth characters);
3371⎯⎯3⎯⎯33⎯⎯7⎯⎯7
 (you can remove the third and sixth characters);
3371⎯⎯3⎯⎯3⎯⎯37⎯⎯7
 (you can remove the fourth and sixth characters).

 ### ideas
 1. 最简单的方式，是不是 1 3333 7 333 7 
 2. ni = n3 * (n3 - 1) / 2 (如果这个成立，可以用这个方式弄出来)
 3. 所以还得想其他的办法
 4. 假设最后一个   37 它前面有x[k]个3，已经1个1 =》x[k]
 5. 然后在某个位置放置一个 1 3737373737
 6. 13737373737
 7. 对于每一个位置 = 前面3的数量 * 后面7的数量 的和
 8. 如果n = a * b 满足 a + b + 1 <= 1e5 bigo
 9. s = 13(a)7(b)
 10. m = (a + 1) * a / 2 * b <= n （任何一个7都和前面的两个3进行匹配）
 11. 假设再添加37 ， 对于最后一个7来说，增加的是 (a + 2) * (a + 1) / 2
 12. 但是问题是，这个添加的有可能会超了
 13. 所以不大对
 14. 13(a)7
 15. m = (a + 1) * a / 2
 16. 如果 m != n， n -= m
 17. 然后在找出一个位置，放置一个7 直到剩下两个3
 18. 1337....7....7....7 
 19. 必须知道一开始的位置
 20. 