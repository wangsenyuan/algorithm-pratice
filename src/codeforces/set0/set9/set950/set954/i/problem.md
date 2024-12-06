Suppose you have two strings s and t, and their length is equal. You may perform the following operation any number of times: choose two different characters c1 and c2, and replace every occurence of c1 in both strings with c2. Let's denote the distance between strings s and t as the minimum number of operations required to make these strings equal. For example, if s is abcd and t is ddcb, the distance between them is 2 — we may replace every occurence of a with b, so s becomes bbcd, and then we may replace every occurence of b with d, so both strings become ddcd.

You are given two strings S and T. For every substring of S consisting of |T| characters you have to determine the distance between this substring and T.

### ideas
1. 6个字符，假设一个映射，比如a到b，c到d，之类的（也就是最后的改变结果）
2. 这样映射，共有多少种呢？a可以映射b（这种情况下，b不用改变成a）
3. 所以，大体上有 6 * 4 * 2种？
4. 然后在这种映射的情况下，可以将T先变成映射后的结果， 形成新的t
5. 用新的t去匹配S。
6. 如果这个新的t和变化后的s也能匹配到，那么用到的字符的数量，就是操作数