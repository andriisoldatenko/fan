import pprint
import sys
import re

FILE = sys.stdin
#FILE = open('sample.in')

def is_palindrome(n):
    k = str(n)
    return list(k) == list(reversed(k))

def reverse_add(n):
    return n + int(str(n)[::-1])

#import ipdb;ipdb.set_trace() 
test_cases = range(int(FILE.readline()))
#import ipdb; ipdb.set_trace()
for tc in test_cases:
    n = int(FILE.readline().strip())
    total_sum_count = 1
    n = reverse_add(n)
    while not is_palindrome(n):
        n = reverse_add(n)
        total_sum_count += 1
    print(total_sum_count, n)
