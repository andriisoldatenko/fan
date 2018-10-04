import sys
import calendar

ordinary_msg = 'This is an ordinary year.'

leap_msg = 'This is leap year.'
huluculu_msg = 'This is huluculu festival year.'
bulukulu_msg = 'This is bulukulu festival year.'

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
    y = int(line)
    msg = []
    if calendar.isleap(y):
        msg.append(leap_msg)
        if y % 55 == 0:
            msg.append(bulukulu_msg)
    if y % 15 == 0:
        msg.append(huluculu_msg)

    if not msg:
        print(ordinary_msg)
    else:
        print('\n'.join(sorted(msg, reverse=True)))
