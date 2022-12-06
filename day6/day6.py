def get_input():
    with open("input.txt") as f:
        return f.read()

def get_index(count):
    input = get_input()
    chars = []

    for i,c in enumerate(input):
        # append new character and limit length to 4
        chars.append(c)
        chars = chars[-count:]
        if len(set(chars)) == count:
            return i+1


print(get_index(4))
print(get_index(14))