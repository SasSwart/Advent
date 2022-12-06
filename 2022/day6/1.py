f = open("test.txt", "r")
signal = f.readline()

for i in range(4, len(signal)):
    potential_marker = signal[i-4:i]
    if len(set([i for i in potential_marker])) == 4:
        print(i)
        break

for i in range(14, len(signal)):
    potential_marker = signal[i-14:i]
    if len(set([i for i in potential_marker])) == 14:
        print(i)
        break