import sys


FILE = sys.stdin
# FILE = open('sample.in')

for line in FILE:
    p = int(input())
    if line == '':
        break
    n = int(line)
    print(round(p ** (1/float(n))))
