import re

f = open("test.txt", "r")
board, instructions = map(lambda buf : buf.split("\n"), f.read().split("\n\n"))

board_height = len(board)-1
num_stacks = (len(board[-1])+1)//4
stacks = []


for i in range(num_stacks):
    stack = []
    for line in range(board_height-1, -1, -1):
        crate = board[line][4*i+1]
        if crate != " ":
            stack.append(crate)
    stacks.append(stack)

for instruction in instructions:
    num_crates, from_col, to_col = map(int, re.findall("\d+", instruction))
    for i in range(num_crates):
        stacks[to_col-1].append(stacks[from_col-1].pop())

print([s[-1] for s in stacks])