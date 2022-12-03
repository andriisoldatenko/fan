FILE = open("input2.txt")


def main(file):
    result = 0
    buf = []
    line_num = 1
    for line in file:
        cl_line = line.strip()
        buf.append(set(cl_line))
        if line_num % 3 == 0:
            for ch in set.intersection(*buf):
                if ord(ch) - 96 > 0:
                    result += ord(ch) - 96
                else:
                    result += ord(ch) - 38
            buf = []
        line_num += 1
    return result

print(main(FILE))
