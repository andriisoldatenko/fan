import pprint
import sys

FILE = sys.stdin
# FILE = open('sample.in')


for line in FILE:
    v, t = map(int, line.strip().split())
    print(2*t*v)
