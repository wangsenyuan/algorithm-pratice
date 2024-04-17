Let's call some positive integer classy if its decimal representation contains no more than 3
non-zero digits. For example, numbers 4
, 200000
, 10203
are classy and numbers 4231
, 102306
, 7277420000
are not.

You are given a segment [𝐿;𝑅]
. Count the number of classy integers 𝑥
such that 𝐿≤𝑥≤𝑅
.

Each testcase contains several segments, for each of them you are required to solve the problem separately.

Input
The first line contains a single integer 𝑇
(1≤𝑇≤104
) — the number of segments in a testcase.

Each of the next 𝑇
lines contains two integers 𝐿𝑖
and 𝑅𝑖
(1≤𝐿𝑖≤𝑅𝑖≤1018
).

Output
Print 𝑇
lines — the 𝑖
-th line should contain the number of classy integers on a segment [𝐿𝑖;𝑅𝑖]
.

### ideas

1. 如果数字 < 1000 => classy
2. 如果数字的长度=length > 3,
3. 那么第一位必然不是0， 所以在剩下的length-1中，选择3位，使他们不等于0，那么这个数字就是 non-classy, 其他的数字随便选
4. C(length - 1, 3) * pow(9, 3) * pow(10, length - 4) * 9 (第一个数字)
5. 