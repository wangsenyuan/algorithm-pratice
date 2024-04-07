Jury has hidden a permutation p of integers from 0 to n - 1. You know only the length n. Remind that in permutation all
integers are distinct.

Let b be the inverse permutation for p, i.e. pbi = i for all i. The only thing you can do is to ask xor of elements pi
and bj, printing two indices i and j (not necessarily distinct). As a result of the query with indices i and j you'll
get the value , where denotes the xor operation. You can find the description of xor operation in notes.

Note that some permutations can remain indistinguishable from the hidden one, even if you make all possible n2 queries.
You have to compute the number of permutations indistinguishable from the hidden one, and print one of such
permutations, making no more than 2n queries.

The hidden permutation does not depend on your queries.

Input
The first line contains single integer n (1 ≤ n ≤ 5000) — the length of the hidden permutation. You should read this
integer first.

Output
When your program is ready to print the answer, print three lines.

In the first line print "!".

In the second line print single integer answers_cnt — the number of permutations indistinguishable from the hidden one,
including the hidden one.

In the third line print n integers p0, p1, ..., pn - 1 (0 ≤ pi<n, all pi should be distinct) — one of the permutations
indistinguishable from the hidden one.

Your program should terminate after printing the answer.

Interaction
To ask about xor of two elements, print a string "? i j", where i and j — are integers from 0 to n - 1 — the index of
the permutation element and the index of the inverse permutation element you want to know the xor-sum for. After that
print a line break and make flush operation.

After printing the query your program should read single integer — the value of .

For a permutation of length n your program should make no more than 2n queries about xor-sum. Note that printing answer
doesn't count as a query. Note that you can't ask more than 2n questions. If you ask more than 2n questions or at least
one incorrect question, your solution will get "Wrong answer".

If at some moment your program reads -1 as an answer, it should immediately exit (for example, by calling exit(0)). You
will get "Wrong answer" in this case, it means that you asked more than 2n questions, or asked an invalid question. If
you ignore this, you can get other verdicts since your program will continue to read from a closed stream.

Your solution will get "Idleness Limit Exceeded", if you don't print anything or forget to flush the output, including
for the final answer .

To flush you can use (just after printing line break):

fflush(stdout) in C++;
System.out.flush() in Java;
stdout.flush() in Python;
flush(output) in Pascal;
For other languages see the documentation.

### observation

1. cpu干烧了
2. 先举个列子 p = 0, 1, 2
3. 那么 b = 0, 1, 2
4. 如果 p = 0, 2, 1, 那么 b = 0, 1, 2
5. p[i] ^ b[j]
6. p[b[j]] = j
7. b[p[i]] = i
8. 固定 i = 0, 询问 (i, j) j >= 1
9. 这个能给出什么信息呢？
10. 如果 query(i, j) = 0 => p[i] = b[j] => j = p[i] 或者 i = b[j]
11. 这样子可以找出 p[0] = 多少， 或者 b[j] = 0
12. 那么就进而知道 b[1], b[2], ....
13. 那这样子，岂不是只用n次查询？