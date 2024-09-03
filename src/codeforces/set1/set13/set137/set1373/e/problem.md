Let 𝑓(𝑥)
 be the sum of digits of a decimal number 𝑥
.

Find the smallest non-negative integer 𝑥
 such that 𝑓(𝑥)+𝑓(𝑥+1)+⋯+𝑓(𝑥+𝑘)=𝑛
.

Input
The first line contains one integer 𝑡
 (1≤𝑡≤150
) — the number of test cases.

Each test case consists of one line containing two integers 𝑛
 and 𝑘
 (1≤𝑛≤150
, 0≤𝑘≤9
).

Output
For each test case, print one integer without leading zeroes. If there is no such 𝑥
 that 𝑓(𝑥)+𝑓(𝑥+1)+⋯+𝑓(𝑥+𝑘)=𝑛
, print −1
; otherwise, print the minimum 𝑥
 meeting that constraint.


 ### ideas
 1. 看起来这个x会非常大
 2. 另外，应该是不符合2分性质的
 3. 如果全部是9, 且 n % 9 = 0, 那么很简单就是 n / 9 个9（不会有比这更小的）
 4. dp[i][j]表示用i位，sum = j时的最小值（没有leading 0)
 5. dp[i][j] = dp[i-1][j-?] (? 0 ~ 9)
 6. 这个式子不对。
 7. s(x) = f(x) + f(x + 1) + f(x + 2)... + f(x + k)
 8. k <= 9
 9. 如果x的最低位w, 如果 w + k <= 9, 
 10. 那么s(x) = k * f(x) - k * w + w + w + 1 + w... + w + k
 11. 如果w + k > 9, 也能处理，影响的只是最后两位而已
 12. g(w) 表示以w为结尾时，能够得到的最小的x（固定最后一个数），从而知道f(x)应该是多少。
 13. 要看最后两位的