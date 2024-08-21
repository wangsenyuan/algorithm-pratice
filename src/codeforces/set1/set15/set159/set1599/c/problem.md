Little Johnny Bubbles enjoys spending hours in front of his computer playing video games. His favorite game is Bubble Strike, fast-paced bubble shooting online game for two players.

Each game is set in one of the N maps, each having different terrain configuration. First phase of each game decides on which map the game will be played. The game system randomly selects three maps and shows them to the players. Each player must pick one of those three maps to be discarded. The game system then randomly selects one of the maps that were not picked by any of the players and starts the game.

Johnny is deeply enthusiastic about the game and wants to spend some time studying maps, thus increasing chances to win games played on those maps. However, he also needs to do his homework, so he does not have time to study all the maps. That is why he asked himself the following question: "What is the minimum number of maps I have to study, so that the probability to play one of those maps is at least P
"?

Can you help Johnny find the answer for this question? You can assume Johnny's opponents do not know him, and they will randomly pick maps.

Input
The first line contains two integers N
 (3
 ≤
 N
 ≤
 103
) and P
 (0
 ≤
 P
 ≤
 1
) – total number of maps in the game and probability to play map Johnny has studied. P
 will have at most four digits after the decimal point.

 ### ideas
 1. 假设John查看了m个地图, 
 2. 丛中选择3个地图，分4种情况，分别是选中0、1、2、3个其中的地图
 3. C(n - m, 3) (0 个)
 4. C(n - m, 2) * C(m, 1) (1)
 5. C(n - m, 1) * C(m, 2) 
 6. C(n - m, 0) * C(m, 3) 
 7. 在选中0个情况下， John的期望胜率是0
 8.  在选中1个情况下，这个地图被保留下来的概率是多少？John肯定会选择去掉其中一个不是的，而对手有50%的概率保留下选中的这个
 9.  在选中2/3个情况下，John获胜的概率都是1
 10. 一共有多少种可能性呢？ C(n, 3) * P(3) 