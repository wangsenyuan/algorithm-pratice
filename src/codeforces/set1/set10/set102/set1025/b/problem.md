During the research on properties of the greatest common divisor (GCD) of a set of numbers, Ildar, a famous
mathematician, introduced a brand new concept of the weakened common divisor (WCD) of a list of pairs of integers.

For a given list of pairs of integers (𝑎1,𝑏1)
, (𝑎2,𝑏2)
, ..., (𝑎𝑛,𝑏𝑛)
their WCD is arbitrary integer greater than 1
, such that it divides at least one element in each pair. WCD may not exist for some lists.

For example, if the list looks like [(12,15),(25,18),(10,24)]
, then their WCD can be equal to 2
, 3
, 5
or 6
(each of these numbers is strictly greater than 1
and divides at least one number in each pair).

You're currently pursuing your PhD degree under Ildar's mentorship, and that's why this problem was delegated to you.
Your task is to calculate WCD efficiently.

### ideas

1. wcd = 至少整除pair中的一个数