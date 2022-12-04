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


print(main(FILE))
