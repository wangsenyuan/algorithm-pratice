In the Catowice city next weekend the cat contest will be held. However, the jury members and the contestants haven't been selected yet. There are n
 residents and n
 cats in the Catowice, and each resident has exactly one cat living in his house. The residents and cats are numbered with integers from 1
 to n
, where the i
-th cat is living in the house of i
-th resident.

Each Catowice resident is in friendship with several cats, including the one living in his house. In order to conduct a contest, at least one jury member is needed and at least one cat contestant is needed. Of course, every jury member should know none of the contestants. For the contest to be successful, it's also needed that the number of jury members plus the number of contestants is equal to n
.

Please help Catowice residents to select the jury and the contestants for the upcoming competition, or determine that it's impossible to do.

### ideas
1. 显然每个house中，要么resident参与，要么cat参与（否则没法保证 j + p = n)
2. 所以有 a => !b[a]
3. 问题是怎么避免出现0个jury的情况？