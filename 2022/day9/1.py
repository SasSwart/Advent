import itertools
import math
import numpy as np
from os import system
import time

def print_grid(rope):
    size = 16
    for y in range(size, -size, -1):
        print(y,'\t', end='')
        for x in range(-size, size,):
            if rope[0] == [x, y]:
                print("H", end='')
            elif [x, y] == [0,0]:
                print("s", end='')
            elif [x, y] in rope:
                print(str(rope.index([x,y])), end='')
            else:
                print(".", end='')
        print()
    print()

f = open("test.txt", "r")
input = f.readlines()

rope = [[0, 0] for _ in range(10)]
tail_positions = [rope[1][:]]

for instruction in input:
    direction, step = instruction.strip().split(" ")
    for i in range(int(step)):
        match direction:
            case "L":
                rope[0][0] -= 1
            case "R":
                rope[0][0] += 1
            case "U":
                rope[0][1] += 1
            case "D":
                rope[0][1] -= 1
        for head, tail in itertools.pairwise(rope):
            deltax = head[0] - tail[0]
            deltay = head[1] - tail[1]
            distance = math.sqrt(deltax**2 + deltay**2)
            if distance <= 1:
                continue

            if deltax == 0 and abs(deltay) == 2:
                tail[1] += np.sign(deltay)
            elif deltay == 0 and abs(deltax) == 2:
                tail[0] += np.sign(deltax)
            elif abs(deltax) == 1 and abs(deltay) == 2:
                tail[1] += np.sign(deltay)
                tail[0] += np.sign(deltax)
            elif abs(deltay) == 1 and abs(deltax) == 2:
                tail[0] += np.sign(deltax)
                tail[1] += np.sign(deltay)
            elif abs(deltay) == 2 and abs(deltax) == 2:
                tail[0] += np.sign(deltax)
                tail[1] += np.sign(deltay)
            tail_positions.append(rope[-1][:])
        system('clear')
        print(instruction)
        print_grid(rope)
        time.sleep(0.02)

print(len(set([str(tail) for tail in tail_positions])))