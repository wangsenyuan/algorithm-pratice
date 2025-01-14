Heidi is now just one code away from breaking the encryption of the Death Star plans. The screen that should be presenting her with the description of the next code looks almost like the previous one, though who would have thought that the evil Empire engineers would fill this small screen with several million digits! It is just ridiculous to think that anyone would read them all...

Heidi is once again given a sequence A and two integers k and p. She needs to find out what the encryption key S is.

Let X be a sequence of integers, and p a positive integer. We define the score of X to be the sum of the elements of X modulo p.

Heidi is given a sequence A that consists of N integers, and also given integers k and p. Her goal is to split A into k parts such that:

Each part contains at least 1 element of A, and each part consists of contiguous elements of A.
No two parts overlap.
The total sum S of the scores of those parts is minimized (not maximized!).
Output the sum S, which is the encryption code.

### ideas
1. 考虑c2的算法为啥是不行的？
2. 是可行的。但是会超时。
3. 