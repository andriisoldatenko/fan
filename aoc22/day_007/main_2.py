result = []


DISK = 70000000

class Dir:
    def __init__(self, name, parent=None, size=0):
        self.name = name
        self.parent = parent
        self.children = []
        self.size = size
        self.total_size = size

    def add_children(self, name, size):
        self.children.append(Dir(name, size=size, parent=self))

    def cd(self, path):
        if path == "..":
            return self.parent

        for p in self.children:
            if p.name == path:
                return p
            else:
                self.cd(p.name)
        return

    def find_size(self):
        total_size = 0
        for p in self.children:
            if p.size > 0:
                total_size += p.size
            else:
                total_size += p.find_size()
        self.total_size = total_size
        return total_size

    def find_folders(self):
        for p in self.children:
            if p.size == 0:
                if p.total_size >= 0:
                    result.append(p.total_size)
            if p.children is not None:
                p.find_folders()

    def __repr__(self):
        return f"{self.name}, parent: {self.parent}, size: {self.size}, total: {self.total_size}"


FILE = open("input.txt")


# d = {"/": []}


# def f(tree):

def find_sing(arr):
    count = 0
    for line in arr:
        if line.startswith("$"):
            return count
        count += 1
    return count


def main(file):
    # result = []
    pos = 0
    f = list(map(str.rstrip, file.readlines()))
    tree = Dir("/")
    while pos < len(f):
        if f[pos] == "$ cd /":
            cwd = tree
            pos = + 1
            continue
        elif f[pos].startswith("$ cd"):
            name = f[pos].split()[-1]
            cwd = cwd.cd(name)
        elif f[pos] == "$ ls":
            sign = find_sing(f[pos + 1:]) + 1
            if sign == 0:
                break
            for el in f[pos + 1:pos + sign]:
                a, b = el.split()
                if a == "dir":
                    cwd.add_children(
                        b, 0
                    )
                elif a.isdigit():
                    cwd.add_children(b, int(a))
            pos += (sign-1)
            continue
        pos += 1
    tree.find_size()
    tree.find_folders()
    for s in sorted(result):
        if (70000000 - tree.total_size + s) >= 30_000_000:
            print(s)
            break
    return


main(FILE)

