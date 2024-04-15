Chouti and his classmates are going to the university soon. To say goodbye to each other, the class has planned a big
farewell party in which classmates, teachers and parents sang and danced.

Chouti remembered that 𝑛
persons took part in that party. To make the party funnier, each person wore one hat among 𝑛
kinds of weird hats numbered 1,2,…𝑛
. It is possible that several persons wore hats of the same kind. Some kinds of hats can remain unclaimed by anyone.

After the party, the 𝑖
-th person said that there were 𝑎𝑖
persons wearing a hat differing from his own.

It has been some days, so Chouti forgot all about others' hats, but he is curious about that. Let 𝑏𝑖
be the number of hat type the 𝑖
-th person was wearing, Chouti wants you to find any possible 𝑏1,𝑏2,…,𝑏𝑛
that doesn't contradict with any person's statement. Because some persons might have a poor memory, there could be no
solution at all.

Input
The first line contains a single integer 𝑛
(1≤𝑛≤105
), the number of persons in the party.

The second line contains 𝑛
integers 𝑎1,𝑎2,…,𝑎𝑛
(0≤𝑎𝑖≤𝑛−1
), the statements of people.

Output
If there is no solution, print a single line "Impossible".

Otherwise, print "Possible" and then 𝑛
integers 𝑏1,𝑏2,…,𝑏𝑛
(1≤𝑏𝑖≤𝑛
).

If there are multiple answers, print any of them.

### ideas

1. a[i] 表示，其他a[i]个人的帽子和自己的不同，也就是表示 n - 1 - a[i]个人的帽子和自己相同
2. 且相同帽子的人的a[i]是相同的
3. 但是相同a[i]的，不一定相同帽子