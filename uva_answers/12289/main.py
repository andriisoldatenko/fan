import sys
FILE = sys.stdin
# FILE = open('sample.in')

k = FILE.readline()

numbers = {1:'one',
           2: 'two',
           3: 'three',
           4: 'four',
           5: 'five',
           6: 'six',
           7: 'seven',
           8: 'eight',
           9: 'nine',
           10: 'ten'}

for line in FILE:
    word = line.strip()
    for k, v in numbers.items():
        if len(word) != len(v):
            continue
        count = 0
        for c1, c2 in zip(word, v):
            if c1 != c2:
                count +=1
        if count == 1 or count == 0:
            print(k)
            break
