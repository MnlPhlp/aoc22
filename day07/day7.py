maxSize = 100000

dirs = {"/": 0}
currentDir = "/"
parentDirs = []
with open("input.txt") as f:
    for line in f:
        # skip cd /
        if line == "$ cd /\n":
            continue
        # keep track of current directory
        elif line.startswith("$ cd .."):
            # add size to parent directory
            dirs[parentDirs[-1]] += dirs[currentDir]
            currentDir = parentDirs.pop()
        elif line.startswith("$ cd "):
            parentDirs.append(currentDir)
            currentDir = currentDir+line.split()[2]+"/"
            # add new directory
            if currentDir not in dirs:
                dirs[currentDir] = 0
        # add file to current directory
        elif line.split()[0].isdigit():
            size = int(line.split()[0])
            dirs[currentDir] += size
# get back to root
while currentDir != "/":
    dirs[parentDirs[-1]] += dirs[currentDir]
    currentDir = parentDirs.pop()

values = []
for d in dirs:
    if dirs[d] <= maxSize:
        values.append(dirs[d])

print("Total size of dirs <= 100000: ", sum(values))

fsSize = 70000000
print("Total used space: ", dirs["/"])
print("Free space: ", fsSize - dirs["/"])
neededSpace = 30000000 - (fsSize - dirs["/"])
print("Needed Space: ", neededSpace)
# find smallest dir to delete
minSize = fsSize
smallestDir = ""
for d in dirs:
    size = dirs[d]
    if size < minSize and size >= neededSpace:
        minSize = size
        smallestDir = d

print("smallest dir to delete: ", smallestDir, minSize)
