def contains(a, b):
  intersection = a.intersection(b)

  # Check if subset
  if intersection == a:
    return True
  if intersection == b:
    return True
  return False

def overlap(a, b):
  intersection = a.intersection(b)
  if len(intersection) > 0:
    return True
  return False

f = open("test.txt", "r")
vector_pairs = [[r.split("-") for r in line.split(",")] for line in f.read().split("\n")]
set_pairs = [[set(range(int(r[0]), int(r[1])+1)) for r in line] for line in vector_pairs]
intersections = [contains(pair[0], pair[1]) for pair in set_pairs]
overlaps = [overlap(pair[0], pair[1]) for pair in set_pairs]

print(len(list(filter(lambda a : a, intersections))))
print(len(list(filter(lambda a : a, overlaps))))