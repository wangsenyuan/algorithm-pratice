The "Bulls and Cows" game needs two people to play. The thinker thinks of a number and the guesser tries to guess it.

The thinker thinks of a four-digit number in the decimal system. All the digits in the number are different and the number may have a leading zero. It can't have more than one leading zero, because all it's digits should be different. The guesser tries to guess the number. He makes a series of guesses, trying experimental numbers and receives answers from the first person in the format "x bulls y cows". x represents the number of digits in the experimental number that occupy the same positions as in the sought number. y represents the number of digits of the experimental number that present in the sought number, but occupy different positions. Naturally, the experimental numbers, as well as the sought number, are represented by four-digit numbers where all digits are different and a leading zero can be present.

For example, let's suppose that the thinker thought of the number 0123. Then the guessers' experimental number 1263 will receive a reply "1 bull 2 cows" (3 occupies the same positions in both numbers and 1 and 2 are present in both numbers but they occupy different positions). Also, the answer to number 8103 will be "2 bulls 1 cow" (analogically, 1 and 3 occupy the same positions and 0 occupies a different one).

When the guesser is answered "4 bulls 0 cows", the game is over.

Now the guesser has already made several guesses and wants to know whether his next guess can possibly be the last one.

Input
The first input line contains an integer n (1 ≤ n ≤ 10) which represents the number of already made guesses. Then follow n lines in the form of "ai bi ci", where ai is the i-th experimental number, bi is the number of bulls, ci is the number of cows (1 ≤ i ≤ n, 0 ≤ bi, ci, bi + ci ≤ 4). The experimental numbers are correct, i.e., each of them contains exactly four digits, in each of them all the four digits are different, and there can be a leading zero. All the experimental numbers are different. As the guesser hasn't guessed the number yet, the answer "4 bulls 0 cows" is not present.

Output
If the input data is enough to determine the sought number, print the number with four digits on a single line. If it has less than four digits, add leading zero. If the data is not enough, print "Need more data" without the quotes. If the thinker happens to have made a mistake in his replies, print "Incorrect data" without the quotes.


### ideas
1. 我们就从10000个数去匹配就可以了