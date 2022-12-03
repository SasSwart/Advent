def duplicates(string_a, string_b):
  chars = dict()
  duplicates = []
  for char in string_a:
    chars[char] = 1

  for char in string_b:
    if char in chars:
      duplicates.append(char)
  
  return set(duplicates)

flatten = lambda l : [i for sl in l for i in sl]

def priority(c):
  offset = ord(c) - ord('a') + 1
  offset = offset + 58 if offset < 0 else offset
  return offset


f = open("test.txt", "r")
rucksacks = f.read().split("\n")
compartmentalised_rucksacks = [[r[:len(r)//2],r[len(r)//2:]] for r in rucksacks]

duplicates_per_rucksack = [duplicates(r[0],r[1]) for r in compartmentalised_rucksacks]
all_duplicates = flatten(duplicates_per_rucksack)

print(sum([priority(d) for d in all_duplicates]))