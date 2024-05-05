Petya has n positive integers a1, a2, ..., an.

His friend Vasya decided to joke and replaced all digits in Petya's numbers with a letters. He used the lowercase
letters of the Latin alphabet from 'a' to 'j' and replaced all digits 0 with one letter, all digits 1 with another
letter and so on. For any two different digits Vasya used distinct letters from 'a' to 'j'.

Your task is to restore Petya's numbers. The restored numbers should be positive integers without leading zeros. Since
there can be multiple ways to do it, determine the minimum possible sum of all Petya's numbers after the restoration. It
is guaranteed that before Vasya's joke all Petya's numbers did not have leading zeros.

### ideas

1. 在首位的不能是0
2. 然后为了得到最小的sum，应该让最高位的尽量的小