You are given an integer value 𝑥 and a string 𝑠 consisting of digits from 1 to 9 inclusive.

A substring of a string is a contiguous subsequence of that string.

Let 𝑓(𝑙,𝑟) be the sum of digits of a substring 𝑠[𝑙..𝑟].

Let's call substring 𝑠[𝑙1..𝑟1] 𝑥-prime if

𝑓(𝑙1,𝑟1)=𝑥;
there are no values 𝑙2,𝑟2 such that
𝑙1≤𝑙2≤𝑟2≤𝑟1;
𝑓(𝑙2,𝑟2)≠𝑥;
𝑥 is divisible by 𝑓(𝑙2,𝑟2).
You are allowed to erase some characters from the string. If you erase a character, the two resulting parts of the string are concatenated without changing their order.

What is the minimum number of characters you should erase from the string so that there are no 𝑥-prime substrings in it? If there are no 𝑥-prime substrings in the given string 𝑠, then print 0.

Input
The first line contains a string 𝑠 (1≤|𝑠|≤1000). 𝑠 contains only digits from 1 to 9 inclusive.

The second line contains an integer 𝑥 (1≤𝑥≤20).

Output
Print a single integer — the minimum number of characters you should erase from the string so that there are no 𝑥-prime substrings in it. If there are no 𝑥-prime substrings in the given string 𝑠, then print 0.