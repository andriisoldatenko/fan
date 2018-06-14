import pprint
import sys

FILE = sys.stdin
#FILE = open('sample.in')
def f91(n):
    if n <= 100:
        return 91
    elif n>= 101:
        return n - 10


for line in FILE:
    if line.strip() == '0':
        break
    n = int(line.strip())
    print('f91({}) = {}'.format(n, f91(n)))
