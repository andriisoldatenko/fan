import sys
# from os import getenv
#
# if bool(getenv('DEBUG')):
#     FILE = open('sample.in')
# else:

FILE = sys.stdin

for line in sys.stdin:
    cl_line = line.strip()
    start, p1, p2, p3 = map(int, cl_line.split())
    if line.strip() == '0 0 0 0':
        break

    # turn the dial clockwise 2 full turns
    # stop at the first number of the combination
    # turn the dial counter-clockwise 1 full turn
    # continue turning counter-clockwise until the 2nd number is reached
    # turn the dial clockwise again until the 3rd number is reached
    # pull the shank and the lock will open.
    # Example:
    # 0 30 0 30 -> 1350

    r1 = 3 * 360
    a1 = 40 - abs(start - p1) if start < p1 else abs(start - p1)
    a2 = 40 - abs(p2 - p1) if p1 > p2 else abs(p1 - p2)
    a3  = 40 - abs(p2 - p3) if p2 < p3 else abs(p3 - p2)
    print(r1 + (a1 + a2 + a3)*9)
    # print('-' * 20)
