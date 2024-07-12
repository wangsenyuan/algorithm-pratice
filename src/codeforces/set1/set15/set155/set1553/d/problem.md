You are given two strings 𝑠
 and 𝑡
, both consisting of lowercase English letters. You are going to type the string 𝑠
 character by character, from the first character to the last one.

When typing a character, instead of pressing the button corresponding to it, you can press the "Backspace" button. It deletes the last character you have typed among those that aren't deleted yet (or does nothing if there are no characters in the current string). For example, if 𝑠
 is "abcbd" and you press Backspace instead of typing the first and the fourth characters, you will get the string "bd" (the first press of Backspace deletes no character, and the second press deletes the character 'c'). Another example, if 𝑠
 is "abcaa" and you press Backspace instead of the last two letters, then the resulting text is "a".

Your task is to determine whether you can obtain the string 𝑡
, if you type the string 𝑠
 and press "Backspace" instead of typing several (maybe zero) characters of 𝑠
.

### ideas.
1. 如果 s[i] != t[j], backspace, 这个时候，是会把前面的某个字符也给删掉
2. 所以，从后面往前
3. 如果 s[i] == t[j], continue
4. else 找到最近的k, s[k] = t[j], 那么 i - k 必须是偶数