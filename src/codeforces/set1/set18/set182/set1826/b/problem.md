* https://codeforces.com/problemset/problem/1826/B
* description
  You have an array 𝑎
  of 𝑛
  non-negative integers. Let's define 𝑓(𝑎,𝑥)=[𝑎1mod𝑥,𝑎2mod𝑥,…,𝑎𝑛mod𝑥]
  for some positive integer 𝑥
  . Find the biggest 𝑥
  , such that 𝑓(𝑎,𝑥)
  is a palindrome.

  Here, 𝑎mod𝑥
  is the remainder of the integer division of 𝑎
  by 𝑥
  .
  
  An array is a palindrome if it reads the same backward as forward. More formally, an array 𝑎
  of length 𝑛
  is a palindrome if for every 𝑖
  (1≤𝑖≤𝑛
  ) 𝑎𝑖=𝑎𝑛−𝑖+1
  .
  
  Input
  The first line contains a single integer 𝑡
  (1≤𝑡≤105
  ) — the number of test cases.
  
  The first line of each test case contains a single integer 𝑛
  (1≤𝑛≤105
  ).
  
  The second line of each test case contains 𝑛
  integers 𝑎𝑖
  (0≤𝑎𝑖≤109
  ).
  
  It's guaranteed that the sum of all 𝑛
  does not exceed 105
  .
  
  Output
  For each test case output the biggest 𝑥
  , such that 𝑓(𝑎,𝑥)
  is a palindrome. If 𝑥
  can be infinitely large, output 0
  instead.