While Vasya finished eating his piece of pizza, the lesson has already started. For being late for the lesson, the
teacher suggested Vasya to solve one interesting problem. Vasya has an array a and integer x. He should find the number
of different ordered pairs of indexes (i, j) such that ai ≤ aj and there are exactly k integers y such that ai ≤ y ≤ aj
and y is divisible by x.

In this problem it is meant that pair (i, j) is equal to (j, i) only if i is equal to j. For example pair (1, 2) is not
the same as (2, 1).

### ideas

1. 先理解题目
2. ai <= y <= aj, 且y能够被x整除
3. 先考虑 ai = aj 的情况，如果 ai % x = 0, 那么就是1， else 0
4. 然后考虑ai < aj， cnt = aj / x - (ai - 1) / x ?
5. 