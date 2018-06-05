import pprint
import sys
import re
import itertools

FILE = sys.stdin
#FILE = open('sample.in')

for line in FILE:
    j= list(map(int, line.strip().split()))
    jolly = j[1:]
    l = []
    for i,j in zip(jolly[1:], jolly[:-1]):
        l.append(abs(i - j))
    l.sort()
    if sorted(l) == list(range(1, len(l)+1)):
        print('Jolly')
    else:
        print('Not jolly')

