import sys
from pprint import pprint

FILE = sys.stdin
# FILE = open('sample.in')

while True:
    num_banks, num_debentures = map(int, FILE.readline().strip().split())
    if num_banks == 0 and num_debentures == 0:
        break
    monetary_reserve = list(map(int, FILE.readline().strip().split()))

    for j in range(num_debentures):
        d, c, v = map(int, FILE.readline().strip().split())
        monetary_reserve[d-1] = monetary_reserve[d-1] - v
        monetary_reserve[c-1] = monetary_reserve[c-1] + v
    print('S' if all([x >= 0 for x in monetary_reserve]) else 'N')
