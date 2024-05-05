In a dream Marco met an elderly man with a pair of black glasses. The man told him the key to immortality and then
disappeared with the wind of time.

When he woke up, he only remembered that the key was a sequence of positive integers of some length n, but forgot the
exact sequence. Let the elements of the sequence be a1, a2, ..., an. He remembered that he calculated gcd(ai, ai +
1, ..., aj) for every 1 ≤ i ≤ j ≤ n and put it into a set S. gcd here means the greatest common divisor.

Note that even if a number is put into the set S twice or more, it only appears once in the set.

Now Marco gives you the set S and asks you to help him figure out the initial sequence. If there are many solutions,
print any of them. It is also possible that there are no sequences that produce the set S, in this case print -1.

### ideas

1. 可以先排序
2. 最小的数，肯定是[1...n]的gcd
3. 然后第二小的数，它要们是1..n-1, 或者 2....n的数，但是它们必须能整除s[0]
4. 然后我们可以假设吗？s[1] 是 [1...n]，那么有 gcd(s[1], a[n-1]) = s[0]
5. 那么不妨设a[0] = s[0]
6. 然后[0...n-1] = s[2]， a[n-1] = s[0]
7. 然后s[3]....s[m] 就有点麻烦了，这是因为，要将它们分组，有一部分是s[1]一组，有一部分是s[2]一组
8. 等等，某个s[i]肯定是原有数组的a[j]
9. a[i] = s[m]
10. 或者，我们其实不需要计算a，我们只要验证s是满足条件就可以了
11. 如果 gcd(s[i], s[i+1]) 存在
12. graph 