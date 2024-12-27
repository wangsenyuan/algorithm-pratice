Vus the Cossack has two binary strings, that is, strings that consist only of "0" and "1". We call these strings 𝑎
 and 𝑏
. It is known that |𝑏|≤|𝑎|
, that is, the length of 𝑏
 is at most the length of 𝑎
.

The Cossack considers every substring of length |𝑏|
 in string 𝑎
. Let's call this substring 𝑐
. He matches the corresponding characters in 𝑏
 and 𝑐
, after which he counts the number of positions where the two strings are different. We call this function 𝑓(𝑏,𝑐)
.

For example, let 𝑏=00110
, and 𝑐=11000
. In these strings, the first, second, third and fourth positions are different.

Vus the Cossack counts the number of such substrings 𝑐
 such that 𝑓(𝑏,𝑐)
 is even.

For example, let 𝑎=01100010
 and 𝑏=00110
. 𝑎
 has four substrings of the length |𝑏|
: 01100
, 11000
, 10001
, 00010
.

𝑓(00110,01100)=2
;
𝑓(00110,11000)=4
;
𝑓(00110,10001)=4
;
𝑓(00110,00010)=1
.
Since in three substrings, 𝑓(𝑏,𝑐)
 is even, the answer is 3
.

Vus can not find the answer for big strings. That is why he is asking you to help him.


### ideas
1. 假设已经计算好了第一个字符串的结果
2. 现在计算第二个字符串
3. f(b, x1) = w （b[i] != x1[i] 的个数）
4. f(b, x2) = ？， 假设b[i] = 0, a[i] = 1 (+1), a[i+1] = 1 (w不变)
5. 但是如果 a[i+1] = 0, 那么 w - 1, 同理，如果b[i] = 1, a[i] 和 a[i+1] 相同，w不变，否则w变化+/-1
6. 而且这里关心的奇偶行，也不需要知道+/1, 只需要知道变化1即可
7. 头尾不需要特殊处理
8. 如果 a[i] == b[i]， 但是a[i] != a[i+1] diff++
9. 如果 a[i] != b[i], 但是 a[i] != a[i+1], diff-- 