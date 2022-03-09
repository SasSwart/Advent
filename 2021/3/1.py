raw_data = open("power_cons.txt", "r").read()

ordered_data = raw_data.split()

lineLength = len(ordered_data[0])
bitFrequency = [0 for _ in ordered_data[0]]
gamma = 0
epsilon = 0
mask = (1 << lineLength) - 1

for line in ordered_data:
    for column in range(lineLength):
      bitFrequency[column] += int(line[column])

for column in bitFrequency:
  gamma = gamma << 1
  gamma += (column << 1) / len(ordered_data)

epsilon = mask ^ gamma

print(bitFrequency)
print("{0:b}".format(gamma))
print("{0:b}".format(epsilon))
print(gamma * epsilon)