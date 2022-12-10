import re

f = open("test.txt", "r")
input = f.readlines()

visibility_map = [[False for _ in range(len(input[0].strip()))] for _ in range(len(input))]
height_map = [[int(i) for i in s.strip()] for s in input]


for y in range(len(height_map)):
    row = height_map[y]
    max_height = -1
    for x in range(len(row)):
        if row[x] > max_height:
            visibility_map[y][x] = True
            max_height = row[x]

    max_height = -1
    for x in range(len(row)-1, -1, -1):
        if row[x] > max_height:
            visibility_map[y][x] = True
            max_height = row[x]

for x in range(len(height_map[0])):
    max_height = -1
    for y in range(len(height_map)):
        tree = height_map[y][x]
        if tree > max_height:
            visibility_map[y][x] = True
            max_height = tree

    max_height = -1
    for y in range(len(height_map)-1, -1, -1):
        tree = height_map[y][x]
        if tree > max_height:
            visibility_map[y][x] = True
            max_height = tree

total = 0
for y in visibility_map:
    for x in y:
        if x:
            total += 1

# print(height_map)
# print(visibility_map)
print(total)