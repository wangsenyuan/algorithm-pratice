You are an experienced Codeforces user. Today you found out that during your activity on Codeforces you have made y
submissions, out of which x have been successful. Thus, your current success rate on Codeforces is equal to x / y.

Your favorite rational number in the [0;1] range is p / q. Now you wonder: what is the smallest number of submissions
you have to make if you want your success rate to be p / q?

### ideas

### solution

This problem can be solved using binary search without any special cases. You can also solve it using a formula,
handling a couple of special cases separately.

If our success rate is p / q, it means we have pt successful submissions out of qt for some t. Since p / q is already
irreducible, t has to be a positive integer.

Let's reformulate things a bit. Initially we have x successful submissions and y - x unsuccessful ones. Suppose we fix
t, and in the end we want to have pt successful submissions and qt - pt unsuccessful ones. It's clear we can achieve
that if pt ≥ x and qt - pt ≥ y - x, since we can only increase both the number of successful and unsuccessful
submissions.

Note that both inequalities have constant right hand side, and their left hand side doesn't decrease as t increases.
Therefore, if the inequalities are satisfied for some t, they will definitely be satisfied for larger t as well.

It means we can apply binary search to find the lowest value of t satisfying the inequality. Then, the answer to the
problem is qt - y. If even for very large t the inequalities are not satisfied, the answer is -1.

It can be proved that one "very large" value of t is t = y — that is, if the inequalities are not satisfied for t = y,
then they cannot be satisfied for any value of t. Alternatively, picking t = 109 works too, since y ≤ 109.

Actually, inequalities pt ≥ x and qt - pt ≥ y - x are easy to solve by hand. From pt ≥ x it follows that . From qt -
pt ≥ y - x it follows that . In order to satisfy both inequalities, t has to satisfy .

Don't forget that t has to be integer, so the division results have to be rounded up to the nearest integer. In general,
if we want to find rounded up where both a and b are positive integers, the usual way to do that is to calculate (a +
b - 1)/ b, where "/" is the standard integer division (rounded down).

In this solution, cases when p / q = 0 / 1 or p / q = 1 / 1 have to be handled separately.