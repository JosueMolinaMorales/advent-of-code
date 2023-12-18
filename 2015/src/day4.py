import hashlib


def md5(key, leadingZeroes):
    hash = ""
    i = 0
    while not hash.startswith(leadingZeroes):
        i += 1
        hash = hashlib.md5((key+str(i)).encode()).hexdigest()
    return i


def partOne(key):
    return md5(key, "00000")


def partTwo(key):
    return md5(key, "000000")


if __name__ == "__main__":
    key = "yzbqklnj"
    print(partOne(key))
    print(partTwo(key))
