import sys
from itertools import cycle

FILE = sys.stdin
# FILE = open('sample.in')
# inpt = FILE.readline().strip()

test_case = 1
test_cases = range(1, int(input())+1)

words = "Happy birthday to you Happy birthday to you Happy birthday to " \
        "Rujia Happy birthday to you".split()

people = []
for tc in test_cases:
    people.append(input())

estimated_length = 0
while estimated_length < len(people):
    estimated_length += len(words)

cycled_people = cycle(people)
cycled_words = cycle(words)

for _ in range(estimated_length):
    print('{}: {}'.format(next(cycled_people), next(cycled_words)))
