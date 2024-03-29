You are given a permutation of integers from 1 to n. Exactly once you apply the following operation to this permutation:
pick a random segment and shuffle its elements. Formally:

Pick a random segment (continuous subsequence) from l to r. All segments are equiprobable.
Let k = r - l + 1, i.e. the length of the chosen segment. Pick a random permutation of integers from 1 to k, p1,
p2, ..., pk. All k! permutation are equiprobable.
This permutation is applied to elements of the chosen segment, i.e. permutation a1, a2, ..., al - 1, al, al + 1, ...,
ar - 1, ar, ar + 1, ..., an is transformed to a1, a2, ..., al - 1, al - 1 + p1, al - 1 + p2, ..., al - 1 + pk - 1, al -
1 + pk, ar + 1, ..., an.
Inversion if a pair of elements (not necessary neighbouring) with the wrong relative order. In other words, the number
of inversion is equal to the number of pairs (i, j) such that i<j and ai>aj. Find the expected number of inversions
after we apply exactly one operation mentioned above.

### solution

Lets calculate all available segments count – . It will be a denominator of answer fraction.

Also, necessary to understand, that expected value of inversion count in shuffled permutation with length len equal  (It
can be prooved by fact, that for each permutation there are mirrored permutation (or biection i ->n - i + 1) and sum of
inversion count his permutations are equal inversions).

Next, we will find expected value of difference between expected value of inversion count after operation and inversion
count in the source array. And add it to the inversion count in the source array.

Naive solution:

For every segment we will calculate the count of inversions in it, and also the expected value of the inversion count in
it after shuffle. Take the difference and divide by a denominator. Sum these values for all segments. This solution has
quadratic asympthotics, try to improve it.

Optimized solution:

We will go through the permutation from right to left and for each position count the sum of inversions in the initial
permutation for all segments that start in the current position (denote that the largest segment which ends at the
position n has the length len), also we will maintain the sum of expected values of the inversion counts on the segments
of lengths 1..len. Knowing these two numbers, increase the answer by their difference divided by the denominator.

To calculate the first value we will use the data structure that can get the sum of numbers on the prefix and modify the
value at the single position (e.g. Fenwick tree). For the position i we need to know how many numbers are less than ai,
and every number should be taken t times, where t is number of segments where it is contained. Suppose we have
calculated answers for some suffix and are now standing at the position i. For every position to the left it will be
added n - i + 1 times (the number of positions j>= i as the candidates for the right bound, in 1-indexation). Perform
fenwick add(a[i], n - i + 1).