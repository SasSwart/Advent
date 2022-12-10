import re

f = open("test.txt", "r")
input = f.readlines()

height_map = [[int(i) for i in s.strip()] for s in input]
view_value_map = [[1 for _ in s.strip()] for s in input]

for y in range(len(height_map)):
	row = height_map[y]
	for x in range(len(row)):
		home_run = True
		for j in range(x-1, -1, -1):
			if row[j] >= row[x]:
				view_value_map[y][x] = view_value_map[y][x] * (x-j)
				home_run = False
				break
		if home_run:
			view_value_map[y][x] = view_value_map[y][x] * x

	for x in range(len(row)-1,-1,-1):
		home_run = True
		for j in range(x+1, len(row)):
			if row[j] >= row[x]:
				view_value_map[y][x] = view_value_map[y][x] * (j-x)
				home_run = False
				break
		if home_run:
			view_value_map[y][x] = view_value_map[y][x] * (len(row)-1-x)

for x in range(len(height_map[0])):
	col = [row[x] for row in height_map]
	for y in range(len(col)):
		home_run = True
		for j in range(y-1, -1, -1):
			if col[j] >= col[y]:
				view_value_map[y][x] = view_value_map[y][x] * (y-j)
				home_run = False
				break
		if home_run:
			view_value_map[y][x] = view_value_map[y][x] * y

	for y in range(len(col)-1,-1,-1):
		home_run = True
		for j in range(y+1, len(col)):
			if col[j] >= col[y]:
				view_value_map[y][x] = view_value_map[y][x] * (j-y)
				home_run = False
				break
		if home_run:
			view_value_map[y][x] = view_value_map[y][x] * (len(col)-1-y)


print(max([max(row) for row in view_value_map]))