def is_password_valid(password):
    # Password must include one increasing straight of at least three letters
    three_letters = [password[i:i+3] for i in range(len(password)-2)]
    if not any(map(lambda s: ord(s[0])+2 == ord(s[1])+1 == ord(s[2]), three_letters)):
        return False

    # Password may not contain the letters i, o, l
    for c in ["i", "o", "l"]:
        if c in password:
            return False
    # Password must contain at least two different, non-overlapping pairs of letters
    pairs = {}
    for i in range(len(password)-1):
        pair = password[i:i+2]
        if pair in pairs:
            pairs[pair] += 1
        elif pair[0] == pair[1]:
            pairs[pair] = 1

    if len(pairs) < 2:
        return False
    return True


def inc_password(password):
    password = list(map(ord, password))
    i = len(password)-1
    wrapped = True
    while wrapped:
        password[i] += 1
        if password[i] > ord("z"):
            password[i] = ord("a")
            i -= 1
        else:
            wrapped = False
    password = "".join(map(chr, password))

    return password


def part_one(inp):
    password = inp
    stop = False
    while not stop:
        # Increment password
        password = inc_password(password)

        if is_password_valid(password):
            stop = True

    return password


def part_two(inp):
    password = part_one(inp)
    return part_one(password)


def run_day_eleven():
    input = "hxbxwxba"
    print("Day 11 part one:", part_one(input))
    print("Day 11 part two:", part_two(input))
