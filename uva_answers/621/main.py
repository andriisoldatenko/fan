for _ in range(int(input())):
    n = input()
    if n in {'1', '4', '78'}:
        print('+')
    elif n.endswith('35'):
        print('-')
    elif n.startswith('9') and n.endswith('4'):
        print('*')
    else:
        print('?')
