import re

FILE = open("input.txt")

N = 9

STACK = {
    1: [],
    2: [],
    3: [],
    4: [],
    5: [],
    6: [],
    7: [],
    8: [],
    9: [],
}


def main(file):
    for line in file:
        cl_line = line
        if line.startswith(" 1 2") or line == "\n":
            continue

        if line.startswith("move"):
            numbers = list(map(int, re.findall(r"\d+", line)))
            sub = []
            for j in range(numbers[0]):
                item = STACK[numbers[1]].pop()
                sub.insert(0, item)
            try:
                STACK[numbers[2]].extend(sub)
            except (IndexError, KeyError):
                pass
            continue

        chars = cl_line.split(" ")
        k = 0
        r = []
        for ch in chars:
            if k % 4 == 0 or ch.startswith("["):
                r.append(ch)
            k += 1

        for i in range(1, N + 1):
            if i - 1 < len(r) and r[i - 1] != '':
                STACK[i].insert(0, r[i - 1])
    result = []
    for i in range(N):
        result.append(STACK[i + 1][-1])
    return "".join(re.findall(r"\w", "".join(result)))


print(main(FILE))

