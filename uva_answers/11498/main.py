import sys

FILE = sys.stdin
# FILE = open('sample.in')

while True:
    k = int(FILE.readline().strip())
    if k == 0:
        break
    n, m = map(int, FILE.readline().strip().split())
    for _ in range(k):
        x, y = map(int, FILE.readline().strip().split())
        if x == n or y == m:
            print('divisa')
        elif x > n and y > m:
            print('NE')
        elif x < n and y < m:
            print('SO')
        elif x < n and y > m:
            print('NO')
        elif x > n and m > y:
            print('SE')
