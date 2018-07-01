import sys
from pprint import pprint

FILE = sys.stdin
# FILE = open('sample.in')

test_cases = range(int(FILE.readline()))
FILE.readline()

MAX_SIZE = 6

for tc in test_cases:
    w_num = int(FILE.readline())
    w = range(w_num)
    d = 0
    for row in w:
        y1, y2 = map(int, FILE.readline().strip().split())
        d += (y1 -y2)
    if d % w_num == 0:
        print('yes')
    else:
        print('no')
    l = FILE.readline()
    if l:
        print()
