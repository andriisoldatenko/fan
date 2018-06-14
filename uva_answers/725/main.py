import pprint
import sys

FILE = sys.stdin
# FILE = open('sample.in')

for line in FILE:
    n = int(line.strip())
    if n == 0:
        break

    for fghij in range(1234, 98765//n):
        abcde = fghij * n
        used = 1 if (fghij < 10000) else 0
        tmp = abcde
        while tmp:
            used |= 1 << (tmp % 10)
            tmp //= 10
        tmp = fghij
        while tmp:
            used |= 1 << (tmp % 10)
            tmp //= 10
        if used == ((1<<10) - 1):
            print('{:05d} / {:05d} = {}'.format(abcde, fghij, n))
