import sys
import string
FILE = sys.stdin
# FILE = open('sample.in')

counter = 1

d = {   
    'HELLO': 'ENGLISH',
    'HOLA': 'SPANISH',
    'HALLO': 'GERMAN',
    'BONJOUR': 'FRENCH',
    'CIAO': 'ITALIAN',
    'ZDRAVSTVUJTE': 'RUSSIAN'
}


for line in FILE:
    if line.strip() == '#':
        break
    print('Case {}: {}'.format(counter, d.get(line.strip(), 'UNKNOWN')))
    counter +=1

map(int, FILE.readline().strip().split())
