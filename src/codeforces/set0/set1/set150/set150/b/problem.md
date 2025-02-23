Just in case somebody missed it: this winter is totally cold in Nvodsk! It is so cold that one gets funny thoughts. For example, let's say there are strings with the length exactly n, based on the alphabet of size m. Any its substring with length equal to k is a palindrome. How many such strings exist? Your task is to find their quantity modulo 1000000007 (109 + 7). Be careful and don't miss a string or two!

Let us remind you that a string is a palindrome if it can be read the same way in either direction, from the left to the right and from the right to the left.

### ideas
1. 对于长度为n的，所有的长度为k的子串都是回文
2. 比如说n = 10, k = 4
3. abbaabba...
4. a[i] = a[i+k-1] 要成立
5. a[i+1] = a[i+k-2] 也要成立
6. abccbaabc， 只有一种情况 ababababa （或者全是a）
7. 