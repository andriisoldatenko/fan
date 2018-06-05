import pprint
import sys
import re

FILE = sys.stdin
#FILE = open('sample.in')

lines = FILE.readlines()
lines = [l.strip() for l in lines]
max_length = len(max(lines, key=len))
for c in range(max_length):
    for s in reversed(lines):
        try:
            print(s[c], end='')
        except IndexError:
            print(' ', end='')
    print()

