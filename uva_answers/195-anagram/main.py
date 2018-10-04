import string
import sys

test_case = 1
FILE = sys.stdin
test_cases = int(FILE.readline())

LETTERS_ORDER = ''
for l, u in zip(string.ascii_uppercase, string.ascii_lowercase):
    LETTERS_ORDER += (l + u)


def next_permutation(a):
    i = len(a) - 2
    while not (i < 0 or LETTERS_ORDER.index(a[i]) < LETTERS_ORDER.index(a[i+1])):
        i -= 1
    if i < 0:
        return False
    j = len(a) - 1
    while not (LETTERS_ORDER.index(a[j]) > LETTERS_ORDER.index(a[i])):
        j -= 1
    a[i], a[j] = a[j], a[i]
    a[i+1:] = reversed(a[i+1:])
    return True


for _ in range(1, test_cases+1):
    anagram = FILE.readline().strip()
    sorted_anagram = sorted(anagram, key=lambda x: LETTERS_ORDER.index(x))
    if len(sorted_anagram) == 1:
        print(''.join(sorted_anagram))
        continue
    if len(sorted_anagram) == 2:
        print(''.join(sorted_anagram))
        print(''.join(reversed(sorted_anagram)))
        continue

    while True:
        print(''.join(sorted_anagram))
        tt =  next_permutation(sorted_anagram)
        if not tt:
            break
