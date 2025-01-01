Melody Pond was stolen from her parents as a newborn baby by Madame Kovarian, to become a weapon of the Silence in their crusade against the Doctor. Madame Kovarian changed Melody's name to River Song, giving her a new identity that allowed her to kill the Eleventh Doctor.

Heidi figured out that Madame Kovarian uses a very complicated hashing function in order to change the names of the babies she steals. In order to prevent this from happening to future Doctors, Heidi decided to prepare herself by learning some basic hashing techniques.

The first hashing function she designed is as follows.

Given two positive integers (𝑥,𝑦)
 she defines 𝐻(𝑥,𝑦):=𝑥2+2𝑥𝑦+𝑥+1
.

Now, Heidi wonders if the function is reversible. That is, given a positive integer 𝑟
, can you find a pair (𝑥,𝑦)
 (of positive integers) such that 𝐻(𝑥,𝑦)=𝑟
?

If multiple such pairs exist, output the one with smallest possible 𝑥
. If there is no such pair, output "NO".

### ideas
1. brute force on x