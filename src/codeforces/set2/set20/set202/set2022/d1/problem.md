This is an interactive problem.

It is a tradition in Mexico's national IOI trainings to play the game "Asesino", which is similar to "Among Us" or "Mafia".

Today, 𝑛
 players, numbered from 1
 to 𝑛
, will play "Asesino" with the following three roles:

Knight: a Knight is someone who always tells the truth.
Knave: a Knave is someone who always lies.
Impostor: an Impostor is someone everybody thinks is a Knight, but is secretly a Knave.
Each player will be assigned a role in the game. There will be exactly one Impostor but there can be any (possible zero) number of Knights and Knaves.

As the game moderator, you have accidentally forgotten the roles of everyone, but you need to determine the player who is the Impostor.

To determine the Impostor, you will ask some questions. In each question, you will pick two players 𝑖
 and 𝑗
 (1≤𝑖,𝑗≤𝑛
; 𝑖≠𝑗
) and ask if player 𝑖
 thinks that player 𝑗
 is a Knight. The results of the question is shown in the table below.

Knight	Knave	Impostor
Knight	Yes	No	Yes
Knave	No	Yes	No
Impostor	No	Yes	—
The response of the cell in row 𝑎
 and column 𝑏
 is the result of asking a question when 𝑖
 has role 𝑎
 and 𝑗
 has row 𝑏
. For example, the "Yes" in the top right cell belongs to row "Knight" and column "Impostor", so it is the response when 𝑖
 is a Knight and 𝑗
 is an Impostor.
Find the Impostor in at most 𝑛+69
 questions.

Note: the grader is adaptive: the roles of the players are not fixed in the beginning and may change depending on your questions. However, it is guaranteed that there exists an assignment of roles that is consistent with all previously asked questions under the constraints of this problem.


### ideas
1. 如果不考虑Imposter, 如果 ? i j 返回Yes， 那么i, j是同一组的，就是它们要么都是Knight,要么都是Knave
2. ? i j 返回False => 它们肯定是不同组的
3. 现在考虑Imposter, 如果 ? i j, i是Imposter， 那么它和Kave是同一组，如果j时Imposter，那么i如果是Knight，就会以为j是同一组的
4. 假设1不是Imposter，那么 ? 1 i, 就可以把人分成两组， 和1一组的，和1不是一组的
5. 如果1是Knight，那么Imposter肯定在第一组里面, 似乎没法保证在额外的69次询问后得到答案
6. 如果没有把Imposter当做i来问，似乎没有特征哪
7. 考虑3个人，a, b, c
8. 根据 f(a, b), f(a, c), f(b, c)的结果能判断出什么呢？
9. 如果它们中存在Imposter, 结果有可能是不一致的
10. 假设不存在的情况下， (a, b) / c  那么 (a, b) = true, (b, c) = false, (c, a) = false
11. 如果 (a, b, c) => (a, b) = true, (b, c) = true, (c, a) = true
12. 假设其中a是Imposter, (b, c) = false, 那么 b, c肯定是两组
13. 那么 (a, b) = true (b是knave), (c, a) = true (c是knight)
14.     (a, b) = false (b是Knight), (c, a) = false (c是Knave)
15. 如果a是Knight， (b, c) = false
16.     (a, b) = true, (b是Knight) (c, a) = false (c是knave)
17.     (a, b) = false, (b 是Knave), (c, a) = true (c是Knight)
18. 如果a是Knave (b, c) = false
19.     (a, b) = true, (b时Knave) (c, a) = false (c时Knight)
20.     (a, b) = false, (b时Knight) （c, a) = true (C 是Knave)
21.  如果一组里面没有Imposter， 那么这一组循环问下来，有可能全部为true，但不可能全部为false
22.  这是因为3个人中，如果职业全部相同，那么就是全为true，但是至少有两个人职业相同，那么就是一个true， 两个false
23.   如果一组里面Imposter，如果另外两个人职业不同，那么会出现全部为false的情况 a是I, b是T，c是N
24.     或者两个true，一个false a是I，b是N，C是T
25.     或者两个true，一个false  a是I，b是T，c是T => 
26.     或者两个true，一个false  a是I，b是N，c是N
27.  也就是说3个一组，直到遇到上面的情况（两个true，或者3个false）
28.  如果是3个false再问额外 ask(b, a), ask(c, b)，ask(c, a)
29.  如果是两个true 
  