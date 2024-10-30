Squirrel Liss is interested in sequences. She also has preferences of integers. She thinks n integers a1, a2, ..., an are good.

Now she is interested in good sequences. A sequence x1, x2, ..., xk is called good if it satisfies the following three conditions:

The sequence is strictly increasing, i.e. xi < xi + 1 for each i (1 ≤ i ≤ k - 1).
No two adjacent elements are coprime, i.e. gcd(xi, xi + 1) > 1 for each i (1 ≤ i ≤ k - 1) (where gcd(p, q) denotes the greatest common divisor of the integers p and q).
All elements of the sequence are good integers.
Find the length of the longest good sequence.

### ideas
1. a[i] is already sorted