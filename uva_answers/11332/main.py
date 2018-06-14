import pprint
import sys
import re

FILE = sys.stdin
#FILE = open('sample.in')

def sum_numbers(n):
    if n < 10:
        return n
    else:
        return sum_numbers(sum(map(int, list(str(n)))))


for line in FILE:
    n = int(line.strip())
    if n == 0:
        break
    print(sum_numbers(n))

