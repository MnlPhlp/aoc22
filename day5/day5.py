def get_input():
    with open("input.txt") as f:
        return f.read()

def setup_stacks(input,stacks):
    stackLines = []
    for line in input.splitlines():
        if line.startswith("["):
            stackLines.append(line)
        else:
            stackCount = len(line.replace(" ",""))
            break
    print(f"creating {stackCount} stacks from following lines:")
    for line in stackLines:
        print(line)
    
    for i in range(stackCount):
        stacks.append([])
    for line in reversed(stackLines):
        for i in range(stackCount):
            item = line[1+i*4]
            if item != " ":
                stacks[i].append(item)

def setup_instructions(input,instructions):
    for line in input.splitlines():
        if line.startswith("move"):
            parts = line.split(" ")
            inst = [int(parts[1]), int(parts[3])-1, int(parts[5])-1]
            instructions.append(inst)

def do_moves(stacks,instructions,moveAtOnce):
    for inst in instructions:
        count = inst[0]
        fromStack = inst[1]
        toStack = inst[2]
        if moveAtOnce:
            items = stacks[fromStack][-count:]
            stacks[fromStack] = stacks[fromStack][:-count]
            stacks[toStack] += items          
        else:
            for i in range(count):
                item = stacks[fromStack].pop()
                stacks[toStack].append(item)

intputText = get_input()
stacks = []
instructions = []
setup_stacks(intputText,stacks)
stacks_copy =  [[e for e in stack] for stack in stacks]
setup_instructions(intputText,instructions)

# move one item at a time
do_moves(stacks,instructions,False)
for stack in stacks:
    print(" " if len(stack) == 0 else stack[-1],end="")
print()

# move all items at once
do_moves(stacks_copy,instructions,True)
for stack in stacks_copy:
    print(" " if len(stack) == 0 else stack[-1],end="")
print()

