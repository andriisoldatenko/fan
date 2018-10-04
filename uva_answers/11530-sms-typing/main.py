import sys
import string

LETTERS = list('abcdefghijklmnopqrtuvwxy ')
EXCEPTION_LETTERS = list('sz')

test_case = 1
FILE = sys.stdin
test_cases = int(input())

for test_case in range(1, test_cases+1):
    chars = FILE.readline()
    total_typed = 0
    for char in chars:
        if char in EXCEPTION_LETTERS:
            total_typed += 4
        elif char == '\n':
            pass
        else:
            if ((LETTERS.index(char)+1) % 3) == 0:
                total_typed += 3
            else:
                total_typed += ((LETTERS.index(char)+1) % 3)
    print('Case #{}: {}'.format(test_case, total_typed))
