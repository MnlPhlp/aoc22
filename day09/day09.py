class vec2:
    x = 0
    y = 0

    def __init__(self, x, y):
        self.x = x
        self.y = y

    def move(self, m):
        self.x += m.x
        self.y += m.y

    def touching(v1, v2):
        return abs(v1.x-v2.x) <= 1 and abs(v1.y-v2.y) <= 1

    def follow(v1, v2):
        dir = vec2(v2.x - v1.x, v2.y - v1.y)
        if dir.x != 0:
            v1.x += dir.x / abs(dir.x)
        if dir.y != 0:
            v1.y += dir.y / abs(dir.y)

    def __str__(self):
        return f"({self.x}, {self.y})"


def NewVec(dir):
    if dir == 'U':
        return vec2(0, 1)
    if dir == 'D':
        return vec2(0, -1)
    if dir == 'R':
        return vec2(1, 0)
    if dir == 'L':
        return vec2(-1, 0)
    return vec2(0, 0)


class move:
    direction = vec2(0, 0)
    distance = 0


def parseMoves(input):
    lines = input.splitlines()
    moves = [move() for _ in range(len(lines))]
    for i, line in enumerate(lines):
        if len(line) == 0:
            continue
        moves[i].direction = NewVec(line[0])
        moves[i].distance = int(line[2:])
    return moves


def makeMoves(moves, knotCount):
    visited = {}
    knots = [vec2(0, 0) for _ in range(knotCount)]
    visited[str(vec2(0, 0))] = True
    for m in moves:
        for rep in range(m.distance):
            # move head
            knots[0].move(m.direction)
            # move following knots
            prev = knots[0]
            for i in range(1, knotCount):
                while not knots[i].touching(prev):
                    knots[i].follow(prev)
                    if i == knotCount-1:
                        visited[str(knots[i])] = True
                prev = knots[i]
    return visited


with open("input.txt", "r") as myfile:
    input = myfile.read()
moves = parseMoves(input)
print(f"Doing {len(moves)} moves with 2 knots")
visited = makeMoves(moves, 2)
print(f"Visited {len(visited)} locations")

print(f"Doing {len(moves)} moves with 10 knots")
visited = makeMoves(moves, 10)
print(f"Visited {len(visited)} locations")
