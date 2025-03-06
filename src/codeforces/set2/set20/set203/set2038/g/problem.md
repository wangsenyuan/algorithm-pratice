The jury has a string 𝑠
 consisting of characters 0 and/or 1. The length of this string is 𝑛
.

You can ask the following queries:

1
 𝑡
 — "how many times does 𝑡
 appear in 𝑠
 as a contiguous substring?" Here, 𝑡
 should be a string consisting of characters 0 and/or 1; its length should be at least 1
 and at most 𝑛
. For example, if the string 𝑠
 is 111011 and the string 𝑡
 is 11, the response to the query is 3
.
You have to guess at least one character in the string 𝑠
 by asking no more than 3
 queries. Note that giving the answer does not count as a query.

In every test and in every test case, the string 𝑠
 is fixed beforehand.

Interaction
Initially, the jury program sends one integer 𝑡
 (1≤𝑡≤1000
) — the number of test cases.

At the start of each test case, the jury program sends one integer 𝑛
 (2≤𝑛≤50
) — the length of the string.

After that, your program can submit queries to the jury program by printing the following line (do not forget to flush the output after printing a line!):

1
 𝑡
 means asking a query "how many times does 𝑡
 appear in 𝑠
 as a contiguous substring?"
For every query, the jury prints one integer on a separate line. It is either:

the answer to your query, if the query is correct, and you haven't exceeded the query limit;
or the integer −1
, if your query is incorrect (for example, the constraint 1≤|𝑡|≤𝑛
 is not met or the string 𝑡
 contains characters other than 0 and 1) or if you have asked too many queries while processing the current test case.
To submit the answer, your program should send a line in the following format (do not forget to flush the output after printing a line!):

0
 𝑖
 𝑐
, where 1≤𝑖≤𝑛
 and 𝑐
 is either 0 or 1, meaning that 𝑠𝑖=𝑐
.
If your guess is correct, the jury program will print one integer 1
 on a separate line, indicating that you may proceed to the next test case (or terminate the program, if it was the last test case) and that the number of queries you have asked is reset. If it is not correct, the jury program will print one integer −1
 on a separate line.

After your program receives −1
 as the response, it should immediately terminate. This will lead to your submission receiving the verdict "Wrong Answer". If your program does not terminate, the verdict of your submission is undefined.

 ### ideas
 1. 只能猜3次
 2. 如果直接问0， 如果ans = n 或者等于0，=》 直接输出答案
 3. 说明此时既有0也有1
 4. 就是判断s[0]到底是0，还是1
 5. ans的数量表示的是0的个数，n-ans 就是1的个数
 6. 01的个数，是不是更特殊，这个代表的就是是 01转换的数量
 7. 如果= 0， 那么要么全0，要么全1，或者是 1110这种结构
 8. 然后再询问 10, 这两个好像是相关的
 9. 假设第一个是0， 那么 010101 （01 = x, 10 = y)
 10. 0101 x - y = 1
 11. 01010 x - y = 0
 12. 1010 y - x = 1
 13. 10101 y - x = 0
 14. 主要就是 x= y 时，到底是 10101, 还是 01010 (x != 0 and y != 0)
 15. 100001001, 1110111
 16. x = 0 and y = 0, 再问一次 0就可以了
 17. x != 0 and y != 0, and x = y
 18. 问一下 11， z = 11的个数 = 10 10 10 这样的，大概3段（不算最后一段）y, 
 19. z = 所有1的个数 - y - (1 or 0 看最后一位是不是1)
 20. 我们假设最后一个是1
 21. z = c1 - y - 1 => c1 = z + y + 1
 22. 1101101  (y = 2, z = 2, c1 = 5)
 23. 第一步询问1的个数得到c1
 24. 如果 c1 = n, or c1 = 0, 都很好处理
 25. 然后询问 10 的个数 y
 26. 然后询问 11 的个数 z
 27. 然后判断给出最后一位的答案