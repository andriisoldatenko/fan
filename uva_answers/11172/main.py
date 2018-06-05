import pprint
import sys
import re

FILE = sys.stdin
#FILE = open('sample.in')

test_cases = int(FILE.readline().strip())

for t in range(test_cases):
    a, b = map(int, FILE.readline().strip().split())
    if a == b:
        print('=')
    elif a > b:
        print('>')
    else:
        print('<')
