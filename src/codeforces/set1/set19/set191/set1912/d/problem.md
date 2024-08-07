Daisy has recently learned divisibility rules for integers and she is fascinated by them. One of the tests she learned is a divisibility test by 3. You can find a sum of all digits of a decimal number and check if the resulting sum is divisible by 3. Moreover, the resulting sum of digits is congruent modulo 3 to the original number — the remainder modulo 3 is preserved. For example, 75≡7+5(mod3)
. Daisy is specifically interested in such remainder preserving divisibility tests.

There are more examples like that that she learned for decimal integers (integers base 10):

To test divisibility modulo 11, find an alternating sum of digits. Counting digits from the last (least significant) digit, add digits on odd positions (the last, 3rd to the last, etc) and subtract digits on even positions (2nd to the last, 4th to the last, etc) to get the sum that is congruent modulo 11 to the original number. For example, 123≡1−2+3(mod11)
.
To test divisibility modulo 4, keep the last two digits. Their value is congruent modulo 4 to the original number. For example, 876543≡43(mod4)
.
To test divisibility modulo 7, find an alternating sum of groups of three digits. For example, 4389328≡4−389+328(mod7)
.
Similar tests can be found in other bases. For example, to test divisibility modulo 5 for octal numbers (base 8), find an alternating sum of groups of two digits. For example, 12348≡−128+348(mod5)
.

Daisy wants to find such rules for a given base 𝑏
. She is interested in three kinds of divisibility rules:

Kind 1 — take the last 𝑘
 digits of an integer in base 𝑏
.
Kind 2 — take a sum of groups of 𝑘
 digits of an integer in base 𝑏
.
Kind 3 — take an alternating sum of groups of 𝑘
 digits of an integer in base 𝑏
.
It is not always possible to find such a divisibility rule. For example, in base 10 there is no such test for divisibility modulo 6, even though different approaches to testing divisibility by 6 exist.

Given base 𝑏
 and modulo 𝑛
, Daisy wants to know the smallest group size 𝑘
 for which such a divisibility test exits.


 ### ideas
 1. 在给定base的情况下，如何计算一个num(base) % n(base)呢？
 2. rem = num - (n) 一直到 rem < n为止