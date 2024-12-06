Polycarp came up with a new programming language. There are only two types of statements in it:

"x := s": assign the variable named x the value s (where s is a string). For example, the statement var := hello assigns the variable named var the value hello. Note that s is the value of a string, not the name of a variable. Between the variable name, the := operator and the string contains exactly one space each.
"x = a + b": assign the variable named x the concatenation of values of two variables a and b. For example, if the program consists of three statements a := hello, b := world, c = a + b, then the variable c will contain the string helloworld. It is guaranteed that the program is correct and the variables a and b were previously defined. There is exactly one space between the variable names and the = and + operators.
All variable names and strings only consist of lowercase letters of the English alphabet and do not exceed 5
 characters.

The result of the program is the number of occurrences of string haha in the string that was written to the variable in the last statement.

Polycarp was very tired while inventing that language. He asks you to implement it. Your task is — for given program statements calculate the number of occurrences of string haha in the last assigned variable.

### ideas
1. n虽然比较小 （50），但是最终的字符串的长度，还是非常可观的
2. f(s) 表示字符串s中的ha的数量
3. s = a + b => f(s) = f(a) + f(b) + 1 (如果a以h结尾，且b以a开始)
4. 好办了