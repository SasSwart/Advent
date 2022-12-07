import re

f = open("test.txt", "r")
input = f.readlines()

class Node:
    def __init__(self):
        self.parent = None
        self.children = {}
        self.ls_entry = ""
        self.sizes = []

    def Print(self):
        print(self.size)
        for c in self.children:
            print(self.children[c].Print())

    def GetSizes(self):
        sizes = self.sizes[:]
        for c in self.children:
            sizes.extend(self.children[c].GetSizes())
        return sizes
    
    def TotalSize(self):
        return sum(self.sizes) + sum([self.children[c].TotalSize() for c in self.children])

    def Traverse(self, f):
        f(self)
        for c in self.children:
            self.children[c].Traverse(f)


cur_node = Node()
root_node = cur_node

for line in input:
    if line == "$ cd ..\n":
        cur_node = cur_node.parent
    elif line.startswith("$ cd "):
        if line[5:] in cur_node.children:
            cur_node = cur_node.children[line[5:-1]]
        else:
            child = Node()
            child.parent = cur_node
            cur_node.children[line[5:-1]] = child
            cur_node = child
    if line[0] != "$":
        cur_node.ls_entry += line
        size = re.findall("^\d+", line)
        size = int(size[0]) if len(size) else 0
        cur_node.sizes.append(size)

one = 0
two = []

def getSizes(n):
    global one
    size = n.TotalSize()
    if size < 100000:
        one += size

def getBiggieSmall(n):
    global two
    size = n.TotalSize()
    if size >= 389918:
        two.append(size)

root_node.Traverse(getSizes)
root_node.Traverse(getBiggieSmall)

target = 70000000 - 30000000
delta = target - root_node.TotalSize()

print(delta)
print(one)
print(min(two))