One of the first programming problems by K1o0n looked like this: "Noobish_Monk has 𝑛
 (1≤𝑛≤100)
 friends. Each of them gave him 𝑎
 (1≤𝑎≤10000)
 apples for his birthday. Delighted with such a gift, Noobish_Monk returned 𝑏
 (1≤𝑏≤min(10000,𝑎⋅𝑛))
 apples to his friends. How many apples are left with Noobish_Monk?"

K1o0n wrote a solution, but accidentally considered the value of 𝑛
 as a string, so the value of 𝑛⋅𝑎−𝑏
 was calculated differently. Specifically:

when multiplying the string 𝑛
 by the integer 𝑎
, he will get the string 𝑠=𝑛+𝑛+⋯+𝑛+𝑛𝑎 times
when subtracting the integer 𝑏
 from the string 𝑠
, the last 𝑏
 characters will be removed from it. If 𝑏
 is greater than or equal to the length of the string 𝑠
, it will become empty.
Learning about this, ErnKor became interested in how many pairs (𝑎,𝑏)
 exist for a given 𝑛
, satisfying the constraints of the problem, on which K1o0n's solution gives the correct answer.

"The solution gives the correct answer" means that it outputs a non-empty string, and this string, when converted to an integer, equals the correct answer, i.e., the value of 𝑛⋅𝑎−𝑏
.

### ideas
1. n * a - b = (nnnnn) - "b"
2. 找规律
3. 如果n是一位数字，（1.。。9）
4. f(1, a, b) = 1111 - b = 111 (删掉后面|b|个1)
5. 但是 g(1, a, b) = 1 * a - b = 111
6. a = b + (111) => s(1, a) 至少 (111)个1
7. f(1, 3, 5) = 111 - "5" = 11
8. g(1, 3, 5) = 3 - 5 = -2
9. f(1, 2, ?) = 11 - "?" = 1
10. g(1, 2, ?) = 2 - ? = 1 => ? = 1
11. f(1, 2, 1) = 11 - "1" = 1
12. 当a = 1 时， f(n, a, b) = "n" - "b" = x
13. g(n, a, b) = n - b = x
14. x既是 n * a - b 的结果，也是 repeat(n, a) 的一个前缀
15. 