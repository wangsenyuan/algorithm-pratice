Piegirl got bored with binary, decimal and other integer based counting systems. Recently she discovered some
interesting properties about number , in particular that q2 = q + 1, and she thinks it would make a good base for her
new unique system. She called it "golden system". In golden system the number is a non-empty string containing 0's and
1's as digits. The decimal value of expression a0a1...an equals to .

Soon Piegirl found out that this system doesn't have same properties that integer base systems do and some operations
can not be performed on it. She wasn't able to come up with a fast way of comparing two numbers. She is asking for your
help.

Given two numbers written in golden system notation, determine which of them has larger decimal value.

### ideas

1. 如果a和b的长度相等，很好比较
2. 但是如果不相等

### solution

The important observation one needs to make is that qn = qn - 1 + qn - 2, which means that we can replace two
consecutive ‘1’ digits with one higher significance digit without changing the value. Note that sometimes the next digit
may become more than ‘1’, but that doesn’t affect the solution.

There are two different kinds of solutions for this problem

The first kind of solution involves normalizing both numbers first. The normalization itself can be done in two ways —
from the least significant digit or from the highest significant one using the replacement operation mentioned above. In
either we will need O(n) operations for each number and we then just need to compare them lexicographically.

Other kind of solutions compare numbers digit by digit. We can start from the highest digit of the numbers, and
propagate them to the lower digits. On each step we can do the following:

If both numbers have ones in the highest bit, then we can replace both ones with zeroes, and move on to the next highest
bit.

Now only one number has one in the highest bit. Without loss of generality let’s say it’s the first number. We subtract
one from the highest bit, and add it to the next two highest bits. Now the next two bits of the first number are at
least as big as the first two bits of the second number. Let’s subtract the values of these two bits of the second
number from both first and second number. By doing so we will make the next two bits of the second numbers become 0. If
first number has at least one two, then it is most certainly bigger (because the sum of all the qi for i from 0 to n is
smaller than twice qn + 1). Otherwise we still have only $0$s and $1$s, and can move on to the next highest bit, back to
step (1). Since the ordinal of the highest bit is now smaller, and we only spent constant amount of time, the complexity
of the algorithm is linear.