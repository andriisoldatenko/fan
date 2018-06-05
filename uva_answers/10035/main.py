import pprint
import sys
import re

FILE = sys.stdin
#FILE = open('sample.in')

for line in FILE:
#    import ipdb; ipdb.set_trace()
    if line.strip() == '0 0':
        break
    a, b = line.strip().split()
    num1 = list(reversed(list(map(int, list(a)))))
    num2 = list(reversed(list(map(int, list(b)))))
    short = max(num1, num2, key=len)
    total_carry = 0
    carry = 0
    for i in range(len(short)):
        try:
            n = num1[i]
        except IndexError:
            n = 0
        try:
            m = num2[i]
        except IndexError:
            m = 0
         
        if carry + n + m >= 10:
            carry = 1
            total_carry += 1
        else:
            carry = 0
    print('{} carry operation{}.'.format(
        'No' if total_carry == 0 else total_carry, 
        '' if total_carry <= 1 else 's' ))

