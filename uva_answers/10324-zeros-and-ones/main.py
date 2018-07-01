import sys

FILE = sys.stdin
# FILE = open('sample.in')
test_case = 1

while True:
    num = FILE.readline().strip()
    if num == '':
        break
    positions = int(FILE.readline().strip())
    results = []
    print('Case {}:'.format(test_case))
    a = [0 for _ in range(len(num))]
    a[0] = 1 if num[0] == '1' else 0
    k = 1
    while k < len(num):
        a[k] = a[k - 1] + 1 if num[k] == '1' else a[k - 1]
        k += 1
    for _ in range(positions):
        ii, jj = map(int, FILE.readline().strip().split())
        i = min(ii, jj)
        j = max(ii, jj)
        if num[i] == '0' and num[j] == '0' and a[j] - a[i] == 0:
            print('Yes')
        elif num[i] == '1' and num[j] == '1' and a[j] - a[i] == j - i:
            print('Yes')
        else:
            print('No')
    test_case += 1
