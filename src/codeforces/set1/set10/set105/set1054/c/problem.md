There are 𝑛
children numbered from 1
to 𝑛
in a kindergarten. Kindergarten teacher gave 𝑎𝑖
(1≤𝑎𝑖≤𝑛
) candies to the 𝑖
-th child. Children were seated in a row in order from 1
to 𝑛
from left to right and started eating candies.

While the 𝑖
-th child was eating candies, he calculated two numbers 𝑙𝑖
and 𝑟𝑖
— the number of children seating to the left of him that got more candies than he and the number of children seating to
the right of him that got more candies than he, respectively.

Formally, 𝑙𝑖
is the number of indices 𝑗
(1≤𝑗<𝑖
), such that 𝑎𝑖<𝑎𝑗
and 𝑟𝑖
is the number of indices 𝑗
(𝑖<𝑗≤𝑛
), such that 𝑎𝑖<𝑎𝑗
.

Each child told to the kindergarten teacher the numbers 𝑙𝑖
and 𝑟𝑖
that he calculated. Unfortunately, she forgot how many candies she has given to each child. So, she asks you for help:
given the arrays 𝑙
and 𝑟
determine whether she could have given the candies to the children such that all children correctly calculated their
values 𝑙𝑖
and 𝑟𝑖
, or some of them have definitely made a mistake. If it was possible, find any way how she could have done it.

### ideas

1. 如果 l[i] = 0, 那么 他的左边的都比他小
2. 如果a[i]都不一样，那么l[i] + r[i]最大的数，就是最小的数
3. 现在考虑有相等的数，a[i] = a[j], i < j
4. 貌似还是 l[i] + r[i] = l[j] + r[j] 还是成立的