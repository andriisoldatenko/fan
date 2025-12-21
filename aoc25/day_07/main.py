with open("input.txt", mode="r", encoding="utf-8") as f:
    data = f.read().strip().splitlines()[::2]
print(data)

processed_rows = [[1 if char == "S" else 0 for char in data[0]]]
print(f"processed_rows={processed_rows}")
splits = 0
for row in data[1:]:
    current = [-1 if char == "^" else 0 for char in row]
    for i, char in enumerate(row[1:-1], 1):
        above = processed_rows[-1][i]
        if above > 0:  # -1 signifies split, 0 not possible
            if char == "^":
                splits += 1
                current[i - 1] += above
                current[i + 1] += above
            else:
                current[i] += above
    processed_rows.append(current)


timelines = sum(i for i in processed_rows[-1] if i > 0)


print(f"Part 1: {splits}\nPart 2: {timelines}")