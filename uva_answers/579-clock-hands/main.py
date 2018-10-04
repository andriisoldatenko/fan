import sys

# FILE = sys.stdin
FILE = open('sample.in')


hour_angels = 360 / 12 # 30
minute_angels = 360 / 60 # 6
hour_min_angels = 360 / (12 * 60)

while True:
    line = FILE.readline().strip()
    if line == '0:00':
        break
    hours, mins = map(int, line.split(':'))
    diff_mins = min(abs(0 - mins), abs(60 - mins))
    diff_hours = min(abs(0 - hours), abs(12 - hours))
    # print(diff_hours * hour_angels)
    # print(hour_min_angels * mins)
    # print(diff_mins * minute_angels)
    t = min(abs(diff_hours*hour_angels + diff_mins*minute_angels), abs(diff_hours*hour_angels - diff_mins*minute_angels))
    print("%.3f" % abs(hour_min_angels*mins + t))
