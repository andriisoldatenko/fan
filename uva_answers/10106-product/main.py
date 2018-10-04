import sys


FILE = sys.stdin
# FILE = open('sample.in')

while True:
    line = FILE.readline().strip()
    if line == '':
        break
    x = int(line)
    y = int(FILE.readline().strip())
    print(x * y)
