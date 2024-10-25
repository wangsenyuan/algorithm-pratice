Fox Ciel is going to publish a paper on LIS (Fox-Operated Intelligent Systems). She heard that the list of authors of a scientific paper is always sorted in lexicographic order.

After studying the examples of the design, the fox discovered that sometimes this statement is not true. In some articles, the authors' names are not sorted in lexicographical order in the usual sense. But it turns out that after some change in the order of the letters in the alphabet, the order of the authors becomes lexicographical !

The fox wants to know if there is such an order of the letters of the Latin alphabet that the names of the authors of the article she is proposing follow in lexicographic order. If so, then she also needs to find any such order.

The lexicographical order is determined as follows. We compare s and t by first finding the leftmost position with different characters: s i  ≠  t i . If there is no such position (that is, s is a prefix of t or vice versa), then the shorter string is smaller. Otherwise, we compare the characters s i and t i according to their order in the alphabet.

### ideas
1. n <= 100
2. 对于每个字符对,a, b 需要判断它们的关系，如果 a < b 建立一条a -> b的边， 如果没有限制，就不建立边
3. 如果在最终的途中存在环 => -1
4. 否则的话，用前序访问的顺序就可以了
5. 