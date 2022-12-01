import numpy as np

f = open("test.txt", "r")
input = f.read()
elves = input.split("\n\n")
elves = [[int(snack) for snack in elf.split("\n")] for elf in elves]
elves = [np.sum(elf) for elf in elves]
elves.sort()
elves = elves[-3:]

print(np.sum(elves))