import sys
# from os import getenv
#
# if bool(getenv('DEBUG')):
#     FILE = open('sample.in')
# else:

FILE = sys.stdin


def process_matrix(matrix):
    result_matrix = [[0 for _ in range(len(matrix[0]))]
                     for _ in range(len(matrix))]
    for i in range(len(matrix)):
        for j in range(len(matrix[0])):
            if matrix[i][j] == '*':
                result_matrix[i][j] = '*'
                continue
            # x1 x2 x3
            # x4 XX x5
            # x6 x7 x8
            x11, y11 = (i - 1, j - 1)
            x22, y22 = (i - 1, j)
            x33, y33 = (i - 1, j + 1)
            x44, y44 = (i, j - 1)
            x55, y55 = (i, j + 1)
            x66, y66 = (i + 1, j - 1)
            x77, y77 = (i + 1, j)
            x88, y88 = (i + 1, j + 1)
            if min(x11, y11) >= 0 and matrix[x11][y11] == '*':
                result_matrix[i][j] += 1
            if min(x22, y22) >= 0 and matrix[x22][y22] == '*':
                result_matrix[i][j] += 1
            if x33 >= 0 and y33 <= len(matrix[0]) - 1 and matrix[x33][y33] == '*':
                result_matrix[i][j] += 1
            if 0 <= y44 <= len(matrix[0]) - 1 and matrix[x44][y44] == '*':
                result_matrix[i][j] += 1
            if y55 <= len(matrix[0]) - 1 and matrix[x55][y55] == '*':
                result_matrix[i][j] += 1
            if 0 <= x66 <= len(matrix) - 1 and 0 <= y66 <= len(matrix[0]) - 1 and matrix[x66][y66] == '*':
                result_matrix[i][j] += 1
            if x77 <= len(matrix) - 1 and y77 <= len(matrix[0]) - 1 and matrix[x77][y77] == '*':
                result_matrix[i][j] += 1
            if x88 <= len(matrix) - 1 and y88 <= len(matrix[0]) - 1 and matrix[x88][y88] == '*':
                result_matrix[i][j] += 1
    return result_matrix

field_index = 1
result = ''
for line in sys.stdin:
    # print(line)
    if line.strip() == '0 0':
        print(result[:len(result) - 2])
        break
    if line.strip() == '':
        continue
    a, b = map(int, line.split())
    matrix = [[0 for _ in range(b)] for _ in range(a)]
    for i in range(a):
        l = list(FILE.readline().strip())
        for j in range(b):
            matrix[i][j] = l[j]
    # pprint(matrix)
    m = process_matrix(matrix)
    result += 'Field #{}:\n'.format(field_index )
    res = ''
    for i in m:
       res = res + ''.join(map(str, i)) + "\n"
    res += '\n'
    field_index += 1
    result += res