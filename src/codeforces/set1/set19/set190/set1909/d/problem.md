There are 𝑛
positive integers 𝑎1,𝑎2,…,𝑎𝑛
on a blackboard. You are also given a positive integer 𝑘
. You can perform the following operation some (possibly 0
) times:

choose a number 𝑥
on the blackboard;
erase one occurrence of 𝑥
;
write two positive integers 𝑦
, 𝑧
such that 𝑦+𝑧=𝑥+𝑘
on the blackboard.
Is it possible to make all the numbers on the blackboard equal? If yes, what is the minimum number of operations you
need?

### ideas

1. (sum + m * k) % (n + m) = 0
2. 假设最后的数字是x
3. sum + m * k = (n + m) * x
4. 所以第一步里面的(x, y) 其中一个必然是最后的x， 另外一个数y要么 = x, 要么 y + k = 2 * x
5. y = 2 * x - k
6. x + 2 * x - k = a[i] + k
7. 这里x已经可以确定出来了

### solution

Consider the "shifted" problem, where each 𝑥
on the blackboard (at any moment) is replaced with 𝑥′=𝑥−𝑘
.

Now, the operation becomes "replace 𝑥
with 𝑦+𝑧
such that 𝑦+𝑧=𝑥+𝑘⟹(𝑦′+𝑘)+(𝑧′+𝑘)=(𝑥′+𝑘)+𝑘⟹𝑦′+𝑧′=𝑥′
". Therefore, in the shifted problem, 𝑘′=0
.

Now you can replace every 𝑎′𝑖:=𝑎𝑖−𝑘
with any number of values with sum 𝑎′𝑖
, and the answer is the amount of numbers on the blackboard at the end, minus 𝑛
.

If we want to make all the values equal to 𝑚′
, it must be a divisor of every 𝑎′𝑖
.

If all the 𝑎′𝑖
are positive, it is optimal to choose 𝑚′=gcd𝑎′𝑖
.
If all the 𝑎′𝑖
are zero, the answer is 0
.
If all the 𝑎′𝑖
are negative, it is optimal to choose 𝑚′=−gcd−𝑎′𝑖
.
Otherwise, the answer is −1
.
Alternative way to get this result:

You have to split each 𝑎𝑖
into 𝑝𝑖
pieces equal to 𝑚
, and their sum must be equal to 𝑎𝑖+𝑘(𝑝𝑖−1)=(𝑎𝑖−𝑘)+𝑘𝑝𝑖=𝑚𝑝𝑖
. Then, (𝑎𝑖−𝑘)=(𝑚−𝑘)𝑝𝑖
, so 𝑚′=𝑚−𝑘
must be a divisor of every 𝑎′𝑖=𝑎𝑖−𝑘
.
In both the positive and the negative case, you will only write positive elements (in the original setup), as wanted.

If all the 𝑎′𝑖
are positive, the numbers you will write on the shifted blackboard are positive, so they will be positive also in the
original blackboard.
If all the 𝑎′𝑖
are negative, the numbers you will write on the shifted blackboard are greater than the numbers you will erase, so they
will be greater than the numbers in input (and positive) in the original blackboard.
Complexity: 𝑂(𝑛+log(max𝑎𝑖))
