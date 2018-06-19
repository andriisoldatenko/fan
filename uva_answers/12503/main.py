import sys

FILE = sys.stdin
# FILE = open('sample.in')
test_case = 1
total_test_cases = range(int(input()))

for tc in total_test_cases:
    num_instructions = range(int(input()))
    instructions = []
    for instruction in num_instructions:
        current_instruction = FILE.readline().strip()
        if current_instruction == 'LEFT':
            instructions.append(-1)
        elif current_instruction == 'RIGHT':
            instructions.append(1)
        elif current_instruction.startswith('SAME AS'):
            same_as = int(current_instruction.split('SAME AS')[-1])
            instructions.append(instructions[same_as-1])
    print(sum(instructions))
