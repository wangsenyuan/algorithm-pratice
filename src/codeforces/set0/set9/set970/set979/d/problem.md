Kuro is currently playing an educational game about numbers. The game focuses on the greatest common divisor (GCD), the XOR value, and the sum of two numbers. Kuro loves the game so much that he solves levels by levels day by day.

Sadly, he's going on a vacation for a day, and he isn't able to continue his solving streak on his own. As Katie is a reliable person, Kuro kindly asked her to come to his house on this day to play the game for him.

Initally, there is an empty array 𝑎
. The game consists of 𝑞
 tasks of two types. The first type asks Katie to add a number 𝑢𝑖
 to 𝑎
. The second type asks Katie to find a number 𝑣
 existing in 𝑎
 such that 𝑘𝑖∣𝐺𝐶𝐷(𝑥𝑖,𝑣)
, 𝑥𝑖+𝑣≤𝑠𝑖
, and 𝑥𝑖⊕𝑣
 is maximized, where ⊕
 denotes the bitwise XOR operation, 𝐺𝐶𝐷(𝑐,𝑑)
 denotes the greatest common divisor of integers 𝑐
 and 𝑑
, and 𝑦∣𝑥
 means 𝑥
 is divisible by 𝑦
, or report -1 if no such numbers are found.

Since you are a programmer, Katie needs you to automatically and accurately perform the tasks in the game to satisfy her dear friend Kuro. Let's help her!


### ideas
1. ki | gcd(xi, v) 那么 let g = gcd(ki, xi), 那么 g | gcd(xi, v) 
2. 也就是说xi和v的gcd是 g的除数
3. 搞反了，是ki是xi的除数