Time Limit: 2 sec / Memory Limit: 1024 MB

Score: 
400 points

Problem Statement
Takahashi is hosting a contest with 
N players numbered 
1 to 
N. The players will compete for points. Currently, all players have zero points.

Takahashi's foreseeing ability lets him know how the players' scores will change. Specifically, for 
i=1,2,…,T, the score of player 
A 
i
​
  will increase by 
B 
i
​
  points at 
i seconds from now. There will be no other change in the scores.

Takahashi, who prefers diversity in scores, wants to know how many different score values will appear among the players' scores at each moment. For each 
i=1,2,…,T, find the number of different score values among the players' scores at 
i+0.5 seconds from now.

For example, if the players have 
10, 
20, 
30, and 
20 points at some moment, there are three different score values among the players' scores at that moment.

### ideas
1. 假设到目前为止的分数的集合是scores set， 那么就是从中减去一个，再增加一个