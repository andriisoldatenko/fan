import sys

FILE = sys.stdin

test_case = 1

for line in FILE:
    l = line.strip()
    if l == 'Hajj':
        print('Case {}: {}'.format(test_case, 'Hajj-e-Akbar'))
    elif l == '*':
        break
    else:
        print('Case {}: {}'.format(test_case, 'Hajj-e-Asghar'))
    test_case += 1
