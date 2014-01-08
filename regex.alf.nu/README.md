# [http://regex.alf.nu/](http://regex.alf.nu/)

1. Plain strings (207) -> foo

	> "foo" is simply the smallest match.

2. Anchors (208) -> k$

	> All Good end with "ick" but there is no Bad ending with "k" so "k$" alone is sufficient.

3. Ranges (202) -> ^[a-f]+$

	> All Bad have always some letters not in range a-f.

4. Backrefs (201) -> (...).*\1

	> All Good have a three letters reoccurrence.

5. Abba (193) -> ^(?!.*(.)(.)\2\1)

	> All Bad have one palindrome in them.

6. A man, a plan (176) -> ^(.)(.).*\2\1$

	> All Good are palindromes. At least two letters are needed because of the Bad "sporous".

7. Prime (286) -> ^(?!(..+)\1+$)

	> Prime factorization with just one prime. For example the Bad with 9 letters can be seen as 3 3x blocks.

8. Four (199) -> ([aeiou])(.\1){3}

	> All Good have 4 blocks of X<some letter>. (I cannot figure out what the tip means)

9. Order (198) -> ^[^o].{4,5}$

	> All Good have 5-6 letters. There is one Bad "oriole" so "o" has to be excluded from the start.

10. Triples (456) -> ^([0369]*([258][0369]*[147][0369]*)*)*(([147][0369]*|[258][0369]*[258][0369]*)([147][0369]*[258][0369]*)*([258][0369]*|[147][0369]*[147][0369]*)([258][0369]*[147][0369]*)*)*$

	> Freaking FSM took forever. There is a lot that can be improved but I'm happy now.

11. Glob (342) -> ^(?:(.+) .+ \1|\*([^*]+) .+ .+\2|([^*]+)\* .+ \3.+|.*\*([^*]+)\*.* .+ .+\4.+)$

	> Implemented globs: exact, left, right and left-right

12. Balance (273) -> ^(<(<(<(<(<(<(<>)*>)*>)*>)*>)*>)*>)+$

	> This is pretty straight forward but I guess there are better solutions...

13. Powers (72) -> ^(?:(x{64})+|x{32}|x{16}|x{8}|x{4}|xx?)$

	> Every Good > 64chars is divideable by 64. The rest must be matched exactly. I gues that there is a MUCH better solution to this.

14. Long count (253) -> ^((.+)0 \2+1 ?)+$

	> The Good have repeating two blocks with one ending with 0 the other with 1. A "+" after \2 is needed because \21 would be another reference.

15. Long count v2 (253) -> ^((.+)0 \2+1 ?)+$

	> This is the same as v1... why?!

16. Alphabetical -> Cannot figure it out.... yet

**SUM = 3519**
