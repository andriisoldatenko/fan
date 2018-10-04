import sys

LETTERS = list('`1234567890-=QWERTYUIOP[]\\ASDFGHJKL;\'ZXCVBNM,./')

FILE = sys.stdin
# FILE = open('sample.in')

while True:
    line = FILE.readline()
    if line == '':
        break
    chars = list(line)
    for ch in chars:
        if ch in LETTERS:
            print(LETTERS[LETTERS.index(ch) - 1], end='')
        else:
            print(ch, end='')
