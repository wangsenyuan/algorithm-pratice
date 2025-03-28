In his very young years the hero of our story, king Copa, decided that his private data was hidden not enough securely, what is unacceptable for the king. That's why he invented tricky and clever password (later he learned that his password is a palindrome of odd length), and coded all his data using it.

Copa is afraid to forget his password, so he decided to write it on a piece of paper. He is aware that it is insecure to keep password in such way, so he decided to cipher it the following way: he cut x characters from the start of his password and from the end of it (x can be 0, and 2x is strictly less than the password length). He obtained 3 parts of the password. Let's call it prefix, middle and suffix correspondingly, both prefix and suffix having equal length and middle always having odd length. From these parts he made a string A + prefix + B + middle + C + suffix, where A, B and C are some (possibly empty) strings invented by Copa, and « + » means concatenation.

Many years have passed, and just yesterday the king Copa found the piece of paper where his ciphered password was written. The password, as well as the strings A, B and C, was completely forgotten by Copa, so he asks you to find a password of maximum possible length, which could be invented, ciphered and written by Copa.

###

输入长度 ≤1e5 的字符串 s，只包含小写英文字母。

有一个奇回文串 t，设 t = pre + mid + suf，其中 mid 的长度是奇数，pre 和 suf 的长度相等（可以为 0）。
已知 s = A + pre + B + mid + C + suf（A B C 可以为空），求 t 的最大长度。

### ideas
1. 如果B和C都是空的，那么就是一段最长的奇数长度的回文串
2. 现在假设B不为空，那么pre = reverse(suf), 可以用kmp计算出，每个位置i, 能够匹配的最长的suf
3. 然后再计算对于每个位置c，计算以它为中心的回文串(l, r)的最大值？ 这个地方还有点问题
4. 就是pre有可能被中心的回文给覆盖到。如果被覆盖，其实肯定是第一种情况吗？
5. 好像也不一定， 考虑 abcdcbaxba 
6. 中心是cdc, pre = ab, suf = ba
7. 这里是不是可以直接忽略这种情况？假设pre被包含进去了，那么它肯定不会产生3段？
8. A + pre + B + mid + C + suf
9. 在固定mid的情况下，如何保证suf和mid没有交集
10. 对于每个suf，固定它最靠近左边的前缀位置, 且这个前缀位置肯定是递增的
11. 那么就是在这个前缀位置中，二分查找 i - l/2 的地方就可以了
12. aaa B aaaaa aaa
13. aaa aaaaa C aaa