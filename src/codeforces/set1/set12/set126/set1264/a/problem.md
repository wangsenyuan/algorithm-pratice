So the Beautiful Regional Contest (BeRC) has come to an end! 𝑛
 students took part in the contest. The final standings are already known: the participant in the 𝑖
-th place solved 𝑝𝑖
 problems. Since the participants are primarily sorted by the number of solved problems, then 𝑝1≥𝑝2≥⋯≥𝑝𝑛
.

Help the jury distribute the gold, silver and bronze medals. Let their numbers be 𝑔
, 𝑠
 and 𝑏
, respectively. Here is a list of requirements from the rules, which all must be satisfied:

for each of the three types of medals, at least one medal must be awarded (that is, 𝑔>0
, 𝑠>0
 and 𝑏>0
);
the number of gold medals must be strictly less than the number of silver and the number of bronze (that is, 𝑔<𝑠
 and 𝑔<𝑏
, but there are no requirements between 𝑠
 and 𝑏
);
each gold medalist must solve strictly more problems than any awarded with a silver medal;
each silver medalist must solve strictly more problems than any awarded a bronze medal;
each bronze medalist must solve strictly more problems than any participant not awarded a medal;
the total number of medalists 𝑔+𝑠+𝑏
 should not exceed half of all participants (for example, if 𝑛=21
, then you can award a maximum of 10
 participants, and if 𝑛=26
, then you can award a maximum of 13
 participants).
The jury wants to reward with medals the total maximal number participants (i.e. to maximize 𝑔+𝑠+𝑏
) so that all of the items listed above are fulfilled. Help the jury find such a way to award medals.

### ideas
1. enum g, try to maximize s & g