import sys
# from os import getenv
#
# if bool(getenv('DEBUG')):
#     FILE = open('sample.in')
# else:

FILE = sys.stdin

results = []
i = True
for c in sys.stdin.read():
    if c == '"' and i:
        results.append('``')
        i = not i
    elif c == '"' and not i:
        results.append("''")
        i = not i
    else:
        results.append(c)
print(''.join(results), end='')
