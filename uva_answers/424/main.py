import pprint
import sys
import re

FILE = sys.stdin
#FILE = open('sample.in')

result = 0
for line in FILE:
    n = int(line.strip())
    result += n
    if n == 0:
        print(result)
        break
