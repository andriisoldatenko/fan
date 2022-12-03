FILE = open("input.txt")


def main(file):
    result = 0
    for line in file:
        cl_line = line.strip()
        first, last = set(cl_line[:len(cl_line)//2]), set(cl_line[len(cl_line)//2:])
        for ch in first.intersection(last):
            if ord(ch) - 96 > 0:
                result += ord(ch) - 96
            else:
                result += ord(ch) - 38
    return result


print(main(FILE))
