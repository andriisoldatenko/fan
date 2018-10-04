import sys


FILE = sys.stdin
# FILE = open('sample.in')

while True:
    b = int(FILE.readline().strip())
    p = int(FILE.readline().strip())
    m = int(FILE.readline().strip())
    t = FILE.readline()
    print(pow(b, p, m))
    if t == '':
        break
