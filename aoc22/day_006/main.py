FILE = open("input.txt")

M = 14

def main(file):
    result = []
    for line in file:
        N = len(line)
        for i in range(0, N - M):
            chars = (line[i:i + M])
            chars_s = set(chars)
            if len(chars_s) == M:
                result.append(i + M)
                break

    return result


print(main(FILE))
