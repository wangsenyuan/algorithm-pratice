Recently personal training sessions have finished in the Berland State University Olympiad Programmer Training Centre. By the results of these training sessions teams are composed for the oncoming team contest season. Each team consists of three people. All the students of the Centre possess numbers from 1 to 3n, and all the teams possess numbers from 1 to n. The splitting of students into teams is performed in the following manner: while there are people who are not part of a team, a person with the best total score is chosen among them (the captain of a new team), this person chooses for himself two teammates from those who is left according to his list of priorities. The list of every person's priorities is represented as a permutation from the rest of 3n - 1 students who attend the centre, besides himself.

You are given the results of personal training sessions which are a permutation of numbers from 1 to 3n, where the i-th number is the number of student who has won the i-th place. No two students share a place. You are also given the arrangement of the already formed teams in the order in which they has been created. Your task is to determine the list of priorities for the student number k. If there are several priority lists, choose the lexicographically minimal one.

### ideas
1. 如果k是某个team的leader，那么可以判断她的优先级
2. 否则，答案 = 1.。。。3*n 除去自己
3. 如果k是leader，假设当前它选择的另外两个伙伴是u, v, 且 u < v
4. 那么在前面所有出现过的，且<v的，都可以排在v的前面，其他的排在v的后面