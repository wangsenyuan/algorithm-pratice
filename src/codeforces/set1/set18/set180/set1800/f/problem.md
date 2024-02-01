Dasha, an excellent student, is studying at the best mathematical lyceum in the country. Recently, a mysterious stranger
brought 𝑛
words consisting of small latin letters 𝑠1,𝑠2,…,𝑠𝑛
to the lyceum. Since that day, Dasha has been tormented by nightmares.

Consider some pair of integers ⟨𝑖,𝑗⟩(1≤𝑖≤𝑗≤𝑛). A nightmare is a string for which it is true:

    It is obtained by concatenation 𝑠𝑖𝑠𝑗 ;
    Its length is odd;
    The number of different letters in it is exactly 25 ;
    The number of occurrences of each letter that is in the word is odd.

For example, if 𝑠𝑖="abcdefg" and 𝑠𝑗="ijklmnopqrstuvwxyz", the pair ⟨𝑖,𝑗⟩ creates a nightmare.

Dasha will stop having nightmares if she counts their number. There are too many nightmares, so Dasha needs your help.
Count the number of different nightmares.

Nightmares are called different if the corresponding pairs ⟨𝑖,𝑗⟩
are different. The pairs ⟨𝑖1,𝑗1⟩
and ⟨𝑖2,𝑗2⟩
are called different if 𝑖1≠𝑖2
or 𝑗1≠𝑗2
.

### thoughts

1. 如果不考虑