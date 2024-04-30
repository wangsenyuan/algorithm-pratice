Jamie is preparing a Codeforces round. He has got an idea for a problem, but does not know how to solve it. Help him
write a solution to the following problem:

Find k integers such that the sum of two to the power of each number equals to the number n and the largest integer in
the answer is as small as possible. As there may be multiple answers, you are asked to output the lexicographically
largest one.

To be more clear, consider all integer sequence with length k (a1, a2, ..., ak) with . Give a value to each sequence.
Among all sequence(s) that have the minimum y value, output the one that is the lexicographically largest.

For definitions of powers and lexicographical order see notes.

### ideas

1. 不考虑k的限制， n可以被分解为二进制表示
2. 假设从高到第 a1, a2, a3.... 表示n = sum(2 ** a1, 2 ** a2....)
3. 如果这个数量 = k, 那么只能是这个序列吗？
4. 假设 > k => no 因为貌似只能增加数量
5. < k, 可以将一个 a1 => 2个a2