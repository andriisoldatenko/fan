import sys

FILE = sys.stdin
# FILE = open('sample.in')
# inpt = FILE.readline().strip()

for line in FILE:
    if line.strip() == '-1 -1':
        break
    a, b = map(int, line.strip().split())
    print(min(max(a, b) - min(a, b), 100 - max(a, b) + min(a, b)))
