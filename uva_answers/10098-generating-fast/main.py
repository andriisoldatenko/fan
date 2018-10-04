import itertools
import sys

test_case = 1
FILE = sys.stdin
test_cases = int(FILE.readline())

for tc in range(1, test_cases+1):
    anagram = FILE.readline().strip()
    sorted_anagram = sorted(anagram)
    an = set()
    for a in itertools.permutations(sorted_anagram):
        an.add(''.join(a))
    print('\n'.join(sorted(list(an))))
    print()
