import pprint
import sys

FILE = sys.stdin
# FILE = open('sample.in')


for line in FILE:
    h, o = map(int, line.strip().split())
    print(abs(o - h))
