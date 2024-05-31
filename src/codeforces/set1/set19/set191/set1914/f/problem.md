BerSoft is the biggest IT corporation in Berland. There are 𝑛
 employees at BerSoft company, numbered from 1
 to 𝑛
.

The first employee is the head of the company, and he does not have any superiors. Every other employee 𝑖
 has exactly one direct superior 𝑝𝑖
.

Employee 𝑥
 is considered to be a superior (direct or indirect) of employee 𝑦
 if one of the following conditions holds:

employee 𝑥
 is the direct superior of employee 𝑦
;
employee 𝑥
 is a superior of the direct superior of employee 𝑦
.
The structure of BerSoft is organized in such a way that the head of the company is superior of every employee.

A programming competition is going to be held soon. Two-person teams should be created for this purpose. However, if one employee in a team is the superior of another, they are uncomfortable together. So, teams of two people should be created so that no one is the superior of the other. Note that no employee can participate in more than one team.

Your task is to calculate the maximum possible number of teams according to the aforementioned rules.


### ideas
1. 计算的不是有多少种获得2人team的数量，而是最多可以有多少个team
2. 似乎顺序也有关系
3. 也不是简单的先匹配子树，再匹配上层