import sys
from collections import OrderedDict

FILE = sys.stdin
# FILE = open('sample.in')
first = True

while True:
    line = FILE.readline()
    if line == '':
        break
    elif not first:
        print()
    first = False

    num = int(line)
    names = FILE.readline().strip().split()
    totals = OrderedDict(zip(names, [0 for _ in range(num)]))
    for x in range(num):
        who, total, num_friends, *friends = FILE.readline().strip().split()
        total = int(total)
        num_friends = int(num_friends)
        pay = total // num_friends if (total > 0 and num_friends > 0) else 0
        penny = total - (pay * num_friends)
        totals[who] = totals[who] - total + penny
        for friend in friends:
            totals[friend] += pay
        # print(who, total, num_friends, friends)
    for user, final in totals.items():
        print(user, final, sep=' ')