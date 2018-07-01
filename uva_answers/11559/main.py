import sys

FILE = sys.stdin
# FILE = open('sample.in')

MAX_NUMBER = 1000000000

while True:
    line = FILE.readline()
    if line == '':
        break
    num_participants, budget, num_hotels, num_weeks = map(
        int,line.strip().split())
    hotels = []
    for _ in range(num_hotels):
        hotel_price = int(FILE.readline().strip())
        total_price = num_participants * hotel_price
        ava_beds = list(map(int, FILE.readline().strip().split()))
        available_beds = any([b >= num_participants for b in ava_beds])
        # print(available_beds)
        # print(total_price)
        if (total_price <= budget and available_beds
                and len(ava_beds) <= num_weeks):
            hotels.append(total_price)
        else:
            hotels.append(MAX_NUMBER)
    print(min(hotels) if min(hotels) != MAX_NUMBER else 'stay home')
