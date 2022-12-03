def duplicates(strings):
  if len(strings) == 1:
    return strings[0]
  
  chars = dict()
  found_duplicates = []
  for char in strings[0]:
    chars[char] = 1

  for char in strings[1]:
    if char in chars:
      found_duplicates.append(char)


  return set(duplicates([found_duplicates] + list(strings[2:])))
  
flatten = lambda l : [i for sl in l for i in sl]

def priority(c):
  offset = ord(c) - ord('a') + 1
  offset = offset + 58 if offset < 0 else offset
  return offset


f = open("test.txt", "r")
rucksacks = f.read().split("\n")
groups = [rucksacks[i:i+3] for i in range(0, len(rucksacks), 3)]

group_badges = [duplicates(r) for r in groups]
badge_priorities = flatten(group_badges)

print(sum([priority(b) for b in badge_priorities]))