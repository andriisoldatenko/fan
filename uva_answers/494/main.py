import pprint
import sys
import re

FILE = sys.stdin
# FILE = open('sample.in')

for line in FILE:
    print(len(re.findall(r'[a-zA-Z]+', line)))
