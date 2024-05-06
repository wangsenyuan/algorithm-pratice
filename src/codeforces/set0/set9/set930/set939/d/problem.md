Valya and Tolya are an ideal pair, but they quarrel sometimes. Recently, Valya took offense at her boyfriend because he
came to her in t-shirt with lettering that differs from lettering on her pullover. Now she doesn't want to see him and
Tolya is seating at his room and crying at her photos all day long.

This story could be very sad but fairy godmother (Tolya's grandmother) decided to help them and restore their
relationship. She secretly took Tolya's t-shirt and Valya's pullover and wants to make the letterings on them same. In
order to do this, for one unit of mana she can buy a spell that can change some letters on the clothes. Your task is
calculate the minimum amount of mana that Tolya's grandmother should spend to rescue love of Tolya and Valya.

More formally, letterings on Tolya's t-shirt and Valya's pullover are two strings with same length n consisting only of
lowercase English letters. Using one unit of mana, grandmother can buy a spell of form (c1, c2) (where c1 and c2 are
some lowercase English letters), which can arbitrary number of times transform a single letter c1 to c2 and vise-versa
on both Tolya's t-shirt and Valya's pullover. You should find the minimum amount of mana that grandmother should spend
to buy a set of spells that can make the letterings equal. In addition you should output the required set of spells.

### ideas

1. c1 -> c2, c2 -> c3, 那么 (c1 -> c3)
2. c1 -> c2 -> c3 -> c4... -> c1会怎样？
3. 这个不大对，如果c1和c2有了， c2和 c3能连通了，那么没有必要使用c1到c3
4. 但是如果 (c1 -> c3) 这个转换特别多，那么优先购买 c1->c3也是比较好的选择