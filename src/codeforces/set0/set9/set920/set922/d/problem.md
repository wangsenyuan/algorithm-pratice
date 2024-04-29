Pushok the dog has been chasing Imp for a few hours already.

Fortunately, Imp knows that Pushok is afraid of a robot vacuum cleaner.

While moving, the robot generates a string t consisting of letters 's' and 'h', that produces a lot of noise. We define
noise of string t as the number of occurrences of string "sh" as a subsequence in it, in other words, the number of such
pairs (i, j), that i<j and and .

The robot is off at the moment. Imp knows that it has a sequence of strings ti in its memory, and he can arbitrary
change their order. When the robot is started, it generates the string t as a concatenation of these strings in the
given order. The noise of the resulting string equals the noise of this concatenation.

Help Imp to find the maximum noise he can achieve by changing the order of the strings.

### ideas

1. 在一个t[i]中，是固定的，
2. 对于t[i]，t[j] 他们的顺序要按照产生的sh来判断
3. 假设 t[i] = (a, b)
4. t[j] = (c, d)
5. 第一个是s的数量，第二个是h的数量
6. 如果t[i]+t[j]， 那么额外产生的sh = a * d,
7. 如果t[j]+t[i] 那么产生的是 c * b
8. 看哪个更大，就排在前面