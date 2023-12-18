VOWELS = ['a', 'e', 'i', 'o', 'u']
BAD_STRINGS = ['ab', 'cd', 'pq', 'xy']


def partOne(lines):
    niceCount = 0
    for s in lines:
        # check for bad strings
        bs = not any([bs in s for bs in BAD_STRINGS])
        # check for at least 3 vowels
        vc = len(list(filter(lambda c: c in VOWELS, s))) >= 3
        # check for repeating letter
        tr = any([s[i] == s[i+1] for i in range(len(s)-1)])
        niceCount += 1 if bs and vc and tr else 0
    return niceCount


def partTwo(lines):
    niceCount = 0
    for s in lines:
        # check for pair of letters
        pairs = any([s[i:i+2] in s[i+2:] for i in range(len(s)-2)])
        # check for repeating letter with one letter between
        repeat = any([s[i] == s[i+2] for i in range(len(s)-2)])
        niceCount += 1 if pairs and repeat else 0

    return niceCount


if __name__ == "__main__":
    with open("./input/day5.txt", "r") as f:
        lines = [x.strip() for x in f.readlines()]
        print(partOne(lines))
        print(partTwo(lines))
