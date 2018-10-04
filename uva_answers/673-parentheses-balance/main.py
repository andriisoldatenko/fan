import sys

FILE = sys.stdin
test_cases = int(input())


def check_parentheses(parentheses):
    p = parentheses
    while p:
        p0, p = p, p.replace('()', '').replace('[]', '')
        if p == p0:
            return 'No'
    return 'Yes'

for test_case in range(test_cases):
    line = FILE.readline().strip()
    print(check_parentheses(line))


def check_parentheses(parentheses):
    opening = tuple('({[')
    closing = tuple(')}]')
    mapping = dict(zip(opening, closing))
    queue = []

    for letter in parentheses:
        if letter in opening:
            queue.append(mapping[letter])
        elif letter in closing:
            if not queue or letter != queue.pop():
                return False
    return not queue
