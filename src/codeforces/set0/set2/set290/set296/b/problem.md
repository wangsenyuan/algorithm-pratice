Yaroslav thinks that two strings s and w, consisting of digits and having length n are non-comparable if there are two
numbers, i and j (1 ≤ i, j ≤ n), such that si>wi and sj<wj. Here sign si represents the i-th digit of string s,
similarly, wj represents the j-th digit of string w.

A string's template is a string that consists of digits and question marks ("?").

Yaroslav has two string templates, each of them has length n. Yaroslav wants to count the number of ways to replace all
question marks by some integers in both templates, so as to make the resulting strings incomparable. Note that the
obtained strings can contain leading zeroes and that distinct question marks can be replaced by distinct or the same
integers.

Help Yaroslav, calculate the remainder after dividing the described number of ways by 1000000007 (109 + 7).

### ideas

1. dp on state