One little-known hacker wanted to get the administrator password for the AtForces testing system in order to find out the tasks for the upcoming contest. To achieve this, he got into the administrator's office and stole a piece of paper with a list of𝑛
passwords - strings consisting of lowercase Latin letters.

The hacker returned home and began preparing for the hack. He discovered that the system contained only the passwords from the stolen piece of paper, and that the system determined the equivalence of the passwords𝑎
And𝑏
as follows:

two passwords𝑎
And𝑏
are equivalent if there is a symbol that is in both𝑎
, and in𝑏
;
two passwords𝑎
And𝑏
equivalent if another password exists𝑐
from the list, which are equivalent to both𝑎
, And𝑏
.
If a certain password is set in the system and an equivalent one is used, the user gains access to the system.

For example, if the list contains passwords " a ", " b ", " ab ", " d ", then from the system's point of view, passwords " a ", " b ", " ab " are equivalent to each other, and password " d " is not equivalent to any other. In other words, if:

the set password is, for example, " b ", then you can log into the system using any of the three passwords: " a ", " b ", " ab ";
the set password is " d ", then you can only log in to the system using the password " d ".
It is known that exactly one password from the list is the administrator password for the testing system. Help the hacker understand what minimum number of passwords will need to be used during hacking to be guaranteed to gain access to the system. Keep in mind that the hacker does not know what password is set in the system.
