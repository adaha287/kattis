def day3():
    file = open("input/day3.txt", "r")
    prioPoints = 0
    for line in file:
        half = len(line) // 2
        s1 = line[:half]
        s2 = line[half:]
        prioPoints += getPoint(existInBoth(s1, s2))
        print(prioPoints)


def existInBoth(s1, s2):
    for c in s1:
        if c in s2:
            return c
    return ""


def getPoint(c):
    if c.islower():
        return ord(c) - 96
    else:
        return ord(c) - 38


if __name__ == "__main__":
    day3()
