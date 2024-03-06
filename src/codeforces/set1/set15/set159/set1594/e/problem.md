Theofanis started playing the new online game called "Among them". However, he always plays with Cypriot players, and
they all have the same name: "Andreas" (the most common name in Cyprus).

In each game, Theofanis plays with 𝑛
other players. Since they all have the same name, they are numbered from 1
to 𝑛
.

The players write 𝑚
comments in the chat. A comment has the structure of "𝑖
𝑗
𝑐
" where 𝑖
and 𝑗
are two distinct integers and 𝑐
is a string (1≤𝑖,𝑗≤𝑛
; 𝑖≠𝑗
; 𝑐
is either imposter or crewmate). The comment means that player 𝑖
said that player 𝑗
has the role 𝑐
.

An imposter always lies, and a crewmate always tells the truth.

Help Theofanis find the maximum possible number of imposters among all the other Cypriot players, or determine that the
comments contradict each other (see the notes for further explanation).

Note that each player has exactly one role: either imposter or crewmate.

### ideas

1. 如果i say j impster, 那么i和j不可能同时时同一个角色
2. 如果i say j crewmate, 那么如果i时imposter, 那么j也必须时imposter, 否则i是crewmate, 那么j也是crewmat
3. 所以对于第二种情况，i和j是同一种角色
4. 对于第一种i和j必须有不同的角色
5. dfs