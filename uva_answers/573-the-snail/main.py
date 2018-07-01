import sys

FILE = sys.stdin
# FILE = open('sample.in')
first = True

while True:
    h, u, d, f = map(int, FILE.readline().strip().split())
