This is an interactive problem.

There is a secret sequence 𝑝0,𝑝1,…,𝑝𝑛−1
, which is a permutation of {0,1,…,𝑛−1}
.

You need to find any two indices 𝑖
 and 𝑗
 such that 𝑝𝑖⊕𝑝𝑗
 is maximized, where ⊕
 denotes the bitwise XOR operation.

To do this, you can ask queries. Each query has the following form: you pick arbitrary indices 𝑎
, 𝑏
, 𝑐
, and 𝑑
 (0≤𝑎,𝑏,𝑐,𝑑<𝑛
). Next, the jury calculates 𝑥=(𝑝𝑎∣𝑝𝑏)
 and 𝑦=(𝑝𝑐∣𝑝𝑑)
, where |
 denotes the bitwise OR operation. Finally, you receive the result of comparison between 𝑥
 and 𝑦
. In other words, you are told if 𝑥<𝑦
, 𝑥>𝑦
, or 𝑥=𝑦
.

Please find any two indices 𝑖
 and 𝑗
 (0≤𝑖,𝑗<𝑛
) such that 𝑝𝑖⊕𝑝𝑗
 is maximum among all such pairs, using at most 3𝑛
 queries. If there are multiple pairs of indices satisfying the condition, you may output any one of them.

Input
Each test contains multiple test cases. The first line contains the number of test cases 𝑡
 (1≤𝑡≤103
). The description of the test cases follows.

Interaction
The first line of each test case contains one integer 𝑛
 (2≤𝑛≤104
). At this moment, the permutation 𝑝0,𝑝1,…,𝑝𝑛−1
 is chosen. The interactor in this task is not adaptive. In other words, the sequence 𝑝
 is fixed in every test case and does not change during the interaction.

To ask a query, you need to pick four indices 𝑎
, 𝑏
, 𝑐
, and 𝑑
 (0≤𝑎,𝑏,𝑐,𝑑<𝑛
) and print the line of the following form:

"? a b c d"
After that, you receive:

"<" if (𝑝𝑎∣𝑝𝑏)<(𝑝𝑐∣𝑝𝑑)
;
"=" if (𝑝𝑎∣𝑝𝑏)=(𝑝𝑐∣𝑝𝑑)
;
">" if (𝑝𝑎∣𝑝𝑏)>(𝑝𝑐∣𝑝𝑑)
.
You can make at most 3𝑛
 queries of this form.

Next, if your program has found a pair of indices 𝑖
 and 𝑗
 (0≤𝑖,𝑗<𝑛
) such that 𝑝𝑖⊕𝑝𝑗
 is maximized, print the line of the following form:

"! i j"
Note that this line is not considered a query and is not taken into account when counting the number of queries asked.

After this, proceed to the next test case.

If you make more than 3𝑛
 queries during an interaction, your program must terminate immediately, and you will receive the Wrong Answer verdict. Otherwise, you can get an arbitrary verdict because your solution will continue to read from a closed stream.

After printing a query or the answer for a test case, do not forget to output the end of line and flush the output. Otherwise, you will get the verdict Idleness Limit Exceeded. To do this, use:

fflush(stdout) or cout.flush() in C++;
System.out.flush() in Java;
flush(output) in Pascal;
stdout.flush() in Python;
see the documentation for other languages.
It is guaranteed that the sum of 𝑛
 over all test cases does not exceed 104
.

Hacks

To hack, follow the test format below.

The first line contains the number of test cases 𝑡
 (1≤𝑡≤103
). The description of the test cases follows.

The first line of each test case contains one integer 𝑛
 (2≤𝑛≤104
).

The second line of each test case contains 𝑛
 integers 𝑝0,𝑝1,…,𝑝𝑛−1
, which represent a permutation of integers from 0
 to 𝑛−1
.

The sum of 𝑛
 over all test cases should not exceed 104
.

### ideas
1. n = 2, [0, 1] => 1
2. n = 3, [0, 1, 2] => 1 ^ 2 = 3
3. n = 4, [0, 1, 2, 3] => 1 ^ 2 = 3
4. n = 5, [0, 1, 2, 3, 4] => 4 ^ 3 = 7
5. n = 6, [0, 1, 2, 3, 4, 5] => 7
6. n = 7, ... 7
7. n = 8, [0, 1, 2, 3, 4, 5, 6, 7] = 7
8. n = 9, [0,.....              7, 8] = 15
9. (1 << h) 〉=n , result = （1 << h） - 1
10. 如果n = 1 << h, 比如上面的8，可选范围比较大 
11. 如果n < 1 << h, 比如上面的9，就必须选择 1 << h, 和 (1 <<h) - 1
12. 无论怎么样，似乎找到 3, 7, 15 这样的数，会非常有利
13. 假设n = 4,  且 p = [2, 1, 0, 3]
14. [2 | 1] == [0 | 3]
15. 如果以两个为一组，且每次保留最大的那一组，那么在n/2次后，得到一组最大的
16. 好像不需要a,b,c,d不同
17. 这里首先如果能找出最大的数字n-1, 然后，再想办法找出 mask ^ (n - 1)？
18. 