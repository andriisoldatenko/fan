FILE = open("input.txt")


def main(file):
    result = 0
    for line in file:
        cl_line = line.strip()
        left, right = cl_line.split(",")
        l1, l2 = map(int, left.split("-"))
        r1, r2 = map(int, right.split("-"))
        print(l1, l2, r1, r2)
        if l1 <= r1 <= l2 and l1 <= r2 <= l2:
            result += 1
        elif r1 <= l1 <= r2 and r1 <= l2 <= r2:
            result += 1
    return result

def main2(file):
    reconsider = []
    for line in file:
        x1, x2, y1, y2 = map(int, [y for x in line.split(",") for y in x.split("-")])
        setx = set(range(x1, x2 + 1))
        sety = set(range(y1, y2 + 1))
        if setx.issubset(sety) or sety.issubset(setx):
            reconsider.append(line)
    return len(reconsider)

print(main2(FILE))
print(main(FILE))
