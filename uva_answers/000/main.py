import itertools


def main():
    for perm in sorted(itertools.permutations('aba')):
        print(''.join(perm))
    # aab
    # aab
    # aba
    # aba
    # baa
    # baa

if __name__ == '__main__':
    main()
