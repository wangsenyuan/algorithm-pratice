IA has so many colorful magnets on her fridge! Exactly one letter is written on each magnet, 'a' or 'b'. She loves to
play with them, placing all magnets in a row. However, the girl is quickly bored and usually thinks how to make her
entertainment more interesting.

Today, when IA looked at the fridge, she noticed that the word formed by magnets is really messy. "It would look much
better when I'll swap some of them!" — thought the girl — "but how to do it?". After a while, she got an idea. IA will
look at all prefixes with lengths from 1
to the length of the word and for each prefix she will either reverse this prefix or leave it as it is. She will
consider the prefixes in the fixed order: from the shortest to the largest. She wants to get the lexicographically
smallest possible word after she considers all prefixes. Can you help her, telling which prefixes should be chosen for
reversing?

A string 𝑎
is lexicographically smaller than a string 𝑏
if and only if one of the following holds:

𝑎
is a prefix of 𝑏
, but 𝑎≠𝑏
;
in the first position where 𝑎
and 𝑏
differ, the string 𝑎
has a letter that appears earlier in the alphabet than the corresponding letter in 𝑏
.

### ideas

1. 有个感觉最优结果是把所有的a排在b的前面
2. 假设在前缀i之间的处理，可以保证所有的a要么在前面，要么在后半部分
3. 如果当前遇到一个新的a，那么仍然可以保证这一点
4. 如果前面的a就在后半部分，已经满足了条件（需要转、或者不转）可以变成前半部分
5. 如果前面的a在前部分，那么必然是以b结尾的，而且必然是没有swap的，那么此时转一下即可