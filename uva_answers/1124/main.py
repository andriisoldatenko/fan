import sys
# from os import getenv
#
# if bool(getenv('DEBUG')):
#     FILE = open('sample.in')
# else:

FILE = sys.stdin

for line in sys.stdin:
    print(line, end='')
