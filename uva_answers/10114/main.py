import sys
# from decimal import Decimal as D
# float = D

FILE = sys.stdin
# FILE = open('sample.in')

while True:
    duration, *other = FILE.readline().strip().split()
    duration = int(duration)
    if duration < 0:
        break
    down_payment, amount_loan, counter = other
    down_payment = float(down_payment)
    amount_loan = float(amount_loan)
    counter = int(counter)
    # print(duration, other)
    dep_percentages = [0 for x in range(max(duration, counter)+1)]
    for _ in range(counter):
        m_num, dep_percentage = FILE.readline().strip().split()
        dep_percentages[int(m_num)] = float(dep_percentage)

    all_deps = []
    first_el = dep_percentages[0]
    for index, y in enumerate(dep_percentages):
        if y != 0:
            all_deps.append(y)
            first_el = y
        else:
            all_deps.append(first_el)
    total_months = 0
    car_initial_value = amount_loan + down_payment
    worth_amount = car_initial_value - (car_initial_value * all_deps[0])
    borrower_owes = amount_loan
    # import ipdb; ipdb.set_trace()
    if borrower_owes < worth_amount:
        print('{} months'.format(total_months))
        continue

    for m_index in range(1, duration+1):
        borrower_owes -= (amount_loan / duration)
        worth_amount = worth_amount - (worth_amount * all_deps[m_index])
        total_months += 1
        if borrower_owes < worth_amount:
            if total_months != 1:
                print('{} months'.format(total_months))
                # print(borrower_owes, worth_amount)
            else:
                print('{} month'.format(total_months))
            break
