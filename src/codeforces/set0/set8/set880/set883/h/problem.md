Kolya has a string s of length n consisting of lowercase and uppercase Latin letters and digits.

He wants to rearrange the symbols in s and cut it into the minimum number of parts so that each part is a palindrome and
all parts have the same lengths. A palindrome is a string which reads the same backward as forward, such as madam or
racecar.

Your task is to help Kolya and determine the minimum number of palindromes of equal lengths to cut s into, if it is
allowed to rearrange letters in s before cuttings.

### ideas

1. 显然f(s) = n 是一个有效的答案
2. 假设最优的答案是x， 也就是有x个回文，那么有 n % x = 0, y = n / x 是每个回文的长度
3. 假设cnt[i] 表示字符i的数量
4. 如果cnt[i]是偶数, 且 cnt[i]/2 <= x，那么它似乎对结果是没有影响的, 因为给其中的 cnt[i]/2 中每个分配2个
5. 