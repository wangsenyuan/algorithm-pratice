Max wants to buy a new skateboard. He has calculated the amount of money that is needed to buy a new skateboard. He left a calculator on the floor and went to ask some money from his parents. Meanwhile his little brother Yusuf came and started to press the keys randomly. Unfortunately Max has forgotten the number which he had calculated. The only thing he knows is that the number is divisible by 4.

You are given a string s consisting of digits (the number on the display of the calculator after Yusuf randomly pressed the keys). Your task is to find the number of substrings which are divisible by 4. A substring can start with a zero.

A substring of a string is a nonempty sequence of consecutive characters.

For example if string s is 124 then we have four substrings that are divisible by 4: 12, 4, 24 and 124. For the string 04 the answer is three: 0, 4, 04.

As input/output can reach huge size it is recommended to use fast input/output methods: for example, prefer to use gets/scanf/printf instead of getline/cin/cout in C++, prefer to use BufferedReader/PrintWriter instead of Scanner/System.out in Java.

### ideas
1. l...r 可以整除4
2. pref[r] = pref[r-1] * 10 % 4 + num[r] % 4
3. (pref[r] - pref[r]) % 4 = 0